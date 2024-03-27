package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 生产任务表
type Task struct {
	gorm.Model
	WorkshopID    uint
	TaskName      string
	Description   string
	StartDate     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	EffectiveTime time.Time `gorm:"default:NULL"`
	TaskStatusID  uint
	IsFinished    uint // 0表示未完成，1表示完成
	Procedures    []Procedure
	TaskTypeID    uint
	Field1        string
	Field2        string
	Field3        string
	Field4        string
}
