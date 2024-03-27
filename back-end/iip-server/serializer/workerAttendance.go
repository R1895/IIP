package serializer

import (
	"iip/model"
	"time"
)

type WorkerAttendance struct {
	ID            uint      `json:"attendance_id"`
	WorkerID      uint      `json:"worker_id"`
	Date          time.Time `json:"date"`
	ArrivalTime   time.Time `json:"arrival_time"`
	DepartureTime time.Time `json:"departure_time"`
	Type          string    `json:"type"`
}

type StatAttendance struct{
	WorkerID      uint      
	TotalTime     string
	WorkerName  string   `json:"worker_name"`
}

func BuildWorkerAttendance(workerAttendance model.WorkerAttendance) WorkerAttendance {
	return WorkerAttendance{
		ID:            workerAttendance.ID,
		WorkerID:      workerAttendance.WorkerID,
		Date:          workerAttendance.Date,
		ArrivalTime:   workerAttendance.ArrivalTime,
		DepartureTime: workerAttendance.DepartureTime,
		Type:          workerAttendance.Type,
	}
}

func BuildWorkerAttendances(items []model.WorkerAttendance) (workerAttendances []WorkerAttendance) {
	for _, item := range items {
		workerAttendance := BuildWorkerAttendance(item)
		workerAttendances = append(workerAttendances, workerAttendance)
	}
	return workerAttendances
}

func BuildStatAttendance (statattendance model.StatAttendance) StatAttendance{
	return StatAttendance{
		WorkerID: statattendance.WorkerID,
		TotalTime: statattendance.TotalTime,
		WorkerName: statattendance.WorkerName,
	}
}

func BuildStatAttendances(items []model.StatAttendance) (statAttendances []StatAttendance) {
	for _, item := range items {
		statAttendance := BuildStatAttendance(item)
		statAttendances = append(statAttendances, statAttendance)
	}
	return statAttendances
}
