package serializer

import (
	"iip/model"
	"time"
)

type Message struct {
	MessageID uint      `json:"message_id"`
	MachineID uint      `json:"machine_id"`
	UserID    uint      `json:"user_id"`
	IsRead    bool      `json:"is_read"`
	SentDate  time.Time `json:"sent_date"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Type      string    `json:"type"`
}


func BuildMessage(message model.Message) Message {
	return Message{
		MessageID: message.ID,
		MachineID: message.MachineID,
		UserID:    message.UserID,
		IsRead:    message.IsRead,
		SentDate:  message.SentDate,
		Title:     message.Title,
		Content:   message.Content,
		Type:      message.Type,
	}
}

func BuildMessages(items []model.Message) (messages []Message) {
	for _, item := range items {
		message := BuildMessage(item)
		messages = append(messages, message)
	}
	return messages
}
