package models

import (
	"regexp"
	"time"
)

type Mailing struct {
	ID        int       `json:"id"`
	Message   string    `json:"messages"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
	Filter    Filter    `json:"filter"`
}

type Filter struct {
	MobileOperator string `json:"mobile_operator"`
	Tag            string `json:"tag"`
}

type Client struct {
	ID             int    `json:"id"`
	PhoneNumber    string `json:"phone_number"`
	MobileOperator string `json:"mobile_operator"`
	Tag            string `json:"tag"`
	TimeZone       string `json:"time_zone"`
}

type Message struct {
	ID       int       `json:"id"`
	MailId   int       `json:"mail_id"`
	ClientId int       `json:"client_id"`
	Status   string    `json:"status"`
	SendTime time.Time `json:"send_time"`
}

type AllStatistics struct {
	MailingID      int
	Message        string
	TotalMessages  int
	SentMessages   int
	FailedMessages int
}

type Statistics struct {
}

type ActiveMailing struct {
	MailId   int
	Message  string
	Status   string
	TimeEnd  time.Time
	TimeSend time.Time
	Filter   Filter
	Client   Client
}

type ProcessMailing struct {
	MailingId      int
	Message        string
	Status         string
	MobileOperator string
	Tag            string
	TimeEnd        time.Time
	TimeSend       time.Time
	Client         []Client
}

type DeliverMailing struct {
	ID    int    `json:"id"`
	Phone int    `json:"phone"`
	Text  string `json:"text"`
}

func (c *Client) CheckNumber() bool {
	// Создаем регулярное выражение для проверки формата телефонного номера.
	phoneRegex := regexp.MustCompile(`^7[0-9]{10}$`)

	// Проверяем, соответствует ли номер формату.
	return phoneRegex.MatchString(c.PhoneNumber)
}
