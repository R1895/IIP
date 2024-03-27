package serializer

import (
	"iip/model"
	"time"
)

type MachineAnomaly struct {
	ID              uint      `json:"id"`
	Time            time.Time `json:"time"`
	MachineID       uint      `json:"machine_id"`
	MachineStatusID uint      `json:"machine_status_id"`
	Description     string    `json:"description"`
}

func BuildMachineAnomaly(machineAnomaly model.MachineAnomaly) MachineAnomaly {
	return MachineAnomaly{
		ID:              machineAnomaly.ID,
		Time:            machineAnomaly.Time,
		MachineID:       machineAnomaly.MachineID,
		MachineStatusID: machineAnomaly.MachineStatusID,
		Description:     machineAnomaly.Description,
	}
}

func BuildMachineAnomalys(items []model.MachineAnomaly) (machineAnomalys []MachineAnomaly) {
	for _, item := range items {
		machineAnomaly := BuildMachineAnomaly(item)
		machineAnomalys = append(machineAnomalys, machineAnomaly)
	}
	return machineAnomalys
}
