package serializer

import (
	"iip/model"
	"time"
)

type MachineWorker struct {
	ID         uint      `json:"id"`
	MachineID  uint      `json:"machine_id"`
	WorkerID   uint      `json:"worker_id"`
	WorkerName string    `json:"worker_name"`
	WorkerType string    `json:"worker_type"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

func BuildMachineWorker(machineWorker model.MachineWorker) MachineWorker {
	var worker model.Worker

	model.DB.First(&worker, machineWorker.WorkerID)
	return MachineWorker{
		ID:         machineWorker.ID,
		MachineID:  machineWorker.MachineID,
		WorkerID:   machineWorker.WorkerID,
		WorkerName: worker.WorkerName,
		WorkerType: worker.WorkerType,
		StartDate:  machineWorker.StartDate,
		EndDate:    machineWorker.EndDate,
	}
}

func BuildMachineWorkers(items []model.MachineWorker) (machineWorkers []MachineWorker) {
	for _, item := range items {
		machineWorker := BuildMachineWorker(item)
		machineWorkers = append(machineWorkers, machineWorker)
	}
	return machineWorkers
}
