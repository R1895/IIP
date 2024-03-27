package serializer

import (
	"iip/model"
	"time"
)

type Machine struct {
	ID              uint             `json:"id"`
	WorkshopID      uint             `json:"workshop_id"`
	ProcedureType   string             `json:"procedure_type"`
	MachineName     string           `json:"machine_name"`
	MachineStatusID uint             `json:"status_id"`
	CurrentTaskID   uint             `json:"current_task_id"`
	PurchaseDate    time.Time        `json:"purchase_date"`
	WarrantyUntil   time.Time        `json:"warranty_until"`
	Workers         []MachineWorker  `json:"machine_workers"`
	MachineAnomalys []MachineAnomaly `json:"machine_anomalys"`
	MachineLogs     []MachineLog     `json:"machine_logs"`
	Procedures      []Procedure      `json:"procedures"`
}

func BuildMachine(machine model.Machine) Machine {
	return Machine{
		ID:              machine.ID,
		WorkshopID:      machine.WorkshopID,
		ProcedureType:   machine.ProcedureType,
		MachineName:     machine.MachineName,
		MachineStatusID: machine.MachineStatusID,
		CurrentTaskID:   machine.CurrentTaskID,
		PurchaseDate:    machine.PurchaseDate,
		WarrantyUntil:   machine.WarrantyUntil,
		Workers:         BuildMachineWorkers(machine.MachineWorkers),
		MachineAnomalys: BuildMachineAnomalys(machine.MachineAnomalys),
		MachineLogs:     BuildMachineLogs(machine.MachineLogs),
		Procedures:      BuildProcedures(machine.Procedures),
	}
}

func BuildMachines(items []model.Machine) (machines []Machine) {
	for _, item := range items {
		machine := BuildMachine(item)
		machines = append(machines, machine)
	}
	return machines
}

