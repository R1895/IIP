package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
)

type CreateSmellService struct {
	WorkshopID  uint      `form:"workshop_id" json:"workshop_id" binding:"required"`
	Value       float64   `form:"value" json:"value"`
	Date        time.Time `form:"date" json:"date" binding:"required"`
	Description string    `form:"description" json:"description"`
}

type UpdateSmellService struct {
	Value       float64 `form:"value" json:"value"`
	Description string  `form:"description" json:"description"`
}

type ListSmellService struct {
	WorkshopID uint `form:"workshop_id" json:"workshop_id" `
	Date        time.Time `form:"date" json:"date" `
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start" binding:"required"`
}

type DeleteSmellService struct {
}

type ShowSmellService struct {
}


func (service *CreateSmellService) Create() *serializer.Response {
	code := e.SUCCESS
	var smell model.Smell

	smell = model.Smell{
		WorkshopID:  service.WorkshopID,
		Value:       service.Value,
		Date:        time.Now(), //
		Description: service.Description,
	}

	if err := model.DB.Create(&smell).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildSmell(smell),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListSmellService) List() serializer.Response {
	var smells []model.Smell
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.Smell{})

	if service.WorkshopID != 0 {
		tx = tx.Where("workshop_id =?", service.WorkshopID)
	}
	//TODO date
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&smells)

	return serializer.BuildListResponse(serializer.BuildSmells(smells), uint(total))
}

func (service *ShowSmellService) Show(id string) serializer.Response {
	var smell model.Smell
	code := e.SUCCESS
	err := model.DB.First(&smell, id).Error
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
		Data:   serializer.BuildSmell(smell),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteSmellService) Delete(id string) serializer.Response {
	var smell model.Smell
	code := e.SUCCESS
	err := model.DB.First(&smell, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&smell).Error
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

func (service *UpdateSmellService) Update(id string) serializer.Response {
	var smell model.Smell
	model.DB.Model(model.Smell{}).Where("id = ?", id).First(&smell)

	if service.Value != smell.Value {
		smell.Value = service.Value
	}
	if service.Description != "" {
		smell.Description = service.Description
	}

	code := e.SUCCESS
	err := model.DB.Save(&smell).Error
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

