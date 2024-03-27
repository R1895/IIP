package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
	"fmt"
)

type CreateIventoryService struct {
	ItemType     string    `form:"item_type" json:"item_type" binding:"required"`
	IventoryType uint      `form:"iventory_type" json:"iventory_type" binding:"required"`
	IventoryNum  uint      `form:"iventory_num" json:"iventory_num" binding:"required"`
	Date         time.Time `form:"date" json:"date"`
	Description  string    `form:"description" json:"description"`
	Field1       string    `form:"field1" json:"field1"`
}

type UpdateIventoryService struct {
	ItemType     string    `form:"item_type" json:"item_type" `
	IventoryType uint      `form:"iventory_type" json:"iventory_type" `
	IventoryNum  uint      `form:"iventory_num" json:"iventory_num" `
	Date         time.Time `form:"date" json:"date"`
	Description  string    `form:"description" json:"description"`
	Field1       string    `form:"field1" json:"field1"`
}

type ListIventoryService struct {
	ItemType string    `form:"item_type" json:"item_type" `
	Field1   string    `form:"field1" json:"field1"`
	Date     time.Time `form:"date" json:"date"`
	Limit    int       `form:"limit" json:"limit"`
	Start    int       `form:"start" json:"start" binding:"required"`
}

type DeleteIventoryService struct {
}

type ShowIventoryService struct {
}

func (service *CreateIventoryService) Create() *serializer.Response {
	code := e.SUCCESS
	var iventory model.Iventory
	iventory = model.Iventory{
		ItemType:     service.ItemType,
		IventoryType: service.IventoryType,
		IventoryNum:  service.IventoryNum,
		Date:         time.Now(),
		Description:  service.Description,
		Field1:       service.Field1,
	}

	if err := model.DB.Create(&iventory).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildIventory(iventory),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListIventoryService) List() serializer.Response {
	var iventorys []model.Iventory
	var total int64

	if service.Limit == 0 {
		service.Limit = 10
	}
	
	// 分页
	tx := model.DB.Model(model.Iventory{})
	if service.Field1 !="" {
		tx = tx.Where(fmt.Sprintf("type like '%%%s%%'", service.Field1))
	}
	if service.ItemType != "" {
		tx = tx.Where(fmt.Sprintf("type like '%%%s%%'", service.ItemType))
	}
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&iventorys)

	return serializer.BuildListResponse(serializer.BuildIventorys(iventorys), uint(total))
}

func (service *ShowIventoryService) Show(id string) serializer.Response {
	var iventory model.Iventory
	code := e.SUCCESS
	err := model.DB.First(&iventory, id).Error
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
		Data:   serializer.BuildIventory(iventory),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteIventoryService) Delete(id string) serializer.Response {
	var iventory model.Iventory
	code := e.SUCCESS
	err := model.DB.First(&iventory, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&iventory).Error
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

func (service *UpdateIventoryService) Update(id string) serializer.Response {
	var iventory model.Iventory
	model.DB.Model(model.Iventory{}).Where("id = ?", id).First(&iventory)

	if service.ItemType != "" {
		iventory.ItemType = service.ItemType
	}

	if service.IventoryType != 0 {
		iventory.IventoryType = service.IventoryType
	}

	if service.Description != "" {
		iventory.Description = service.Description
	}

	service.Date = time.Now()

	if service.IventoryNum != 0 {
		iventory.IventoryNum = service.IventoryNum
	}

	if service.Field1 != "" {
		iventory.Field1 = service.Field1
	}

	code := e.SUCCESS
	err := model.DB.Save(&iventory).Error
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
