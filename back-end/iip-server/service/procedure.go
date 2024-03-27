package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
)

type CreateProcedureService struct {
	ProcedureType     string    `form:"procedure_type" json:"procedure_type"`
	ProcedureName     string    `form:"procedure_name" json:"procedure_name" binding:"required"`
	TaskID            uint      `form:"task_id" json:"task_id" `
	MachineID         uint      `form:"machine_id" json:"machine_id" binding:"required" `
	WorkpiecesNum     uint      `form:"workpieces_num" json:"workpieces_num"`
	ProcessedNum      uint      `form:"processed_num" json:"processed_num"`
	MachineSequence   uint      `form:"machine_sequence" json:"machine_sequence"`
	Sequence          uint      `form:"sequence" json:"sequence"`
	StartDate         time.Time `form:"start_date" json:"start_date"`
	IsFinished        uint      `form:"is_finished" json:"is_finished"`
	ProcedureStatusID uint      `form:"procedure_status_id" json:"procedure_status_id"`
}

type UpdateProcedureService struct {
	ProcedureType     string    `form:"procedure_type" json:"procedure_type"`
	ProcedureName     string    `form:"procedure_name" json:"procedure_name" binding:"required"`
	TaskID            uint      `form:"task_id" json:"task_id"`
	MachineID         uint      `form:"machine_id" json:"machine_id" binding:"required"`
	WorkpiecesNum     uint      `form:"workpieces_num" json:"workpieces_num"`
	ProcessedNum      uint      `form:"processed_num" json:"processed_num"`
	MachineSequence   uint      `form:"machine_sequence" json:"machine_sequence"`
	Sequence          uint      `form:"sequence" json:"sequence"`
	StartDate         time.Time `form:"start_date" json:"start_date"`
	IsFinished        uint      `form:"is_finished" json:"is_finished"`
	ProcedureStatusID uint      `form:"procedure_status_id" json:"procedure_status_id"`
}

type ListProcedureService struct {
	ProcedureType     string `form:"procedure_type" json:"procedure_type"`
	TaskID            uint   `form:"task_id" json:"task_id"`
	MachineID         uint   `form:"machine_id" json:"machine_id"`
	IsFinished        uint   `form:"is_finished" json:"is_finished"`
	ProcedureStatusID uint   `form:"procedure_status_id" json:"procedure_status_id"`
	Limit             int    `form:"limit" json:"limit"`
	Start             int    `form:"start" json:"start" binding:"required"`
}

type DeleteProcedureService struct {
}

type ShowProcedureService struct {
}

func (service *CreateProcedureService) Create() *serializer.Response {
	code := e.SUCCESS
	var procedure model.Procedure
	var count int64
	model.DB.Model(&model.Procedure{}).Where("procedure_name=?", service.ProcedureName).First(&procedure).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	procedure = model.Procedure{
		ProcedureName:     service.ProcedureName,
		ProcedureType:     service.ProcedureType,
		TaskID:            service.TaskID,
		MachineID:         service.MachineID,
		WorkpiecesNum:     service.WorkpiecesNum,
		ProcessedNum:      service.ProcessedNum,
		Sequence:          service.Sequence,
		MachineSequence:   service.MachineSequence,
		StartDate:         service.StartDate,
		IsFinished:        service.IsFinished,
		ProcedureStatusID: service.ProcedureStatusID,
	}

	if err := model.DB.Create(&procedure).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildProcedure(procedure),
		Msg:    e.GetMsg(code),
	}
}

// TODO 搜索条件
func (service *ListProcedureService) List() serializer.Response {
	var procedures []model.Procedure
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.Procedure{}).Preload("ProcedureItems")
	if service.TaskID != 0 {
		tx = tx.Where("task_id =?", service.TaskID)
	}
	if service.MachineID != 0 {
		tx = tx.Where("machine_id =?", service.MachineID)
	}
	if service.ProcedureType != "" {
		tx = tx.Where("procedure_type =?", service.ProcedureType)
	}
	if service.ProcedureStatusID != 0 {
		tx = tx.Where("procedure_status_id =?", service.ProcedureStatusID)
	}
	if service.IsFinished != 0 {
		tx = tx.Where("is_finished =?", service.IsFinished)
	}
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&procedures)

	return serializer.BuildListResponse(serializer.BuildProcedures(procedures), uint(total))
}

func (service *ShowProcedureService) Show(id string) serializer.Response {
	var procedure model.Procedure
	code := e.SUCCESS
	err := model.DB.Model(model.Procedure{}).Preload("ProcedureItems").First(&procedure, id).Error
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
		Data:   serializer.BuildProcedure(procedure),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteProcedureService) Delete(id string) serializer.Response {
	var procedure model.Procedure
	code := e.SUCCESS
	err := model.DB.First(&procedure, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&procedure).Error
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

// TODO 更新条件
func (service *UpdateProcedureService) Update(id string) serializer.Response {
	var procedure model.Procedure
	model.DB.Model(model.Procedure{}).Where("id=?", id).First(&procedure)
	if procedure.TaskID != 0 {
		procedure.TaskID = service.TaskID
	}
	if procedure.MachineID != 0 {
		procedure.MachineID = service.MachineID
	}
	if procedure.WorkpiecesNum != 0 {
		procedure.WorkpiecesNum = service.WorkpiecesNum

	}
	if procedure.ProcessedNum != 0 {
		procedure.ProcessedNum = service.ProcessedNum
	}
	if procedure.Sequence != 0 {
		procedure.Sequence = service.Sequence
	}
	if procedure.StartDate != service.StartDate {
		procedure.StartDate = service.StartDate
	}
	if procedure.IsFinished != service.IsFinished {
		procedure.IsFinished = service.IsFinished
	}

	code := e.SUCCESS
	err := model.DB.Save(&procedure).Error
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
