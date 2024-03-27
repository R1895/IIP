package serializer

import "iip/model"

type TaskStatus struct {
	ID           uint      `json:"id"`
	TaskStatus string      `json:"task_status"`
	Tasks      []Task      `json:"tasks"`
	Description  string    `json:"description"`
}

func BuildTaskStatus(taskStatus model.TaskStatus) TaskStatus {
	return TaskStatus{
		ID:           taskStatus.ID,
		TaskStatus: taskStatus.TaskStatus,
		Tasks:      BuildTasks(taskStatus.Tasks),
		Description:   taskStatus.Description,
	}
}

func BuildTaskStatuss(items []model.TaskStatus) (taskStatuss []TaskStatus) {
	for _, item := range items {
		taskStatus := BuildTaskStatus(item)
		taskStatuss = append(taskStatuss, taskStatus)
	}
	return taskStatuss
}
