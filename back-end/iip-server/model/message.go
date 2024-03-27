package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 消息表 给用户发通知
type Message struct {
	gorm.Model
	MachineID uint
	UserID    uint // 外键，指向用户表
	IsRead    bool
	SentDate  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Title     string
	Content   string
	Type      string
	Field1    string
	Field2    string
	Field3    string
	Field4    string
}
