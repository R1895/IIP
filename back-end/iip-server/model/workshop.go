package model

import (
	"github.com/jinzhu/gorm"
)

// 车间表
type Workshop struct {
	gorm.Model
	WorkshopName string `gorm:"unique;comment:'车间名'"`
	WorkerId     uint   `gorm:"not null"` // 外键，关联工人表
	Worker       Worker
	Workers      []Worker `gorm:"many2many:worker_workshop"`
	Machines     []Machine
	Tasks        []Task
	Environments []Environment
	Noises       []Noise
	Smells       []Smell
	Safes        []Safe
	Field1       string
	Field2       string
	Field3       string
	Field4       string
}


