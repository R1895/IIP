package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"fmt"
)

type CreateProcedureTypeService struct {
	ProcedureName string `form:"procedure_name" json:"procedure_name" binding:"required" `
	ProcedureType string `form:"procedure_type" json:"procedure_type"  binding:"required"`
	Decription    string `form:"description" json:"description" `
}

type UpdateProcedureTypeService struct {
	ProcedureName string `form:"procedure_name" json:"procedure_name" binding:"required" `
	ProcedureType string `form:"procedure_type" json:"procedure_type"  binding:"required"`
	Decription    string `form:"description" json:"description" `
}

type ListProcedureTypeService struct {
	ProcedureName string `form:"procedure_name" json:"procedure_name" `
	ProcedureType string `form:"procedure_type" json:"procedure_type" `
	Limit         int    `form:"limit" json:"limit"`
	Start         int    `form:"start" json:"start" binding:"required"`
}

type DeleteProcedureTypeService struct {
}

type ShowProcedureTypeService struct {
}

func (service *CreateProcedureTypeService) Create() *serializer.Response {
	code := e.SUCCESS
	var proceduretype model.ProcedureType
	var count int64
	model.DB.Model(&model.ProcedureType{}).Where("procedure_type=?", service.ProcedureType).First(&proceduretype).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	proceduretype = model.ProcedureType{
		ProcedureType: service.ProcedureType,
		ProcedureName: service.ProcedureName,
		Decription:    service.Decription,
	}

	if err := model.DB.Create(&proceduretype).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildProcedureType(proceduretype),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListProcedureTypeService) List() serializer.Response {
	var proceduretypes []model.ProcedureType
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.ProcedureType{})
	if service.ProcedureName != "" {
		tx = tx.Where(fmt.Sprintf("procedure_name like '%%%s%%'", service.ProcedureName))
	}
	if service.ProcedureType != "" {
		tx = tx.Where(fmt.Sprintf("procedure_type like '%%%s%%'", service.ProcedureType))
	}
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&proceduretypes)

	return serializer.BuildListResponse(serializer.BuildProcedureTypes(proceduretypes), uint(total))
}

func (service *ShowProcedureTypeService) Show(id string) serializer.Response {
	var proceduretype model.ProcedureType
	code := e.SUCCESS
	err := model.DB.First(&proceduretype, id).Error
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
		Data:   serializer.BuildProcedureType(proceduretype),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteProcedureTypeService) Delete(id string) serializer.Response {
	var proceduretype model.ProcedureType
	code := e.SUCCESS
	err := model.DB.First(&proceduretype, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&proceduretype).Error
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

func (service *UpdateProcedureTypeService) Update(id string) serializer.Response {
	var proceduretype model.ProcedureType
	model.DB.Model(model.ProcedureType{}).Where("id=?",id).First(&proceduretype)

	if service.ProcedureName != "" {
		proceduretype.ProcedureName = service.ProcedureName
	}
	if service.ProcedureType != "" {
		proceduretype.ProcedureType = service.ProcedureType
	}
	if service.Decription != "" {
		proceduretype.Decription = service.Decription
	}

	code := e.SUCCESS
	err := model.DB.Save(&proceduretype).Error
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
