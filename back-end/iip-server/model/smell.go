package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 气味
type Smell struct {
	gorm.Model
	WorkshopID uint
	Value      float64
	Date       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Description string
	Field1     string
	Field2     string
}
