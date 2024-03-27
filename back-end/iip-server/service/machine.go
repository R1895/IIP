package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"strconv"
	"time"
)

type CreateMachineService struct {
	WorkshopID      uint   `form:"workshop_id" json:"workshop_id" binding:"required"`
	MachineName     string `form:"machine_name" json:"machine_name" binding:"required"`
	ProcedureType   string `form:"procedure_type"  json:"procedure_type"`
	MachineStatusID uint   `form:"machine_status_id" json:"machine_status_id" binding:"required"`
}

type UpdateMachineService struct {
	WorkshopID      uint   `form:"workshop_id" json:"workshop_id"`
	MachineName     string `form:"machine_name" json:"machine_name"`
	MachineStatusID uint   `form:"machine_status_id" json:"machine_status_id"`
	ProcedureType   string `form:"procedure_type"  json:"procedure_type"`
	CurrentTaskID   uint   `form:"current_task_id" json:"current_task_id"`
}

type ListMachineService struct {
	WorkshopID      uint   `form:"workshop_id" json:"workshop_id"`
	MachineName     string `form:"machine_name" json:"machine_name"`
	MachineStatusID uint   `form:"machine_status_id" json:"machine_status_id"`
	CurrentTaskID   uint   `form:"current_task_id" json:"current_task_id"`
	ProcedureType   string `form:"procedure_type"  json:"procedure_type"`
	Limit           int    `form:"limit" json:"limit"`
	Start           int    `form:"start" json:"start" binding:"required"`
}

type ModifyMachineStatusService struct {
	// MachineID     string `form:"machine_id" json:"machine_id"`
	// MachineStatusID string   `form:"machine_status_id" json:"machine_status_id"`
}
func (service *ModifyMachineStatusService) ModifyMachineStatus(mid string,sid string  ) serializer.Response {
	var machine model.Machine
	model.DB.Model(model.Machine{}).Where("id = ?", mid).First(&machine)
	machine.PurchaseDate = time.Now()
	machine.WarrantyUntil = time.Now()
	s,_ := strconv.Atoi(sid)
	machine.MachineStatusID = uint(s)
	code := e.SUCCESS
	err := model.DB.Save(&machine).Error
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

type DeleteMachineService struct {
}

type ShowMachineService struct {
}

func (service *CreateMachineService) Create() *serializer.Response {
	code := e.SUCCESS
	var machine model.Machine
	var count int64

	model.DB.Model(&model.Machine{}).Where("machine_name=?", service.MachineName).First(&machine).Count(&count)
	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// TODO 这行代码 干什么 count = 0 查询必须存在该id
	var machineStatus model.MachineStatus
	model.DB.Model(&model.MachineStatus{}).Where("machine_status_id=?", service.MachineStatusID).First(&machineStatus).Count(&count)
	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	machine = model.Machine{
		WorkshopID:      service.WorkshopID,
		MachineName:     service.MachineName,
		MachineStatusID: service.MachineStatusID,
		ProcedureType:   service.ProcedureType,
	}

	if err := model.DB.Create(&machine).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildMachine(machine),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListMachineService) List() serializer.Response {
	var machines []model.Machine
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 预加载和分页
	tx := model.DB.Model(model.Machine{}).Preload("MachineWorkers").Preload("Procedures")
	if service.MachineName != "" {
		tx.Where("machine_name = like '%%%s%%'", service.MachineName)
	}
	if service.ProcedureType != "" {
		tx.Where("procedure_type = ?", service.ProcedureType)
	}
	if service.MachineStatusID != 0 {
		tx.Where("machine_status_id = ?", service.MachineStatusID)
	}
	if service.WorkshopID != 0 {
		tx.Where("workshop_id = ?", service.WorkshopID)
	}
	if service.CurrentTaskID != 0 {
		tx.Where("current_task_id = ?", service.CurrentTaskID)
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&machines)

	return serializer.BuildListResponse(
		serializer.BuildMachines(machines), uint(total))
}

func (service *ShowMachineService) Show(id string) serializer.Response {
	var machine model.Machine
	code := e.SUCCESS
	err := model.DB.Model(model.Machine{}).Preload("MachineWorkers").Preload("Procedures").Preload("MachineAnomalys").Preload("MachineLogs").First(&machine, id).Error
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
		Data:   serializer.BuildMachine(machine),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteMachineService) Delete(id string) serializer.Response {
	var machine model.Machine
	code := e.SUCCESS
	err := model.DB.First(&machine, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&machine).Error
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

func (service *UpdateMachineService) Update(id string) serializer.Response {
	var machine model.Machine
	model.DB.Model(model.Machine{}).Where("id = ?", id).First(&machine)
	machine.WorkshopID = service.WorkshopID
	if service.MachineName != "" {
		machine.MachineName = service.MachineName
	}

	if service.ProcedureType != "" {
		machine.ProcedureType = service.ProcedureType
	}
	if service.MachineStatusID != machine.MachineStatusID {
		machine.MachineStatusID = service.MachineStatusID
	}
	machine.PurchaseDate = time.Now()
	machine.WarrantyUntil = time.Now()

	code := e.SUCCESS
	err := model.DB.Save(&machine).Error
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

// TODO 分配工人
type MachineWorker struct{}

func (service *MachineWorker) AssignWorker() *serializer.Response {
	code := e.SUCCESS

	return &serializer.Response{
		Status: code,

		Msg: e.GetMsg(code),
	}
}

// TODO 分配工序
type MachineTask struct{}

func (service *MachineWorker) AssignTask() *serializer.Response {
	code := e.SUCCESS

	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
