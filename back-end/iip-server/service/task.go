package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
)

type CreateTaskService struct {
	WorkshopID    uint      `form:"workshop_id" json:"workshop_id"  binding:"required"  `
	TaskName      string    `form:"task_name" json:"task_name"     binding:"required"  `
	Description   string    `form:"description" json:"description"`
	StartDate     time.Time `form:"start_date" json:"start_date"`
	EffectiveTime time.Time `form:"effective_time" json:"effective_time"`
	TaskStatusID  uint      `form:"task_status_id" json:"task_status_id" binding:"required"  `
	IsFinished    uint      `form:"is_finished" json:"is_finished"`
	TaskTypeID    uint      `form:"task_type_id" json:"task_type_id" binding:"required"  `
}

type UpdateTaskService struct {
	WorkshopID    uint      `form:"workshop_id" json:"workshop_id" binding:"required"  `
	TaskName      string    `form:"task_name" json:"task_name"     binding:"required"`
	Description   string    `form:"description" json:"description"`
	EffectiveTime time.Time `form:"effective_time" json:"effective_time"`
	TaskStatusID  uint      `form:"task_status_id" json:"task_status_id" binding:"required"  `
	IsFinished    uint      `form:"is_finished" json:"is_finished"`
	TaskTypeID    uint      `form:"task_type_id" json:"task_type_id" binding:"required"  `
}

type ListTaskService struct {
	WorkshopID  uint   `form:"workshop_id" json:"workshop_id"  `
	TaskName    string `form:"task_name" json:"task_name" `
	TaskTypeID  uint   `form:"task_type_id"  json:"task_type_id"`
	StartDate     time.Time `form:"start_date" json:"start_date"`
	EffectiveTime time.Time `form:"effective_time" json:"effective_time"`
	TaskStatusID  uint      `form:"task_status_id" json:"task_status_id" `
	IsFinished  uint   `form:"is_finished" json:"is_finished"`
	Limit       int    `form:"limit" json:"limit"`
	Start       int    `form:"start" json:"start" binding:"required"  `
}

type DeleteTaskService struct {
}

type ShowTaskService struct {
}

func (service *CreateTaskService) Create() *serializer.Response {
	code := e.SUCCESS
	var task model.Task

	// 任务名可以重复

	// var count int64
	// model.DB.Model(&model.Task{}).Where("task_name=?", service.TaskName).First(&task).Count(&count)

	// if count == 1 {
	// 	code = e.ErrorExist
	// 	return &serializer.Response{
	// 		Status: code,
	// 		Msg:    e.GetMsg(code),
	// 	}
	// }
	task = model.Task{
		TaskName:      service.TaskName,
		WorkshopID:    service.WorkshopID,
		Description:   service.Description,
		StartDate:     time.Now(),
		EffectiveTime: time.Now(),
		TaskStatusID:  service.TaskStatusID,
		IsFinished:    service.IsFinished,
		TaskTypeID:    service.TaskTypeID,
	}

	if err := model.DB.Create(&task).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListTaskService) List() serializer.Response {
	var tasks []model.Task
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.Task{}).Preload("Procedures")

	if service.WorkshopID != 0 {
		tx = tx.Where("workshop_id =?", service.WorkshopID)
	}
	if service.TaskName != "" {
		tx = tx.Where("skill_level like '%%%s%%'", service.TaskName)
	}
	if service.TaskStatusID != 0 {
		tx = tx.Where("skill_level =?", service.TaskStatusID)
	}
	if service.TaskTypeID != 0 {
		tx = tx.Where("skill_level =?", service.TaskTypeID)
	}
	if service.IsFinished != 0 {
		tx = tx.Where("skill_level =?", service.IsFinished)
	}
	//TODO
	// if service.StartDate != 0 {
	// 	tx = tx.Where("skill_level =?", service.SkillLevel)
	// }
	// if service.EffectiveTime != 0 {
	// 	tx = tx.Where("skill_level =?", service.SkillLevel)
	// }
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&tasks)

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))
}

func (service *ShowTaskService) Show(id string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	err := model.DB.Model(model.Task{}).Preload("Procedures").First(&task, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	err := model.DB.First(&task, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&task).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *UpdateTaskService) Update(id string) serializer.Response {
	var task model.Task
	var code = e.SUCCESS
	var errs = model.DB.Model(model.Task{}).Where("id = ?", id).First(&task).Error
	if errs != nil {
		util.LogrusObj.Info(errs)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errs.Error(),
		}
	}
	if service.WorkshopID != 0 {
		task.WorkshopID = service.WorkshopID
	}
	if service.TaskTypeID != 0 {
		task.TaskTypeID = service.TaskTypeID
	}
	if service.Description != "" {
		task.Description = service.Description
	}

	if service.TaskName != "" {
		task.TaskName = service.TaskName
	}

	if service.TaskStatusID != 0 {
		task.TaskStatusID = service.TaskStatusID
	}
	if service.IsFinished != 0 {
		task.IsFinished = service.IsFinished
	}

	task.EffectiveTime = time.Now() //
	if service.IsFinished != 0 {
		task.IsFinished = service.IsFinished
	}

	err := model.DB.Save(&task).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "修改成功",
	}
}
