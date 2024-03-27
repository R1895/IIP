package service

import (
	"fmt"
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
)

type CreateWorkshopService struct {
	WorkshopName string `form:"workshop_name" json:"workshop_name" binding:"required,min=2,max=20"`
	WorkerId     uint   `form:"worker_id" json:"worker_id" binding:"required"`
}

type UpdateWorkshopService struct {
	WorkshopName string `form:"workshop_name" json:"workshop_name" `
	WorkerId     uint   `form:"worker_id" json:"worker_id"`
}

type ListWorkshopService struct {
	WorkshopName string `form:"workshop_name" json:"workshop_name"`
	WorkerId     uint   `form:"worker_id" json:"worker_id"`
	Limit        int    `form:"limit" json:"limit" `
	Start        int    `form:"start" json:"start" binding:"required"`
}

type DeleteWorkshopService struct {
}

type ShowWorkshopService struct {
}

func (service *CreateWorkshopService) Create() *serializer.Response {
	code := e.SUCCESS
	var workshop model.Workshop
	var count int64

	model.DB.Model(&model.Workshop{}).Where("workshop_name=?", service.WorkshopName).First(&workshop).Count(&count)
	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 如果worker不存在直接返回
	var worker model.Worker
	model.DB.Model(&model.Worker{}).Where("id=?", service.WorkerId).First(&worker).Count(&count)
	if count == 0 {
		code = e.ErrorNotExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	workshop = model.Workshop{
		WorkshopName: service.WorkshopName,
		WorkerId:     service.WorkerId,
	}

	if err := model.DB.Create(&workshop).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildWorkshop(workshop),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListWorkshopService) List() serializer.Response {
	var workshops []model.Workshop
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}

	// 预加载和分页
	tx := model.DB.Model(model.Workshop{}).
		Preload("Worker").Preload("Workers").
		Preload("Machines").Preload("Tasks").
		Preload("Environments").Preload("Noises").
		Preload("Smells").Preload("Safes")

	if service.WorkshopName != "" {
		tx = tx.Where(fmt.Sprintf("workshop_name like '%%%s%%' ", service.WorkshopName))
	}

	if service.WorkerId != 0 {
		tx = tx.Where("worker_id =?", service.WorkerId)
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&workshops)

	return serializer.BuildListResponse(
		serializer.BuildWorkshops(workshops), uint(total))
}

// 根据id 查询车间详情，包括Worker，Workers，Machines，Tasks
func (service *ShowWorkshopService) Show(id string) serializer.Response {
	var workshop model.Workshop
	code := e.SUCCESS
	err := model.DB.Model(model.Workshop{}).
		Preload("Worker").Preload("Workers").Preload("Machines").
		Preload("Tasks").
		First(&workshop, id).Error
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
		Data:   serializer.BuildWorkshop(workshop),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteWorkshopService) Delete(id string) serializer.Response {
	var workshop model.Workshop
	code := e.SUCCESS
	err := model.DB.First(&workshop, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&workshop).Error
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

// 根据id 更新 workshop
func (service *UpdateWorkshopService) Update(id string) serializer.Response {
	var workshop model.Workshop
	model.DB.Model(model.Workshop{}).Where("id = ?", id)

	if service.WorkshopName == workshop.WorkshopName && service.WorkerId == workshop.WorkerId {
		code := e.ErrorNotUpdate
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if service.WorkshopName != workshop.WorkshopName {
		workshop.WorkshopName = service.WorkshopName
	}
	if service.WorkerId != workshop.WorkerId {
		workshop.WorkerId = service.WorkerId
	}

	code := e.SUCCESS
	err := model.DB.Save(&workshop).Error
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

// 为车间分配工人
type AddWorkerService struct {
	WorkshopID uint `form:"workshop_id" json:"workshop_id" binding:"required"`
	WorkerID   uint `form:"worker_id" json:"worker_id" binding:"required"`
}

func (service *AddWorkerService) AddWorker() serializer.Response {
	code := e.SUCCESS
	workshopid := service.WorkshopID
	workerid := service.WorkerID

	// 如果已经存在 直接返回
	var worker model.Worker
	var count int64
	model.DB.Model(&model.Worker{}).Preload("Workshops").Where("id=?", service.WorkerID).First(&worker).Count(&count)
	if count > 0 {
		for _, workshop := range worker.Workshops {
			if workshop.ID == service.WorkshopID {
				code = e.ErrorExist
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
				}
			}
		}
	}

	insertQuery := "INSERT INTO worker_workshop (workshop_id, worker_id) VALUES (?, ?)"
	err := model.DB.Exec(insertQuery, workshopid, workerid).Error

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
		Data:   "添加成功",
	}
}

// TODO 删除
// 为车间分配工人
type RemoveWorkerService struct {
	WorkshopID uint `form:"workshop_id" json:"workshop_id" binding:"required"`
	WorkerID   uint `form:"worker_id" json:"worker_id" binding:"required"`
}

func (service *RemoveWorkerService) RemoveWorker() serializer.Response {
	code := e.SUCCESS

	err := model.DB.Exec("DELETE FROM worker_workshop WHERE workshop_id = ? AND worker_id = ?", service.WorkshopID, service.WorkerID).Error
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
		Data:   "删除成功",
	}
}
