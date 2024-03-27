package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
)

type CreateMachineWorkerService struct {
	MachineID uint      `form:"machine_id" json:"machine_id" binding:"required"`
	WorkerID  uint      `form:"worker_id"   json:"worker_id"  binding:"required"`
	StartDate time.Time `form:"start_date"  json:"start_date" binding:"required"`
	EndDate   time.Time `form:"end_date"    json:"end_date"`
}

type UpdateMachineWorkerService struct {
	MachineID uint      `form:"machine_id" json:"machine_id" `
	WorkerID  uint      `form:"worker_id"   json:"worker_id"  `
	EndDate   time.Time `form:"end_date"    json:"end_date"`
}

type ListMachineWorkerService struct {
	MachineID uint `form:"machine_id" json:"machine_id" `
	WorkerID  uint `form:"worker_id"   json:"worker_id"  `
	Limit     int  `form:"limit" json:"limit"`
	Start     int  `form:"start" json:"start" binding:"required"`
}

type DeleteMachineWorkerService struct {
}

type ShowMachineWorkerService struct {
}

func (service *CreateMachineWorkerService) Create() *serializer.Response {
	code := e.SUCCESS
	var machineWorker model.MachineWorker

	machineWorker = model.MachineWorker{
		MachineID: service.MachineID,
		WorkerID:  service.WorkerID,
		StartDate: service.StartDate,
		EndDate:   service.EndDate,
	}

	if err := model.DB.Create(&machineWorker).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildMachineWorker(machineWorker),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListMachineWorkerService) List() serializer.Response {
	var machineWorkers []model.MachineWorker
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.MachineWorker{})

	if service.MachineID != 0 {
		tx = tx.Where("machine_id = ?", service.MachineID)
	}
	if service.WorkerID != 0 {
		tx = tx.Where("worker_id = ?", service.WorkerID)
	}
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&machineWorkers)

	return serializer.BuildListResponse(serializer.BuildMachineWorkers(machineWorkers), uint(total))
}

func (service *ShowMachineWorkerService) Show(id string) serializer.Response {
	var machineWorker model.MachineWorker
	code := e.SUCCESS
	err := model.DB.First(&machineWorker, id).Error
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
		Data:   serializer.BuildMachineWorker(machineWorker),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteMachineWorkerService) Delete(id string) serializer.Response {
	var machineWorker model.MachineWorker
	code := e.SUCCESS
	err := model.DB.First(&machineWorker, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&machineWorker).Error
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

func (service *UpdateMachineWorkerService) Update(id string) serializer.Response {
	var machineWorker model.MachineWorker
	model.DB.Model(model.MachineWorker{}).Where("id = ?", id).First(&machineWorker)

	if service.MachineID != machineWorker.MachineID {
		machineWorker.MachineID = service.MachineID
	}
	if service.WorkerID != machineWorker.WorkerID {
		machineWorker.WorkerID = service.WorkerID
	}

	machineWorker.EndDate = time.Now()

	code := e.SUCCESS
	err := model.DB.Save(&machineWorker).Error
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
