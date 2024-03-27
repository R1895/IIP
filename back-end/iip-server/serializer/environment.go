package serializer

import (
	"iip/model"
	"time"
)

type Environment struct {
	EnvironmentID uint      `json:"environment_id"`
	Level         uint      `json:"level"`
	WorkshopID    uint      `json:"workshop_id"`
	Date          time.Time `json:"timestamp"`
	Description   string    `json:"description"`
}

type NumDayEnvironment struct {
	WorkshopID uint `json:"workshop_id"`
	Count      uint `json:"count"`
}

func BuildEnvironment(environment model.Environment) Environment {
	return Environment{
		EnvironmentID: environment.ID,
		Level:         environment.Level,
		WorkshopID:    environment.WorkshopID,
		Date:          environment.Date,
		Description:   environment.Description,
	}
}

func BuildEnvironments(items []model.Environment) (environments []Environment) {
	for _, item := range items {
		environment := BuildEnvironment(item)
		environments = append(environments, environment)
	}
	return environments
}
