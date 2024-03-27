package model

import (
	"github.com/jinzhu/gorm"
)

// 任务状态表 查看历史状态
type TaskStatus struct {
	gorm.Model
	TaskStatus  string
	Tasks       []Task
	Description string
	Field1      string
	Field2      string
	Field3      string
}
