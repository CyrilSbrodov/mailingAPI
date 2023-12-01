package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"mailingAPI/internal/storage/models"
)

// SendMessage отправляет сообщение клиенту через внешний сервис
func SendMessage(client models.Client, message string) error {
	externalServiceURL := "https://probe.fbrq.cloud/send" // Замените на реальный URL вашего внешнего сервиса

	// Создаем JSON-запрос
	requestBody, err := json.Marshal(map[string]interface{}{
		"phone_number": client.PhoneNumber,
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
	defer resp.Body.Close()

	// Обрабатываем ответ
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("некорректный HTTP-статус от внешнего сервиса: %d", resp.StatusCode)
	}

	return nil
}
