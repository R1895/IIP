package model

import (
	"github.com/jinzhu/gorm"
	"time"
)



type SafeRate struct {
	gorm.Model
	WorkshopId   uint 
	Value        float64
	Date         time.Time
	Field1       string
	Field2       string
}

