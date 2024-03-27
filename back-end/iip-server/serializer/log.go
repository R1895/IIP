package serializer

import (
	"iip/model"
	"time"
)

type Log struct {
	LogID     uint      `json:"log_id"`
	UserID    uint      `json:"user_id"`
	Action    string    `json:"action"`
	Method    string    `json:"method"`
	Timestamp time.Time `json:"timestamp"`
	Details   string    `json:"details"`
}

func BuildLog(log model.Log) Log {
	return Log{
		LogID:     log.ID,
		UserID:    log.UserID,
		Action:    log.Action,
		Method:    log.Method,
		Timestamp: log.Timestamp,
		Details:   log.Details,
	}
}

func BuildLogs(items []model.Log) (logs []Log) {
	for _, item := range items {
		log := BuildLog(item)
		logs = append(logs, log)
	}
	return logs
}
