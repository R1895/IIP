package serializer

import (
	"iip/model"
	"time"
)

type Task struct {
	ID            uint        `json:"id"`
	WorkshopID    uint        `json:"workshop_id"`
	TaskName      string      `json:"task_name"`
	Description   string      `json:"description"`
	StartDate     time.Time   `json:"start_date"`
	EffectiveTime time.Time   `json:"effective_time"`
	TaskStatusID  uint        `json:"task_status_id"`
	IsFinished    uint        `json:"is_finished"`
	Procedures    []Procedure `json:"procedures"`
	TaskTypeID    uint        `json:"task_type_id"`
}

func BuildTask(task model.Task) Task {
	return Task{
		ID:            task.ID,
		WorkshopID:    task.WorkshopID,
		TaskName:      task.TaskName,
		Description:   task.Description,
		StartDate:     task.StartDate,
		EffectiveTime: task.EffectiveTime,
		TaskStatusID:  task.TaskStatusID,
		IsFinished:    task.IsFinished,
		Procedures:    BuildProcedures(task.Procedures),
		TaskTypeID:    task.TaskTypeID,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
