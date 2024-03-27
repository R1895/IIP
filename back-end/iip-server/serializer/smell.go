package serializer

import (
	"iip/model"
	"time"
)

type Smell struct {
	ID          uint      `json:"id"`
	WorkshopID  uint      `json:"workshop_id"`
	Value       float64   `json:"value"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type DaySmell struct {
	Workshop_ID uint    `json:"workshop_id"`
	Count       uint    `json:"count"`
	Avg         float64 `json:"avg"`
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Smells      []Smell `json:"smells"`
	Smell       Smell   `json:"smell"`
}

func BuildSmell(smell model.Smell) Smell {
	return Smell{
		ID:          smell.ID,
		WorkshopID:  smell.WorkshopID,
		Value:       smell.Value,
		Date:        smell.Date,
		Description: smell.Description,
	}
}

func BuildSmells(items []model.Smell) (smells []Smell) {
	for _, item := range items {
		smell := BuildSmell(item)
		smells = append(smells, smell)
	}
	return smells
}
