package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 设备异常表
type MachineAnomaly struct {
	gorm.Model
	Time            time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	MachineID       uint
	MachineStatusID uint
	Description     string
	Field1          string
	Field2          string
	Field3          string
	Field4          string
}
