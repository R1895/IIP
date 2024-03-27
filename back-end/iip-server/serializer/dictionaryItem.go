package serializer

import (
	"iip/model"
)

type DictionaryItem struct {
	ID          uint   `json:"id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func BuildDictionaryItem(dictionaryItem model.DictionaryItem) DictionaryItem {
	return DictionaryItem{
		ID:          dictionaryItem.ID,
		Key:         dictionaryItem.Key,
		Value:       dictionaryItem.Value,
		Description: dictionaryItem.Description,
		Category:    dictionaryItem.Category,
	}
}

func BuildDictionaryItems(items []model.DictionaryItem) (dictionaryItems []DictionaryItem) {
	for _, item := range items {
		dictionaryItem := BuildDictionaryItem(item)
		dictionaryItems = append(dictionaryItems, dictionaryItem)
	}
	return dictionaryItems
}
