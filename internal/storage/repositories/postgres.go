package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"mailingAPI/cmd/config"
	"mailingAPI/cmd/loggers"
	"mailingAPI/internal/storage/models"
	"mailingAPI/pkg/postgres"
)

type PGStore struct {
	client postgres.Client
	cfg    *config.ServerConfig
	logger *loggers.Logger
}

// createTable - функция создания новых таблиц в БД.
func createTable(ctx context.Context, client postgres.Client, logger *loggers.Logger) error {
	tx, err := client.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.LogErr(err, "failed to begin transaction")
		return err
	}
	defer tx.Rollback(ctx)

	//создание таблиц
	q := `CREATE TABLE if not exists client (
    		id SERIAL PRIMARY KEY,
    		phone_number VARCHAR(15) NOT NULL,
    		tag VARCHAR(50),
    		code VARCHAR(50) NOT NULL,
    		timezone VARCHAR(50),
    		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE if not exists mailing (
    		id SERIAL PRIMARY KEY,
    		time_start TIMESTAMPTZ NOT NULL,
    		time_end TIMESTAMPTZ NOT NULL,
    		code VARCHAR(50),
    		tag VARCHAR(50),
    		message TEXT,
    		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP                       
		);
		CREATE TABLE if not exists messages (
    		id SERIAL PRIMARY KEY,
    		mail_id INT REFERENCES mailing(id) NOT NULL,
    		client_id INT REFERENCES client(id) NOT NULL,
    		status varchar(50) NOT NULL,
    		send_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP                     
		);`

	_, err = tx.Exec(ctx, q)
	if err != nil {
		logger.LogErr(err, "failed to create table")
		return err
	}
	return tx.Commit(ctx)
}

func NewPGStore(client postgres.Client, cfg *config.ServerConfig, logger *loggers.Logger) (*PGStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := createTable(ctx, client, logger); err != nil {
		logger.LogErr(err, "failed to create table")
		return nil, err
	}
	return &PGStore{
		client: client,
		cfg:    cfg,
		logger: logger,
	}, nil
}

func (p *PGStore) AddClient(c *models.Client) error {
	q := `INSERT INTO client (phone_number, code)
    						VALUES ($1, $2)`
	if _, err := p.client.Exec(context.Background(), q, c.PhoneNumber, c.MobileOperator); err != nil {
		p.logger.LogErr(err, "Failure to insert object into table")
		return err
	}
	return nil
}
func (p *PGStore) UpdateClient(c *models.Client) error {
	q := `UPDATE client SET phone_number = $2, code = $3 WHERE id = $1`
	if _, err := p.client.Exec(context.Background(), q, c.ID, c.PhoneNumber, c.MobileOperator); err != nil {
		p.logger.LogErr(err, "Failure to update object into table")
		return err
	}
	return nil
}
func (p *PGStore) DeleteClient(c *models.Client) error {
	q := `DELETE FROM client WHERE id = $1`
	if _, err := p.client.Exec(context.Background(), q, c.ID); err != nil {
		p.logger.LogErr(err, "Failure to delete object into table")
		return err
	}
	return nil
}
func (p *PGStore) AddMailing(m *models.Mailing) error {
	q := `INSERT INTO mailing (time_start, time_end, code, tag, message)
    						VALUES ($1, $2, $3, $4, $5, $6)`
	if _, err := p.client.Exec(context.Background(), q, m.TimeStart, m.TimeEnd, m.Filter.MobileOperator, m.Filter.Tag, m.Message); err != nil {
		p.logger.LogErr(err, "Failure to insert object into table")
		return err
	}
	return nil
}
func (p *PGStore) GetAllMailingStat() (models.AllStatistics, error) {
	var stat models.AllStatistics
	q := `SELECT m.id AS mailing_id, m.message,
    COUNT(msg.id) AS total_messages,
    SUM(CASE WHEN msg.status = 'ok' THEN 1 ELSE 0 END) AS sent_messages,
    SUM(CASE WHEN msg.status = 'failed' THEN 1 ELSE 0 END) AS failed_messages
	FROM mailing m
    LEFT JOIN messages msg ON m.id = msg.mail_id
	GROUP BY
    m.id, m.message
	ORDER BY
    m.id;`
	err := p.client.QueryRow(context.Background(), q).Scan(&stat.MailingID, &stat.Message, &stat.TotalMessages, &stat.SentMessages, &stat.FailedMessages)
	if err != nil {
		p.logger.LogErr(err, "Failure to insert object into table")
		return stat, err
	}
	return stat, nil
}
func (p *PGStore) GetOneMailingStatByID(m *models.Mailing) (models.Mailing, error) {
	//q := `SELECT m.id AS mail_id, c.id AS client_id, c.phone_number, msg.send_time, msg.status, m.message
	//		FROM messages msg
	//		JOIN mailing m ON msg.mail_id = m.id
	//		JOIN client c ON msg.client_id = c.id
	//		WHERE m.id = $1
	//		ORDER BY mail_id;`

	return *m, nil
}
func (p *PGStore) UpdateMailing(m *models.Mailing) error {
	q := `UPDATE mailing SET time_start = $2, time_end = $3, code = $4, tag = $5, message = $6
               WHERE id = $1`
	if _, err := p.client.Exec(context.Background(), q, m.ID, m.TimeStart, m.TimeEnd, m.Filter.MobileOperator,
		m.Filter.Tag, m.Message); err != nil {
		p.logger.LogErr(err, "Failure to update object into table")
		return err
	}
	return nil
}
func (p *PGStore) DeleteMailing(m *models.Mailing) error {
	q := `DELETE FROM mailing WHERE id = $1`
	if _, err := p.client.Exec(context.Background(), q, m.ID); err != nil {
		p.logger.LogErr(err, "Failure to delete object into table")
		return err
	}
	return nil
}

func (p *PGStore) ActiveProcessMailing() (map[string][]models.ActiveMailing, error) {
	activeMailing := make(map[string][]models.ActiveMailing)
	q := `SELECT mailing.id, mailing.message, mailing.time_end, c.id AS client_id, c.phone_number
				FROM mailing JOIN client c ON mailing.tag = c.tag AND mailing.code = c.code
         		LEFT JOIN messages m ON mailing.id = m.mail_id AND c.id = m.client_id
				WHERE now() AT TIME ZONE 'Europe/Moscow' BETWEEN mailing.time_start AND mailing.time_end
  				AND (m.id IS NULL OR NOT c.id = m.client_id) 
				AND NOT EXISTS (
    							SELECT 1
    							FROM messages msg
    							WHERE msg.mail_id = mailing.id
								);`
	rows, err := p.client.Query(context.Background(), q)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.LogErr(err, "There are no mailing lists")
			return nil, err
		}
		p.logger.LogErr(err, "")
		return nil, err
	}
	for rows.Next() {
		var m models.ActiveMailing
		if err = rows.Scan(&m.MailId, &m.Message, &m.TimeEnd, &m.Client.ID, &m.Client.PhoneNumber); err != nil {
			p.logger.LogErr(err, "")
			return nil, err
		}
		activeMailing[m.Message] = append(activeMailing[m.Message], m)
	}
	return activeMailing, nil
}

func (p *PGStore) UpdateStatusMessage(am []models.ActiveMailing) error {
	q := `INSERT INTO messages (mail_id, client_id, status, send_time) VALUES ($1, $2, $3, $4)`
	tx, err := p.client.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		p.logger.LogErr(err, "failed to begin transaction")
		return err
	}
	defer tx.Rollback(context.Background())

	for _, m := range am {
		if _, err := p.client.Exec(context.Background(), q, m.MailId, m.Client.ID, m.Status, m.TimeSend); err != nil {
			p.logger.LogErr(err, "Failure to insert object from table")
			return err
		}
	}
	return tx.Commit(context.Background())
}
