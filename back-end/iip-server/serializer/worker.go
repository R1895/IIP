package serializer

import "iip/model"

type Worker struct {
	ID                uint               `json:"id"`
	UserID            uint               `json:"user_id"`
	WorkerTypeID      uint               `json:"worker_type_id"`
	WorkerType        string             `json:"worker_type"`
	WorkerName        string             `json:"worker_name"`
	SkillLevel        int                `json:"skill_level"`
	Workshops         []Workshop         `json:"workshops"`
	MachineWorkers    []MachineWorker    `json:"machine_workers"`
	WorkerAttendances []WorkerAttendance `json:"worker_attendances"`
}

// BuildWorker 序列化角色
func BuildWorker(worker model.Worker) Worker {
	return Worker{
		ID:                worker.ID,
		UserID:            worker.UserID,
		WorkerTypeID:      worker.WorkerTypeID,
		WorkerType:        worker.WorkerType,
		WorkerName:        worker.WorkerName,
		SkillLevel:        worker.SkillLevel,
		Workshops:         BuildWorkshops(worker.Workshops),
		MachineWorkers:    BuildMachineWorkers(worker.MachineWorkers),
		WorkerAttendances: BuildWorkerAttendances(worker.WorkerAttendances),
	}
}

func BuildWorkers(items []model.Worker) (workers []Worker) {
	for _, item := range items {
		worker := BuildWorker(item)
		workers = append(workers, worker)
	}
	return
}
