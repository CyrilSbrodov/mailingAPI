package repositories

import (
	"context"
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
    		code BIGINT NOT NULL
		);
		CREATE TABLE if not exists mailing (
    		id SERIAL PRIMARY KEY,
    		time_start TIMESTAMPTZ,
    		time_end TIMESTAMPTZ,
    		code VARCHAR(50),
    		tag VARCHAR(50),
    		content TEXT                         
		);
		CREATE TABLE if not exists message (
    		id BIGINT PRIMARY KEY generated always as identity,
    		mail_id INT REFERENCES mailing(id),
    		client_id INT REFERENCES client(id),
    		status varchar(50),
    		time_start TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP                       
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
	q := `INSERT INTO mailing (time_start, time_end, code, tag, content)
    						VALUES ($1, $2, $3, $4, $5)`
	if _, err := p.client.Exec(context.Background(), q, m.TimeStart, m.TimeEnd, m.Filter.MobileOperator, m.Filter.Tag, m.Messages); err != nil {
		p.logger.LogErr(err, "Failure to insert object into table")
		return err
	}
	return nil
}
func (p *PGStore) GetAllMailingStat() ([]models.Mailing, error) {
	var m []models.Mailing
	q := `SELECT id, time_start, time_end, code, tag, content FROM mailing`
	rows, err := p.client.Query(context.Background(), q)
	if err != nil {
		p.logger.LogErr(err, "Failure to insert object into table")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var mail models.Mailing
		err = rows.Scan(&mail.ID, &mail.TimeStart, &mail.TimeEnd, &mail.Filter.MobileOperator,
			&mail.Filter.Tag, mail.Messages)
		if err != nil {
			p.logger.LogErr(err, "Failure to convert object from table")
			return nil, err
		}
		m = append(m, mail)
	}
	return m, nil
}
func (p *PGStore) GetOneMailingStat(m *models.Mailing) (models.Mailing, error) {
	return *m, nil
}
func (p *PGStore) UpdateMailing(m *models.Mailing) error {
	q := `UPDATE mailing SET time_start = $2, time_end = $3, code = $4, tag = $5, content = $6
               WHERE id = $1`
	if _, err := p.client.Exec(context.Background(), q, m.ID, m.TimeStart, m.TimeEnd, m.Filter.MobileOperator,
		m.Filter.Tag, m.Messages); err != nil {
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

func (p *PGStore) ActiveProcessMailing() ([]models.Mailing, error) {
	q := `SELECT id, time_start, time_end, code, tag, content FROM mailing`
	q1 := `SELECT id, phone_number, code FROM client`
}

func (p *PGStore) UpdateStatus(m *models.Message) error {
	q := `UPDATE message SET status = $2 WHERE client_id = $1`
	if _, err := p.client.Exec(context.Background(), q, m.ID); err != nil {
		p.logger.LogErr(err, "Failure to update object from table")
		return err
	}
	return nil
}
