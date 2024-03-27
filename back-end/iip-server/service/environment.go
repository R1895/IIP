package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
)

type CreateEnvironmentService struct {
	WorkshopID  uint      `form:"workshop_id" json:"workshop_id" binding:"required"`
	Level       uint      `form:"level" json:"level"`
	Date        time.Time `form:"date" json:"date"`
	Description string    `form:"description" json:"description"`
}

type UpdateEnvironmentService struct {
	Level       uint      `form:"level" json:"level"`
	Date        time.Time `form:"date" json:"date"`
	Description string    `form:"description" json:"description"`
}

type ListEnvironmentService struct {
	WorkshopID uint      `form:"workshop_id" json:"workshop_id" `
	Level      uint      `form:"level" json:"level"`
	Date       time.Time `form:"date" json:"date"`
	Limit      int       `form:"limit" json:"limit"`
	Start      int       `form:"start" json:"start" binding:"required"`
}

type DeleteEnvironmentService struct {
}

type ShowEnvironmentService struct {
}



func (service *CreateEnvironmentService) Create() *serializer.Response {
	code := e.SUCCESS
	var environment model.Environment
	environment = model.Environment{
		WorkshopID:  service.WorkshopID,
		Level:       service.Level,
		Date:        time.Now(),
		Description: service.Description,
	}

	if err := model.DB.Create(&environment).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildEnvironment(environment),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListEnvironmentService) List() serializer.Response {
	var environments []model.Environment
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	tx := model.DB.Model(model.Environment{})

	if service.Level != 0 {
		tx = tx.Where("level =?", service.Level)
	}
	if service.WorkshopID != 0 {
		tx = tx.Where("workshop_id =?", service.WorkshopID)
	}
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&environments)

	return serializer.BuildListResponse(serializer.BuildEnvironments(environments), uint(total))
}

func (service *ShowEnvironmentService) Show(id string) serializer.Response {
	var environment model.Environment
	code := e.SUCCESS
	err := model.DB.First(&environment, id).Error
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
		Data:   serializer.BuildEnvironment(environment),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteEnvironmentService) Delete(id string) serializer.Response {
	var environment model.Environment
	code := e.SUCCESS
	err := model.DB.First(&environment, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&environment).Error
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

func (service *UpdateEnvironmentService) Update(id string) serializer.Response {
	var environment model.Environment
	model.DB.Model(model.Environment{}).Where("id = ?", id).First(&environment)
	if service.Level != 0 {
		environment.Level = service.Level
	}

	if service.Date != environment.Date {
		environment.Date = service.Date
	}
	if service.Description != environment.Description {
		environment.Description = service.Description
	}
	code := e.SUCCESS
	err := model.DB.Save(&environment).Error
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

