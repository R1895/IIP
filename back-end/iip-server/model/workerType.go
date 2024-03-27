package model

import (
	"github.com/jinzhu/gorm"
)

type WorkerType struct {
	gorm.Model
	WorkerType  string `gorm:"unique"`
	Workers     []Worker
	Description string
	Field1      string
	Field2      string
	Field3      string
	Field4      string
}

