package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"fmt"
)

type CreateMessageService struct {
	MachineID uint   `form:"machine_id" json:"machine_id"`
	UserID    uint   `form:"user_id" json:"user_id"`
	Title     string `form:"title" json:"title"`
	Content   string `form:"content" json:"content"`
	Type      string `form:"type" json:"type"`
}

type UpdateMessageService struct {
	IsRead  bool   `form:"is_read" json:"is_read"`
	Content string `form:"content" json:"content"`
	Type    string `form:"type" json:"type"`
}

type ListMessageService struct {
	MachineID uint   `form:"machine_id" json:"machine_id"`
	UserID    uint   `form:"user_id" json:"user_id"`
	Type      string `form:"type" json:"type"`
	Limit     int    `form:"limit" json:"limit"`
	Start     int    `form:"start" json:"start" binding:"required"`
}

type DeleteMessageService struct {
}

type ShowMessageService struct {
}

func (service *CreateMessageService) Create() *serializer.Response {
	code := e.SUCCESS
	var message model.Message

	message = model.Message{
		MachineID: service.MachineID,
		UserID:    service.UserID,
		Title:     service.Title,
		Content:   service.Content,
		Type:      service.Type,
		IsRead:    false,
	}
	if err := model.DB.Create(&message).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildMessage(message),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListMessageService) List() serializer.Response {
	var messages []model.Message
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}

	tx := model.DB.Model(model.Message{})
	if service.MachineID != 0 {
		tx = tx.Where("machine_id = ?", service.MachineID)
	}
	if service.UserID != 0 {
		tx = tx.Where("user_id = ?", service.UserID)
	}
	if service.Type != "" {
		tx = tx.Where(fmt.Sprintf("type like '%%%s%%'", service.Type))
	}
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&messages)

	return serializer.BuildListResponse(
		serializer.BuildMessages(messages), uint(total))
}

func (service *ShowMessageService) Show(id string) serializer.Response {
	var message model.Message
	code := e.SUCCESS
	err := model.DB.Model(model.Message{}).First(&message, id).Error
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
		Data:   serializer.BuildMessage(message),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteMessageService) Delete(id string) serializer.Response {
	var message model.Message
	code := e.SUCCESS
	err := model.DB.First(&message, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&message).Error
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

func (service *UpdateMessageService) Update(id string) serializer.Response {
	var message model.Message
	model.DB.Model(model.Message{}).Where("id = ?", id).First(&message)
	message.IsRead = service.IsRead
	if service.Content != "" {
		message.Content = service.Content
	}
	if service.Type != "" {
		message.Type = service.Type
	}
	code := e.SUCCESS
	err := model.DB.Save(&message).Error
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
