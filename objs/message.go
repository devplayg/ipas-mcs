package objs

import "time"

const (

	UnreadMessage = 1

	InfoMessage = 1 << iota
	WarningMessage
	DangerMessage

)

type Message struct {
	MessageId  int `json:"message_id"`
	Date       time.Time `json:"date"`
	Status     int `json:"status"`
	ReceiverId int `json:"receiver_id"`
	SenderId   int `json:"sender_id"`
	Priority   int `json:"priority"`
	Category   int `json:"category"`
	Message    string `json:"message"`
	Url        string `json:"url"`
}

type MessageFilter struct {
	PagingFilter
	MessageId  int
	date       time.Time
	Status     int
	ReceiverId int
	SenderId   int
	Priority   int
	Category   int
	Message    string
	Url        string
}
