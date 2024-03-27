package serializer

import (
	"iip/model"
	"time"
)

type Safe struct {
	ID          uint      `json:"id"`
	WorkShopId  uint      `json:"workshop_id"`
	Value       float64   `json:"value"`
	Date        time.Time `json:"date"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
}

type SafeInformation struct {
	Type       string  `json:"type"`
	TotalValue float64 `json:"total_value"`
}

type NumDaySafe struct {
	Workshop_ID uint `json:"workshop_id"`
	Count       uint `json:"count"`
	SafeInformations       []SafeInformation `json:"safes"`
	Total       uint `json:"total"`   
}

// BuildSafe 序列化角色
func BuildSafe(safe model.Safe) Safe {
	return Safe{
		ID:          safe.ID,
		WorkShopId:  safe.WorkshopID,
		Value:       safe.Value,
		Date:        safe.Date,
		Type:        safe.Type,
		Description: safe.Description,
	}
}

func BuildSafes(items []model.Safe) (safes []Safe) {
	for _, item := range items {
		safe := BuildSafe(item)
		safes = append(safes, safe)
	}
	return safes
}

func BuildSafeInformation(safeInformation model.SafeInformation) SafeInformation{
	return SafeInformation{
		Type: safeInformation.Type,
		TotalValue:  safeInformation.TotalValue,
	}
}

func BuildSafeInformations(items []model.SafeInformation) (safeinformations []SafeInformation) {
	for _, item := range items {
		safeinformation := BuildSafeInformation(item)
		safeinformations = append(safeinformations, safeinformation)
	}
	return safeinformations
}