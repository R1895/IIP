package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"
)

type CreateLogService struct {
	UserID    uint      `form:"user_id" json:"user_id"`
	Action    string    `form:"action" json:"action"`
	Method    string    `form:"method" json:"method"`
	Timestamp time.Time `form:"time_stamp" json:"time_stamp"`
	Details   string    `form:"details" json:"details"`
}

type UpdateLogService struct {
	UserID    uint      `form:"user_id" json:"user_id"`
	Action    string    `form:"action" json:"action"`
	Method    string    `form:"method" json:"method"`
	Timestamp time.Time `form:"time_stamp" json:"time_stamp"`
	Details   string    `form:"details" json:"details"`
}

type ListLogService struct {
	UserID uint `form:"user_id" json:"user_id"`
	Limit  int  `form:"limit" json:"limit"`
	Start  int  `form:"start" json:"start" binding:"required"`
}

type DeleteLogService struct {
}

type ShowLogService struct {
}


func (service *CreateLogService) Create() *serializer.Response {
	code := e.SUCCESS
	var log model.Log
	var count int64
	model.DB.Model(&model.Log{}).Where("user_id=?", service.UserID).First(&log).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	log = model.Log{

		UserID:    service.UserID,
		Action:    service.Action,
		Method:    service.Method,
		Timestamp: service.Timestamp,
		Details:   service.Details,
	}

	if err := model.DB.Create(&log).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildLog(log),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListLogService) List() serializer.Response {
	var logs []model.Log
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.Log{})
	if service.UserID != 0 {
		tx = tx.Where("user_id =?", service.UserID)
	}
	
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&logs)

	return serializer.BuildListResponse(serializer.BuildLogs(logs), uint(total))
}

func (service *ShowLogService) Show(id string) serializer.Response {
	var log model.Log
	code := e.SUCCESS
	err := model.DB.First(&log, id).Error
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
		Data:   serializer.BuildLog(log),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteLogService) Delete(id string) serializer.Response {
	var log model.Log
	code := e.SUCCESS
	err := model.DB.First(&log, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&log).Error
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

func (service *UpdateLogService) Update(id string) serializer.Response {
	var log model.Log
	model.DB.Model(model.Log{}).Where("id = ?", id).First(&log)
	if service.UserID != 0 {
		log.UserID = service.UserID
	}
	log.Action = service.Action
	log.Method = service.Method
	log.Timestamp = time.Now() ///
	log.Timestamp = service.Timestamp
	log.Details = service.Details

	code := e.SUCCESS
	err := model.DB.Save(&log).Error
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
