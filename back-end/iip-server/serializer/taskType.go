package serializer

import "iip/model"

type TaskType struct {
	ID          uint   `json:"task_type_id"`
	TaskType    string `json:"task_type"`
	Tasks       []Task `json:"tasks"`
	Description string `json:"description"`
}

func BuildTaskType(taskType model.TaskType) TaskType {
	return TaskType{
		ID:          taskType.ID,
		TaskType:    taskType.TaskType,
		Tasks:       BuildTasks(taskType.Tasks),
		Description: taskType.Description,
	}
}

func BuildTaskTypes(items []model.TaskType) (taskTypes []TaskType) {
	for _, item := range items {
		taskType := BuildTaskType(item)
		taskTypes = append(taskTypes, taskType)
	}
	return taskTypes
}
