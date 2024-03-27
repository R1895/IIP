package serializer

import "iip/model"

type WorkerType struct {
	ID           uint      `json:"id"`
	WorkerType string    `json:"worker_type"`
	Workers      []Worker  `json:"workers"`
	Description  string    `json:"description"`
}

func BuildWorkerType(workerType model.WorkerType) WorkerType {
	return WorkerType{
		ID:           workerType.ID,
		WorkerType: workerType.WorkerType,
		Workers:      BuildWorkers(workerType.Workers),
		Description:   workerType.Description,
	}
}

func BuildWorkerTypes(items []model.WorkerType) (workerTypes []WorkerType) {
	for _, item := range items {
		workerType := BuildWorkerType(item)
		workerTypes = append(workerTypes, workerType)
	}
	return workerTypes
}
