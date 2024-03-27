package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 噪声
type Noise struct {
	gorm.Model
	WorkshopId uint
	Value      float64
	Date       time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	Description string
	Field1     string
	Field2     string
}
