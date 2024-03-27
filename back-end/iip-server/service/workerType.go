package service

import (
	"fmt"
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
)

type CreateWorkerTypeService struct {
	WorkerType  string `form:"worker_type" json:"worker_type" binding:"required"`
	Description string `form:"description" json:"description"`
}

type UpdateWorkerTypeService struct {
	WorkerType  string `form:"worker_type" json:"worker_type" binding:"required"`
	Description string `form:"description" json:"description"`
}

type ListWorkerTypeService struct {
	WorkerType  string `form:"worker_type" json:"worker_type" `
	Limit       int    `form:"limit" json:"limit" `
	Start       int    `form:"start" json:"start" binding:"required"`
}

type DeleteWorkerTypeService struct {
}

type ShowWorkerTypeService struct {
}

func (service *CreateWorkerTypeService) Create() *serializer.Response {
	code := e.SUCCESS
	var workerType model.WorkerType
	var count int64
	model.DB.Model(&model.WorkerType{}).Where("worker_type=?", service.WorkerType).First(&workerType).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	workerType = model.WorkerType{
		WorkerType:  service.WorkerType,
		Description: service.Description,
	}

	if err := model.DB.Create(&workerType).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildWorkerType(workerType),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListWorkerTypeService) List() serializer.Response {
	var workerTypes []model.WorkerType
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	tx := model.DB.Model(model.WorkerType{}).Preload("Workers")
	if service.WorkerType != "" {
		tx = tx.Where(fmt.Sprintf("worker_type like '%%%s%%'", service.WorkerType))
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&workerTypes)

	return serializer.BuildListResponse(
		serializer.BuildWorkerTypes(workerTypes), uint(total))
}

// 展示详情
func (service *ShowWorkerTypeService) Show(id string) serializer.Response {
	var workerType model.WorkerType
	code := e.SUCCESS
	err := model.DB.Model(model.WorkerType{}).Preload("Workers").First(&workerType, id).Error
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
		Data:   serializer.BuildWorkerType(workerType),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteWorkerTypeService) Delete(id string) serializer.Response {
	var workerType model.WorkerType
	code := e.SUCCESS
	err := model.DB.First(&workerType, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&workerType).Error
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

func (service *UpdateWorkerTypeService) Update(id string) serializer.Response {
	var workerType model.WorkerType
	model.DB.Model(model.WorkerType{}).Where("id = ?", id).First(&workerType)

	// 检查字段是否更新
	if workerType.WorkerType == service.WorkerType && workerType.Description == service.Description {
		code := e.ErrorNotUpdate
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if service.WorkerType != "" {
		workerType.WorkerType = service.WorkerType
	}
	if service.Description != "" {
		workerType.Description = service.Description
	}

	code := e.SUCCESS
	err := model.DB.Save(&workerType).Error
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
