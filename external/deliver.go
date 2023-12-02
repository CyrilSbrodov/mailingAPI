package external

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"

	"mailingAPI/cmd/config"
	"mailingAPI/cmd/loggers"
	"mailingAPI/internal/storage"
	"mailingAPI/internal/storage/models"
)

const (
	statusOk     = "ok"
	statusFailed = "failed"
)

type Deliver struct {
	store  storage.Storage
	cfg    *config.ServerConfig
	logger *loggers.Logger
	ticker *time.Ticker
	wg     *sync.WaitGroup
	client *http.Client
}

func NewDeliver(cfg *config.ServerConfig, store storage.Storage, logger *loggers.Logger) *Deliver {
	ticker := time.NewTicker(cfg.CheckInterval)
	wg := &sync.WaitGroup{}
	client := &http.Client{}
	return &Deliver{
		store:  store,
		cfg:    cfg,
		logger: logger,
		ticker: ticker,
		wg:     wg,
		client: client,
	}
}

// Run запуск доставщика
func (d *Deliver) Run() {
	for range d.ticker.C {
		clients, err := d.store.ActiveProcessMailing()
		if len(clients) == 0 || err != nil && errors.Is(err, pgx.ErrNoRows) {
			d.logger.LogInfo("deliver", "mailing", "there no active mailing")
			continue
		}
		for _, m := range clients {
			go func(am []models.ActiveMailing) {
				timer := time.NewTimer(am[0].TimeEnd.Sub(time.Now()))
				d.SendMessage(am, timer)
			}(m)
		}
	}
}

// SendMessage отправляет сообщение клиенту через внешний сервис
func (d *Deliver) SendMessage(am []models.ActiveMailing, timer *time.Timer) {
	for i := 0; i < len(am); i++ {
		select {
		case <-timer.C:
			timer.Stop()
			break
		default:
			var dm models.DeliverMailing
			var err error

			dm.ID = am[i].MailId
			dm.Phone, err = strconv.Atoi(am[i].Client.PhoneNumber)
			if err != nil {
				d.logger.LogErr(err, "failed to convert phone")
				continue
			}
			dm.Text = am[i].Message

			// Создаем JSON-запрос
			requestBody, err := json.Marshal(dm)
			if err != nil {
				d.logger.LogErr(err, "failed to convert to JSON")
				continue
			}

			// Отправляем HTTP-запрос
			req, err := http.NewRequest("POST", d.cfg.ExternalAddr+strconv.Itoa(dm.ID), bytes.NewBuffer(requestBody))
			if err != nil {
				d.logger.LogErr(fmt.Errorf("ошибка при создании HTTP-запроса: %v", err), "")
				continue
			}
			req.Header.Set("Authorization", "Bearer "+d.cfg.Token)
			req.Header.Set("Content-Type", "application/json")

			resp, err := d.client.Do(req)
			if err != nil {
				d.logger.LogErr(fmt.Errorf("ошибка при отправке HTTP-запроса: %v", err), "")
				continue
			}
			// Обрабатываем ответ
			if resp.StatusCode != http.StatusOK {
				am[i].Status = statusFailed
				d.logger.LogErr(fmt.Errorf("некорректный HTTP-статус от внешнего сервиса: %d", resp.StatusCode), "")
			} else {
				am[i].TimeSend = time.Now()
				am[i].Status = statusOk
			}

			resp.Body.Close()
		}
	}
	if err := d.store.UpdateStatusMessage(am); err != nil {
		d.logger.LogErr(err, "")
		return
	}
}
