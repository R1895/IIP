package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 安全
type Safe struct {
	gorm.Model
	WorkshopID  uint
	Value       float64
	Type        string
	Description string
	Date        time.Time  `gorm:"default:CURRENT_TIMESTAMP"`

	Field1      string
	Field2      string
}

type SafeInformation struct {
	Type       string  
	TotalValue float64 
}