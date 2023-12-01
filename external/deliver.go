package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"mailingAPI/cmd/config"
	"mailingAPI/internal/storage"
	"mailingAPI/internal/storage/models"
)

type Scheduler struct {
	store  storage.Storage
	cfg    *config.ServerConfig
	timer  *time.Timer
	ticker *time.Ticker
}

func NewScheduler(cfg *config.ServerConfig, store storage.Storage) *Scheduler {
	ticker := time.NewTicker(cfg.CheckInterval)
	return &Scheduler{
		store:  store,
		cfg:    cfg,
		ticker: ticker,
	}
}

func (s *Scheduler) CheckNewMailing() error {
	for range s.ticker.C {
		clients, mailing, err := s.store.ActiveProcessMailing()
		if err != nil {
			return err
		}
		if len(mailing) == 0 {
			continue
		} else {
			go func() {
				if err := s.SendMessage(clients, mailing); err != nil {
					//TODO error
					return err
				}
			}()
		}
	}
	return nil
}

// SendMessage отправляет сообщение клиенту через внешний сервис
func (s *Scheduler) SendMessage(client []models.Client, message []models.Mailing) error {
	for _, m := range client {
		//TODO add external addr to cfg
		externalServiceURL := "https://probe.fbrq.cloud/send" // Замените на реальный URL вашего внешнего сервиса

		// Создаем JSON-запрос
		requestBody, err := json.Marshal(map[string]interface{}{
			"phone_number": m.PhoneNumber,
			"message":      message,
		})
		if err != nil {
			return fmt.Errorf("ошибка при формировании JSON-запроса: %v", err)
		}

		// Отправляем HTTP-запрос
		resp, err := http.Post(externalServiceURL, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			return fmt.Errorf("ошибка при отправке HTTP-запроса: %v", err)
		}

		// Обрабатываем ответ
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("некорректный HTTP-статус от внешнего сервиса: %d", resp.StatusCode)
		}

		resp.Body.Close()

	}
	return nil
}
