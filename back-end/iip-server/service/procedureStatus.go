package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"fmt"
)

type CreateProcedureStatusService struct {
	ProcedureStatus string `form:"procedure_status" json:"procedure_status" binding:"required"`
	Description     string `form:"description" json:"description"`
}

type UpdateProcedureStatusService struct {
	ProcedureStatus string `form:"procedure_status" json:"procedure_status" binding:"required"`
	Description     string `form:"description" json:"description"`
}

type ListProcedureStatusService struct {
	ProcedureStatus string `form:"procedure_status" json:"procedure_status" `
	Limit           int    `form:"limit" json:"limit"`
	Start           int    `form:"start" json:"start" binding:"required"`
}

type DeleteProcedureStatusService struct {
}

type ShowProcedureStatusService struct {
}

func (service *CreateProcedureStatusService) Create() *serializer.Response {
	code := e.SUCCESS
	var procedureStatus model.ProcedureStatus
	var count int64
	model.DB.Model(&model.ProcedureStatus{}).Where("procedure_status=?", service.ProcedureStatus).First(&procedureStatus).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	procedureStatus = model.ProcedureStatus{
		ProcedureStatus: service.ProcedureStatus,
		Description:     service.Description,
	}

	if err := model.DB.Create(&procedureStatus).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildProcedureStatus(procedureStatus),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListProcedureStatusService) List() serializer.Response {
	var procedureStatuss []model.ProcedureStatus
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.ProcedureStatus{})
	if service.ProcedureStatus != "" {
		tx = tx.Where(fmt.Sprintf("procedure_status like '%%%s%%'", service.ProcedureStatus))
	}
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&procedureStatuss)

	return serializer.BuildListResponse(serializer.BuildProcedureStatuss(procedureStatuss), uint(total))
}

func (service *ShowProcedureStatusService) Show(id string) serializer.Response {
	var procedureStatus model.ProcedureStatus
	code := e.SUCCESS
	err := model.DB.First(&procedureStatus, id).Error
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
		Data:   serializer.BuildProcedureStatus(procedureStatus),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteProcedureStatusService) Delete(id string) serializer.Response {
	var procedureStatus model.ProcedureStatus
	code := e.SUCCESS
	err := model.DB.First(&procedureStatus, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&procedureStatus).Error
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

func (service *UpdateProcedureStatusService) Update(id string) serializer.Response {
	var procedureStatus model.ProcedureStatus
	model.DB.Model(model.ProcedureStatus{}).Where("id=?",id).First(&procedureStatus)

	if service.ProcedureStatus != "" {
		procedureStatus.ProcedureStatus = service.ProcedureStatus
	}
	if service.Description != "" {
		procedureStatus.Description = service.Description
	}

	code := e.SUCCESS
	err := model.DB.Save(&procedureStatus).Error
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
