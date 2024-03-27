package serializer

import (
	"iip/model"
	"time"
)

type MachineLog struct {
	ID              uint      `json:"id"`
	Time            time.Time `json:"time"`
	MachineID       uint      `json:"machine_id"`
	MachineStatusID uint      `json:"machine_status_id"`
	Description     string    `json:"description"`
}

// BuildMachineLog 序列化角色
func BuildMachineLog(machineLog model.MachineLog) MachineLog {
	return MachineLog{
		ID:              machineLog.ID,
		Time:            machineLog.Time,
		MachineID:       machineLog.MachineID,
		MachineStatusID: machineLog.MachineStatusID,
		Description:     machineLog.Description,
	}
}

func BuildMachineLogs(items []model.MachineLog) (machineLogs []MachineLog) {
	for _, item := range items {
		machineLog := BuildMachineLog(item)
		machineLogs = append(machineLogs, machineLog)
	}
	return machineLogs
}
