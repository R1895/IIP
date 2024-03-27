package service

import (
	"fmt"
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
)

type CreateMachineStatusService struct {
	StatusType  uint `form:"status_type"  json:"status_type" binding:"required"`
	StatusName  string `form:"status_name"  json:"status_name" binding:"required"`
	Description string `form:"description" json:"description"`
}

type UpdateMachineStatusService struct {
	StatusType  uint `form:"status_type"  json:"status_type" binding:"required"`
	StatusName  string `form:"status_name"  json:"status_name" binding:"required"`
	Description string `form:"description" json:"description"`
}

type ListMachineStatusService struct {
	StatusName  string `form:"status_name"  json:"status_name"`
	Limit       int    `form:"limit" json:"limit"`
	Start       int    `form:"start" json:"start" binding:"required"`
}

type DeleteMachineStatusService struct {
}

type ShowMachineStatusService struct {
}

func (service *CreateMachineStatusService) Create() *serializer.Response {
	code := e.SUCCESS
	var machineStatus model.MachineStatus
	var count int64
	model.DB.Model(&model.MachineStatus{}).Where("status_name= ?", service.StatusName).First(&machineStatus).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	machineStatus = model.MachineStatus{
		StatusType: service.StatusType,
		StatusName:  service.StatusName,
		Description: service.Description,
	}

	if err := model.DB.Create(&machineStatus).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildMachineStatus(machineStatus),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListMachineStatusService) List() serializer.Response {
	var machineStatuss []model.MachineStatus
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.MachineStatus{})
	if service.StatusName != "" {
		tx = tx.Where(fmt.Sprintf("status_name like '%%%s%%'", service.StatusName))
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&machineStatuss)

	return serializer.BuildListResponse(serializer.BuildMachineStatuss(machineStatuss), uint(total))
}

func (service *ShowMachineStatusService) Show(id string) serializer.Response {
	var machineStatus model.MachineStatus
	code := e.SUCCESS
	err := model.DB.First(&machineStatus, id).Error
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
		Data:   serializer.BuildMachineStatus(machineStatus),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteMachineStatusService) Delete(id string) serializer.Response {
	var machineStatus model.MachineStatus
	code := e.SUCCESS
	err := model.DB.First(&machineStatus, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&machineStatus).Error
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

func (service *UpdateMachineStatusService) Update(id string) serializer.Response {
	var machineStatus model.MachineStatus
	model.DB.Model(model.MachineStatus{}).Where("id = ?", id).First(&machineStatus)

	if service.Description != "" {
		machineStatus.Description = service.Description
	}
	if service.StatusName != "" {
		machineStatus.StatusName = service.StatusName
	}

	code := e.SUCCESS
	err := model.DB.Save(&machineStatus).Error
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
