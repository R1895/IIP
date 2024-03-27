package serializer

import (
	"iip/model"
	"time"
)

// 库存
type Iventory struct {
	ID           uint      `json:"id"`
	ItemType     string    `json:"item_type"`
	IventoryType uint      `json:"iventory_type"`
	IventoryNum  uint      `json:"iventory_num"`
	Date         time.Time `json:"date"`
	Description  string    `json:"description"`
	Field1       string    `json:"field1"`
}

type StatIventory struct {
	Field1   string `json:"field1"`
	ItemType string `json:"item_type"`
	StatNum  uint   `json:"stat_num"`
}

type StatIventorys struct {
	StatIventorys []StatIventory `json:"stativentorys"`
	Total        int64          `json:"total"`
}

func BuildIventory(iventory model.Iventory) Iventory {
	return Iventory{
		ID:           iventory.ID,
		ItemType:     iventory.ItemType,
		IventoryType: iventory.IventoryType,
		IventoryNum:  iventory.IventoryNum,
		Date:         iventory.Date,
		Description:  iventory.Description,
		Field1:       iventory.Field1,
	}
}

func BuildIventorys(items []model.Iventory) (Iventorys []Iventory) {
	for _, item := range items {
		Iventory := BuildIventory(item)
		Iventorys = append(Iventorys, Iventory)
	}
	return Iventorys
}
