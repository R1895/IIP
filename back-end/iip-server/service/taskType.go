package service

import (
	"fmt"
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
)

type CreateTaskTypeService struct {
	TaskType    string `form:"task_type" json:"task_type" binding:"required"`
	Description string `form:"description" json:"description"`
}

type UpdateTaskTypeService struct {
	TaskType    string `form:"task_type" json:"task_type" binding:"required"`
	Description string `form:"description" json:"description"`
}

type ListTaskTypeService struct {
	TaskType    string `form:"task_type" json:"task_type" `
	Limit       int    `form:"limit" json:"limit" `
	Start       int    `form:"start" json:"start" binding:"required"`
}

type DeleteTaskTypeService struct {
}

type ShowTaskTypeService struct {
}

func (service *CreateTaskTypeService) Create() *serializer.Response {
	code := e.SUCCESS
	var taskType model.TaskType
	var count int64
	model.DB.Model(&model.TaskType{}).Where("task_type=?", service.TaskType).First(&taskType).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	taskType = model.TaskType{
		TaskType:    service.TaskType,
		Description: service.Description,
	}

	if err := model.DB.Create(&taskType).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildTaskType(taskType),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListTaskTypeService) List() serializer.Response {
	var taskTypes []model.TaskType
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 预加载和分页
	tx := model.DB.Model(model.TaskType{}).Preload("Tasks")
	if service.TaskType != "" {
		tx = tx.Where(fmt.Sprintf("task_type like '%%%s%%'", service.TaskType))
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&taskTypes)

	return serializer.BuildListResponse(
		serializer.BuildTaskTypes(taskTypes), uint(total))
}

// 展示详情
func (service *ShowTaskTypeService) Show(id string) serializer.Response {
	var taskType model.TaskType
	code := e.SUCCESS
	err := model.DB.Model(model.TaskType{}).Preload("Tasks").First(&taskType, id).Error
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
		Data:   serializer.BuildTaskType(taskType),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteTaskTypeService) Delete(id string) serializer.Response {
	var taskType model.TaskType
	code := e.SUCCESS
	err := model.DB.First(&taskType, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&taskType).Error
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

func (service *UpdateTaskTypeService) Update(id string) serializer.Response {
	var taskType model.TaskType
	model.DB.Model(model.TaskType{}).Where("id = ?", id).First(&taskType)
	if service.TaskType != "" {
		taskType.TaskType = service.TaskType
	}
	if service.Description != "" {
		taskType.Description = service.Description
	}

	code := e.SUCCESS
	err := model.DB.Save(&taskType).Error
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
