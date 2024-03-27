package model

import (
	"github.com/jinzhu/gorm"
)

// 工件
type ProcedureItem struct {
	gorm.Model
	ProcedureID uint
	MachineID   uint
	WorkerID    uint
	ItemType    string
	ItemName    string
	Quality     string
	Field1      string
	Field2      string
	Field3      string
	Field4      string
}
