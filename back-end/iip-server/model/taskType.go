package model

import (
	"github.com/jinzhu/gorm"
)

type TaskType struct {
	gorm.Model
	TaskType    string `gorm:"unique"`
	Tasks       []Task
	Description string
	Field1      string
	Field2      string
	Field3      string
	Field4      string
}
