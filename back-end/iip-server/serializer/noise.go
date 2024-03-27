package serializer

import (
	"iip/model"
	"time"
)

type Noise struct {
	ID         uint      `json:"id"`
	WorkshopId uint      `json:"workshop_id"`
	Value      float64   `json:"value"`
	Date       time.Time `json:"date"`
}

type NoiseStats struct {
	Avg         float64 `json:"avg"`
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Count       uint    `json:"count"`
	Workshop_ID uint    `json:"workshop_id"`
	Noises      []Noise `json:"noises"`

}

func BuildNoise(noise model.Noise) Noise {
	return Noise{
		ID:         noise.ID,
		WorkshopId: noise.WorkshopId,
		Value:      noise.Value,
		Date:       noise.Date,
	}
}

func BuildNoises(items []model.Noise) (noises []Noise) {
	for _, item := range items {
		noise := BuildNoise(item)
		noises = append(noises, noise)
	}
	return noises
}
