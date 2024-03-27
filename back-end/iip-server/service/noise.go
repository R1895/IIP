package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
)

type CreateNoiseService struct {
	WorkshopID  uint      `form:"workshop_id" json:"workshop_id" binding:"required"`
	Value       float64   `form:"value" json:"value"`
	Date        time.Time `form:"date" json:"date"`
	Description string    `form:"description" json:"description"`
}

type UpdateNoiseService struct {
	Value       float64   `form:"value" json:"value"`
	Date        time.Time `form:"date" json:"date"`
	Description string    `form:"description" json:"description"`
}

type ListNoiseService struct {
	WorkshopID uint      `form:"workshop_id" json:"workshop_id" `
	Date       time.Time `form:"date" json:"date"`  //TODO
	Limit      int       `form:"limit" json:"limit"`
	Start      int       `form:"start" json:"start" binding:"required"`
}

type DeleteNoiseService struct {
}

type ShowNoiseService struct {
}



func (service *CreateNoiseService) Create() *serializer.Response {
	code := e.SUCCESS
	var noise model.Noise
	noise = model.Noise{
		WorkshopId:  service.WorkshopID,
		Value:       service.Value,
		Date:        time.Now(),
		Description: service.Description,
	}

	if err := model.DB.Create(&noise).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildNoise(noise),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListNoiseService) List() serializer.Response {
	var noises []model.Noise
	var total int64

	if service.Limit == 0 {
		service.Limit = 10
	}

	// 分页
	tx := model.DB.Model(model.Noise{})
	if service.WorkshopID != 0 {
		tx = tx.Where("workshop_id =?", service.WorkshopID)
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&noises)

	return serializer.BuildListResponse(serializer.BuildNoises(noises), uint(total))
}

func (service *ShowNoiseService) Show(id string) serializer.Response {
	var noise model.Noise
	code := e.SUCCESS
	err := model.DB.First(&noise, id).Error
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
		Data:   serializer.BuildNoise(noise),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteNoiseService) Delete(id string) serializer.Response {
	var noise model.Noise
	code := e.SUCCESS
	err := model.DB.First(&noise, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&noise).Error
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

func (service *UpdateNoiseService) Update(id string) serializer.Response {
	var noise model.Noise
	model.DB.Model(model.Noise{}).Where("id = ?", id).First(&noise)

	if service.Value != noise.Value {
		noise.Value = service.Value
	}
	if service.Date != noise.Date {
		noise.Date = service.Date
	}
	code := e.SUCCESS
	err := model.DB.Save(&noise).Error
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

