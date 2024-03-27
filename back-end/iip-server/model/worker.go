package model

import (
	"github.com/jinzhu/gorm"
)

// 工人表
type Worker struct {
	gorm.Model
	UserID            uint
	WorkerTypeID      uint
	WorkerType        string
	WorkerName        string
	SkillLevel        int
	Workshops         []Workshop `gorm:"many2many:worker_workshop"`
	MachineWorkers    []MachineWorker
	WorkerAttendances []WorkerAttendance

	Field1 string
	Field2 string
	Field3 string
	Field4 string
}
