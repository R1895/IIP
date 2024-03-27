package model

import (
	"github.com/jinzhu/gorm"
)

// 数据字典
type DictionaryItem struct {
	gorm.Model
	Key         string
	Value       string
	Description string
	Category    string
	Field1      string
	Field2      string
	Field3      string
	Field4      string
}
