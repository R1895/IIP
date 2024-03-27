package model

import (
	"github.com/jinzhu/gorm"
)

// 工序表
type ProcedureStatus struct {
	gorm.Model
	ProcedureStatus string
	Procedures      []Procedure
	Description     string
	Field1          string
	Field2          string
	Field3          string
}
