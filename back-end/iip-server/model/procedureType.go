package model

import (
	"github.com/jinzhu/gorm"
)

type ProcedureType struct {
	gorm.Model
	ProcedureType string
	ProcedureName string
	Decription    string
	Field1        string
	Field2        string
}
