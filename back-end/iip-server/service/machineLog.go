package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
)

type CreateMachineLogService struct {
	Time            time.Time `form:"time" json:"time" binding:"required" `
	MachineID       uint      `form:"machine_id" json:"machine_id" binding:"required"`
	MachineStatusID uint      `form:"machine_status_id" json:"machine_status_id"`
	Description     string    `form:"description" json:"description"`
}

type UpdateMachineLogService struct {
	Time            time.Time `form:"time" json:"time" binding:"required" `
	MachineID       uint      `form:"machine_id" json:"machine_id" binding:"required" `
	MachineStatusID uint      `form:"machine_status_id" json:"machine_status_id"`
	Description     string    `form:"description" json:"description"`
}

type ListMachineLogService struct {
	MachineID       uint `form:"machine_id" json:"machine_id"`
	MachineStatusID uint `form:"machine_status_id" json:"machine_status_id"`
	Limit           int  `form:"limit" json:"limit"`
	Start           int  `form:"start" json:"start" binding:"required" `
}

type DeleteMachineLogService struct {
}

type ShowMachineLogService struct {
}


func (service *CreateMachineLogService) Create() *serializer.Response {
	code := e.SUCCESS
	var machineLog model.MachineLog

	machineLog = model.MachineLog{
		MachineID:       service.MachineID,
		Time:            service.Time,
		MachineStatusID: service.MachineStatusID,
		Description:     service.Description,
	}

	if err := model.DB.Create(&machineLog).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildMachineLog(machineLog),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListMachineLogService) List() serializer.Response {
	var machineLogs []model.MachineLog
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.MachineLog{})
	if service.MachineID != 0 {
		tx = tx.Where("machine_id = ?", service.MachineID)
	}
	if service.MachineStatusID != 0 {
		tx = tx.Where("machine_status_id =?", service.MachineStatusID)
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&machineLogs)

	return serializer.BuildListResponse(serializer.BuildMachineLogs(machineLogs), uint(total))
}

func (service *ShowMachineLogService) Show(id string) serializer.Response {
	var machineLog model.MachineLog
	code := e.SUCCESS
	err := model.DB.First(&machineLog, id).Error
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
		Data:   serializer.BuildMachineLog(machineLog),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteMachineLogService) Delete(id string) serializer.Response {
	var machineLog model.MachineLog
	code := e.SUCCESS
	err := model.DB.First(&machineLog, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&machineLog).Error
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

func (service *UpdateMachineLogService) Update(id string) serializer.Response {
	var machineLog model.MachineLog
	model.DB.Model(model.MachineLog{}).Where("id = ?", id).First(&machineLog)

	if service.MachineID != 0 {
		machineLog.MachineID = service.MachineID
	}
	if service.MachineStatusID != 0 {
		machineLog.MachineStatusID = service.MachineStatusID
	}

	if service.Description != "" {
		machineLog.Description = service.Description
	}

	code := e.SUCCESS
	err := model.DB.Save(&machineLog).Error
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

