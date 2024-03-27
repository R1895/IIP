package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 机器设备表
type Machine struct {
	gorm.Model
	WorkshopID      uint
	ProcedureType   string
	MachineName     string
	MachineStatusID uint      // 外键，关联设备状态表
	CurrentTaskID   uint      // 当前执行的任务ID
	PurchaseDate    time.Time `gorm:"default:NULL"`
	WarrantyUntil   time.Time `gorm:"default:NULL"`
	MachineWorkers  []MachineWorker
	MachineAnomalys []MachineAnomaly
	MachineLogs     []MachineLog
	Procedures      []Procedure
	Field1          string
	Field2          string
	Field3          string
	Field4          string
}
