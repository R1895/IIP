package serializer

import (
	"iip/model"
)

type MachineStatus struct {
	ID          uint      `json:"id"`
	StatusType  uint      `json:"status_type"`
	StatusName  string    `json:"status_name"`
	Machines    []Machine `json:"machines"`
	Description string    `json:"description"`
}

func BuildMachineStatus(machineStatus model.MachineStatus) MachineStatus {
	return MachineStatus{
		ID:          machineStatus.ID,
		StatusType:  machineStatus.StatusType,
		StatusName:  machineStatus.StatusName,
		Machines:    BuildMachines(machineStatus.Machines),
		Description: machineStatus.Description,
	}
}

func BuildMachineStatuss(items []model.MachineStatus) (machineStatuss []MachineStatus) {
	for _, item := range items {
		machineStatus := BuildMachineStatus(item)
		machineStatuss = append(machineStatuss, machineStatus)
	}
	return machineStatuss
}
