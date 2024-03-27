package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 设备日志表
type MachineLog struct {
	gorm.Model
	MachineID       uint
	Time            time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	MachineStatusID uint
	Description     string
	Field1          string
	Field2          string
	Field3          string
	Field4          string
}
