package model

import (
	"github.com/jinzhu/gorm"
	"time"
)


type Log struct {
	gorm.Model
	UserID    uint // 外键，指向用户表
	Action    string
	Method    string
	Timestamp time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Details   string
	Field1    string 
	Field2    string
	Field3    string
	Field4    string
}
