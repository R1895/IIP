package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"fmt"
)

type CreateProcedureItemService struct {
	ProcedureID uint   `form:"procedure_id" json:"procedure_id" binding:"required"`
	MachineID   uint   `form:"machine_id" json:"machine_id" binding:"required"`
	WorkerID    uint   `form:"worker_id" json:"worker_id"`
	ItemType    string `form:"item_type" json:"item_type"`
	ItemName    string `form:"item_name" json:"item_name"`
	Quality     string `form:"quality" json:"quality"`
}

type UpdateProcedureItemService struct {
	ProcedureID uint   `form:"procedure_id" json:"procedure_id"`
	MachineID   uint   `form:"machine_id" json:"machine_id"`
	WorkerID    uint   `form:"worker_id" json:"worker_id"`
	ItemType    string `form:"item_type" json:"item_type"`
	ItemName    string `form:"item_name" json:"item_name"`
	Quality     string `form:"quality" json:"quality"`
}

type ListProcedureItemService struct {
	ProcedureID uint   `form:"procedure_id" json:"procedure_id" `
	MachineID   uint   `form:"machine_id" json:"machine_id"`
	WorkerID    uint   `form:"worker_id" json:"worker_id"`
	ItemType    string `form:"item_type" json:"item_type"`
	Limit       int    `form:"limit" json:"limit"`
	Start       int    `form:"start" json:"start" binding:"required"`
}

type DeleteProcedureItemService struct {
}

type ShowProcedureItemService struct {
}


func (service *CreateProcedureItemService) Create() *serializer.Response {
	code := e.SUCCESS
	var procedureItem model.ProcedureItem

	procedureItem = model.ProcedureItem{
		ProcedureID: service.ProcedureID,
		MachineID:   service.MachineID,
		WorkerID:    service.WorkerID,
		ItemType:    service.ItemType,
		ItemName:    service.ItemName,
		Quality:     service.Quality,
	}

	if err := model.DB.Create(&procedureItem).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildProcedureItem(procedureItem),
		Msg:    e.GetMsg(code),
	}
}

// TODO 分页查询条件
func (service *ListProcedureItemService) List() serializer.Response {
	var procedureItems []model.ProcedureItem
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.ProcedureItem{})
	if service.ProcedureID != 0 {
		tx = tx.Where("procedure_id = ?", service.ProcedureID)
	}
	if service.MachineID != 0 {
		tx = tx.Where("machine_id = ?", service.MachineID)
	}
	if service.WorkerID != 0 {
		tx = tx.Where("worker_id = ?", service.WorkerID)
	}
	if service.ItemType != "" {
		tx = tx.Where(fmt.Sprintf("procedure_status like '%%%s%%'", service.ItemType))
	}
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&procedureItems)

	return serializer.BuildListResponse(serializer.BuildProcedureItems(procedureItems), uint(total))
}

func (service *ShowProcedureItemService) Show(id string) serializer.Response {
	var procedureItem model.ProcedureItem
	code := e.SUCCESS
	err := model.DB.First(&procedureItem, id).Error
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
		Data:   serializer.BuildProcedureItem(procedureItem),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteProcedureItemService) Delete(id string) serializer.Response {
	var procedureItem model.ProcedureItem
	code := e.SUCCESS
	err := model.DB.First(&procedureItem, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&procedureItem).Error
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

func (service *UpdateProcedureItemService) Update(id string) serializer.Response {
	var procedureItem model.ProcedureItem
	model.DB.Model(model.ProcedureItem{}).Where("id=?", id).First(&procedureItem)

	if service.ProcedureID != 0 {
		procedureItem.ProcedureID = service.ProcedureID
	}
	if service.MachineID != 0 {
		procedureItem.MachineID = service.MachineID
	}
	if service.WorkerID != 0 {
		procedureItem.WorkerID = service.WorkerID
	}
	if service.ItemType != "" {
		procedureItem.ItemType = service.ItemType
	}
	if service.ItemName != "" {
		procedureItem.ItemName = service.ItemName
	}
	if service.Quality != "" {
		procedureItem.Quality = service.Quality
	}

	code := e.SUCCESS
	err := model.DB.Save(&procedureItem).Error
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
