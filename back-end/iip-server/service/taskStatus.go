package service

import (

	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
)

type CreateTaskStatusService struct {
	TaskStatus  string `form:"task_status" json:"task_status" binding:"required"`
	Description string `form:"description" json:"description"`
}

type UpdateTaskStatusService struct {
	TaskStatus  string `form:"task_status" json:"task_status" `
	Description string `form:"description" json:"description"`
}

type ListTaskStatusService struct {
	TaskStatus  string `form:"task_status" json:"task_status" `
	Limit       int    `form:"limit" json:"limit" `
	Start       int    `form:"start" json:"start" binding:"required"`
}

type DeleteTaskStatusService struct {
}

type ShowTaskStatusService struct {
}

func (service *CreateTaskStatusService) Create() *serializer.Response {
	code := e.SUCCESS
	var taskStatus model.TaskStatus
	var count int64
	model.DB.Model(&model.TaskStatus{}).Where("task_status=?", service.TaskStatus).First(&taskStatus).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	taskStatus = model.TaskStatus{
		TaskStatus:  service.TaskStatus,
		Description: service.Description,
	}

	if err := model.DB.Create(&taskStatus).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildTaskStatus(taskStatus),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListTaskStatusService) List() serializer.Response {
	var taskStatuss []model.TaskStatus
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 预加载和分页
	tx := model.DB.Model(model.TaskStatus{}).Preload("Tasks")
	if service.TaskStatus != "" {
		tx = tx.Where("task_status =?", service.TaskStatus)
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&taskStatuss)

	return serializer.BuildListResponse(
		serializer.BuildTaskStatuss(taskStatuss), uint(total))
}

// 展示详情
func (service *ShowTaskStatusService) Show(id string) serializer.Response {
	var taskStatus model.TaskStatus
	code := e.SUCCESS
	err := model.DB.Model(model.TaskStatus{}).Preload("Tasks").First(&taskStatus, id).Error
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
		Data:   serializer.BuildTaskStatus(taskStatus),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteTaskStatusService) Delete(id string) serializer.Response {
	var taskStatus model.TaskStatus
	code := e.SUCCESS
	err := model.DB.First(&taskStatus, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&taskStatus).Error
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

func (service *UpdateTaskStatusService) Update(id string) serializer.Response {
	var taskStatus model.TaskStatus
	model.DB.Model(model.TaskStatus{}).Where("id = ?", id).First(&taskStatus)
	if service.TaskStatus != "" {
		taskStatus.TaskStatus = service.TaskStatus
	}
	if service.Description != "" {
		taskStatus.Description = service.Description
	}

	code := e.SUCCESS
	err := model.DB.Save(&taskStatus).Error
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
