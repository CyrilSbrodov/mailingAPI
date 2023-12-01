package models

import "time"

type Mailing struct {
	ID        int       `json:"id"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
	Filter    Filter    `json:"filter"`
	Messages  []Message `json:"messages"`
}

type Filter struct {
	MobileOperator int    `json:"mobile_operator"`
	Tag            string `json:"tag"`
}

type Client struct {
	ID             int    `json:"id"`
	PhoneNumber    int    `json:"phone_number"`
	MobileOperator int    `json:"mobile_operator"`
	Tag            string `json:"tag"`
	TimeZone       string `json:"time_zone"`
}

type Message struct {
	ID        int       `json:"id"`
	MailId    int       `json:"mail_id"`
	ClientId  int       `json:"client_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Statistics struct {
	MailingCount  int
	MessagesCount map[string]int // статус сообщения -> количество сообщений с таким статусом
}
