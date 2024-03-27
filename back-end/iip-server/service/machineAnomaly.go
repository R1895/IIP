package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"strconv"
	"time"
)

type CreateMachineAnomalyService2 struct{}

type CreateMachineAnomalyService struct {
	Time            time.Time `form:"time" json:"time" `
	MachineID       uint      `form:"machine_id" json:"machine_id" binding:"required"`
	MachineStatusID uint      `form:"machine_status_id" json:"machine_status_id"`
	//MachineStatus   MachineStatus `form:"machine_status" json:"machine_status"`
	Description string `form:"description" json:"description"`
}

type UpdateMachineAnomalyService struct {
	Time            time.Time `form:"time" json:"time" `
	MachineID       uint      `form:"machine_id" json:"machine_id"`
	MachineStatusID uint      `form:"machine_status_id" json:"machine_status_id"`
	//MachineStatus   MachineStatus `form:"machine_status" json:"machine_status"`
	Description string `form:"description" json:"description"`
}

type ListMachineAnomalyService struct {
	MachineID       uint `form:"machine_id" json:"machine_id"`
	MachineStatusID uint `form:"machine_status_id" json:"machine_status_id"`
	Limit           int  `form:"limit" json:"limit"`
	Start           int  `form:"start" json:"start"  binding:"required"`
}

type DeleteMachineAnomalyService struct {
}

type ShowMachineAnomalyService struct {
}

func (service *CreateMachineAnomalyService2) Create(mid string, id string) *serializer.Response {
	code := e.SUCCESS
	var machineAnomaly model.MachineAnomaly
	var count int64
	model.DB.Model(&model.MachineAnomaly{}).First(&machineAnomaly).Count(&count)

	t := time.Now()
	machine_id, _ := strconv.Atoi(mid)
	machine_status_id, _ := strconv.Atoi(id)

	machineAnomaly = model.MachineAnomaly{
		MachineID:       uint(machine_id),
		Time:            t,
		MachineStatusID: uint(machine_status_id),
	}

	if err := model.DB.Create(&machineAnomaly).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildMachineAnomaly(machineAnomaly),
		Msg:    e.GetMsg(code),
	}
}
func (service *CreateMachineAnomalyService) Create() *serializer.Response {
	code := e.SUCCESS
	var machineAnomaly model.MachineAnomaly
	service.Time = time.Now()
	machineAnomaly = model.MachineAnomaly{
		MachineID:       service.MachineID,
		Time:            service.Time,
		MachineStatusID: service.MachineStatusID,
		Description:     service.Description,
	}

	if err := model.DB.Create(&machineAnomaly).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildMachineAnomaly(machineAnomaly),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListMachineAnomalyService) List() serializer.Response {
	var machineAnomalys []model.MachineAnomaly
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.MachineAnomaly{})
	if service.MachineID != 0 {
		tx = tx.Where("machine_id = ?", service.MachineID)
	}
	if service.MachineStatusID != 0 {
		tx = tx.Where("machine_status_id =?", service.MachineStatusID)
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&machineAnomalys)

	return serializer.BuildListResponse(serializer.BuildMachineAnomalys(machineAnomalys), uint(total))
}

func (service *ShowMachineAnomalyService) Show(id string) serializer.Response {
	var machineAnomaly model.MachineAnomaly
	code := e.SUCCESS
	err := model.DB.First(&machineAnomaly, id).Error
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
		Data:   serializer.BuildMachineAnomaly(machineAnomaly),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteMachineAnomalyService) Delete(id string) serializer.Response {
	var machineAnomaly model.MachineAnomaly
	code := e.SUCCESS
	err := model.DB.First(&machineAnomaly, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&machineAnomaly).Error
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

func (service *UpdateMachineAnomalyService) Update(id string) serializer.Response {
	var machineAnomaly model.MachineAnomaly
	model.DB.Model(model.MachineAnomaly{}).Where("id = ?", id).First(&machineAnomaly)

	if service.MachineID != 0 {
		machineAnomaly.MachineID = service.MachineID
	}
	if service.MachineStatusID != 0 {
		machineAnomaly.MachineStatusID = service.MachineStatusID
	}

	if service.Description != "" {
		machineAnomaly.Description = service.Description
	}

	code := e.SUCCESS
	err := model.DB.Save(&machineAnomaly).Error
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
