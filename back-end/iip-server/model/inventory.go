package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 库存  TODO type 
type Iventory struct {
	gorm.Model
	ItemType     string    //工件类型
	IventoryType uint      //区分入库出库，值为1表示入库，值为2表示出库
	IventoryNum  uint      //工件数量
	Date         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Description  string
	Field1       string    //大类
	Field2       string
	Field3       string
}

