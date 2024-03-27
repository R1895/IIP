package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
)

type CreateSafeService struct {
	WorkshopID  uint      `form:"workshop_id" json:"workshop_id" binding:"required"`
	Value       float64   `form:"value" json:"value"`
	Description string    `form:"description" json:"description"`
	Type        string    `form:"type" json:"type" binding:"required" `
	Date        time.Time `form:"date" json:"date"`
}

type UpdateSafeService struct {
	Value       float64 `form:"value" json:"value"`
	Type        string    `form:"type" json:"type"`
	Description string  `form:"description" json:"description"`
}

type ListSafeService struct {
	WorkshopID uint      `form:"workshop_id" json:"workshop_id" `
	Type        string    `form:"type" json:"type"`
	Date       time.Time `form:"date" json:"date"`
	Limit      int       `form:"limit" json:"limit"`
	Start      int       `form:"start" json:"start" binding:"required"`
}

type DeleteSafeService struct {
}

type ShowSafeService struct {
}



func (service *CreateSafeService) Create() *serializer.Response {
	code := e.SUCCESS
	var safe model.Safe

	safe = model.Safe{
		WorkshopID:  service.WorkshopID,
		Value:       service.Value,
		Date:        time.Now(),
		Description: service.Description,
	}

	if err := model.DB.Create(&safe).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildSafe(safe),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListSafeService) List() serializer.Response {
	var safes []model.Safe
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.Safe{})

	if service.WorkshopID != 0 {
		tx = tx.Where("workshop_id =?", service.WorkshopID)
	}
	// if service.Date != "" {
	// 	tx = tx.Where(fmt.Sprintf("date(date) =?", service.WorkshopID))
	// }
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&safes)

	return serializer.BuildListResponse(serializer.BuildSafes(safes), uint(total))
}

func (service *ShowSafeService) Show(id string) serializer.Response {
	var safe model.Safe
	code := e.SUCCESS
	err := model.DB.First(&safe, id).Error
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
		Data:   serializer.BuildSafe(safe),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteSafeService) Delete(id string) serializer.Response {
	var safe model.Safe
	code := e.SUCCESS
	err := model.DB.First(&safe, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&safe).Error
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

func (service *UpdateSafeService) Update(id string) serializer.Response {
	var safe model.Safe
	model.DB.Model(model.Safe{}).Where("id = ?", id).First(&safe)

	if service.Value != safe.Value {
		safe.Value = service.Value
	}
	if service.Description != safe.Description {
		safe.Value = service.Value
	}
	code := e.SUCCESS
	err := model.DB.Save(&safe).Error
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

