package model

import (
	"github.com/jinzhu/gorm"
)

type MachineStatus struct {
	gorm.Model
	StatusType  uint `gorm:"unique"`
	StatusName  string
	Machines    []Machine
	Description string
	Field1      string
	Field2      string
	Field3      string
	Field4      string
}
