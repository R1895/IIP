package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 工序表
type Procedure struct {
	gorm.Model
	ProcedureType   string      `gorm:"comment:'工序类别'"`
	ProcedureName   string      `gorm:"comment:'工序名'"`   
	TaskID          uint      `gorm:"comment:'任务ID'"`
	MachineID       uint      `gorm:"comment:'设备ID'"`
	WorkpiecesNum   uint      `gorm:"comment:'工件数量'"`
	ProcessedNum    uint      `gorm:"comment:'已加工数量'"`
	Sequence        uint      `gorm:"comment:'工序顺序'"`
	MachineSequence uint      `gorm:"comment:'设备的顺序'"`
	StartDate       time.Time `gorm:"comment:'开始日期'"`
	IsFinished      uint      `gorm:"comment:'是否结束'"`
	ProcedureStatusID uint
	ProcedureItems   []ProcedureItem
	Field1          string
	Field2          string
}
