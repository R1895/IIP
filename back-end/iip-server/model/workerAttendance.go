package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 工人考勤表
type WorkerAttendance struct {
	gorm.Model
	WorkerID      uint
	Date          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ArrivalTime   time.Time `gorm:"default:NULL"`
	DepartureTime time.Time `gorm:"default:NULL"`
	Type          string
	Filed1        string
	Filed2        string
	Filed3        string
	Filed4        string
}

type StatAttendance struct{
	WorkerID      uint
	TotalTime     string
	WorkerName  string   `json:"worker_name"`
}