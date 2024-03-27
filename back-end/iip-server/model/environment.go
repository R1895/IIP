package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 生产环境
type Environment struct {
	gorm.Model
	WorkshopID  uint
	Level       uint
	Date        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Description string
	Field1      string
	Field2      string
	Field3      string
}
