package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MachineWorker struct {
	gorm.Model
	MachineID uint      // 外键，指向机器设备表
	WorkerID  uint      // 外键，指向工人表
	StartDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	EndDate   time.Time `gorm:"default:NULL"`
	Field1    string
	Field2    string
	field3    string
	Field4    string
}
