package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
)

type CreateDictionaryItemService struct {
	Key         string `form:"key" json:"key" binding:"required"`
	Value       string `form:"value" json:"value" binding:"required"`
	Description string `form:"description" json:"description"`
	Category    string `form:"category" json:"category"`
}

type UpdateDictionaryItemService struct {
	Key         string `form:"key" json:"key" binding:"required"`
	Value       string `form:"value" json:"value"`
	Description string `form:"description" json:"description"`
	Category    string `form:"category" json:"category"`
}

type ListDictionaryItemService struct {
	Category string `form:"category" json:"category"`
	Limit    int    `form:"limit" json:"limit"`
	Start    int    `form:"start" json:"start" binding:"required"`
}

type DeleteDictionaryItemService struct {
}

type ShowDictionaryItemService struct {
}

func (service *CreateDictionaryItemService) Create() *serializer.Response {
	code := e.SUCCESS
	var dictionaryItem model.DictionaryItem

	var count int64
	model.DB.Model(&model.DictionaryItem{}).Where("key=?", service.Key).First(&dictionaryItem).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	dictionaryItem = model.DictionaryItem{
		Key:         service.Key,
		Value:       service.Value,
		Description: service.Description,
		Category:    service.Category,
	}

	if err := model.DB.Create(&dictionaryItem).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildDictionaryItem(dictionaryItem),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListDictionaryItemService) List() serializer.Response {
	var dictionaryItems []model.DictionaryItem
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.DictionaryItem{})
	if service.Category != "" {
		tx.Where("category = ?", service.Category)
	}
	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&dictionaryItems)

	return serializer.BuildListResponse(serializer.BuildDictionaryItems(dictionaryItems), uint(total))
}

func (service *ShowDictionaryItemService) Show(id string) serializer.Response {
	var dictionaryItem model.DictionaryItem
	code := e.SUCCESS
	err := model.DB.First(&dictionaryItem, id).Error
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
		Data:   serializer.BuildDictionaryItem(dictionaryItem),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteDictionaryItemService) Delete(id string) serializer.Response {
	var dictionaryItem model.DictionaryItem
	code := e.SUCCESS
	err := model.DB.First(&dictionaryItem, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&dictionaryItem).Error
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

func (service *UpdateDictionaryItemService) Update(id string) serializer.Response {
	var dictionaryItem model.DictionaryItem
	model.DB.Model(model.DictionaryItem{}).Where("id = ?", id).First(&dictionaryItem)

	if service.Key == dictionaryItem.Key &&
		service.Value == dictionaryItem.Value &&
		service.Description == dictionaryItem.Description &&
		service.Category == dictionaryItem.Category {
		code := e.ErrorNotUpdate
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if service.Key != dictionaryItem.Key {
		dictionaryItem.Key = service.Key
	}
	if service.Value != dictionaryItem.Value {
		dictionaryItem.Value = service.Value
	}
	if service.Description != dictionaryItem.Description {
		dictionaryItem.Description = service.Description
	}
	if service.Category != dictionaryItem.Category {
		dictionaryItem.Category = service.Category
	}
	code := e.SUCCESS
	err := model.DB.Save(&dictionaryItem).Error
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
