package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
)

type CreateWorkerService struct {
	UserID       uint   `form:"user_id" json:"user_id"  binding:"required"`
	WorkerTypeID uint   `form:"worker_type_id" json:"worker_type_id"  binding:"required"`
	WorkerType   string `form:"worker_type" json:"worker_type"`
	WorkerName   string `form:"worker_name" json:"worker_name"  binding:"required,min=2,max=20"`
	SkillLevel   int    `form:"skill_level" json:"skill_level"`
}

type UpdateWorkerService struct {
	WorkerTypeID uint   `form:"woker_type_id" json:"woker_type_id" `
	WorkerType   string `form:"worker_type" json:"worker_type"`
	WorkerName   string `form:"worker_name" json:"worker_name"  binding:"required,min=2,max=20"`
	SkillLevel   int    `form:"skill_level" json:"skill_level"`
}

type ListWorkerService struct {
	WorkerType string `form:"worker_type" json:"worker_type"`
	SkillLevel int    `form:"skill_level" json:"skill_level"`
	Type       string `form:"type" json:"type"`
	Limit      int    `form:"limit" json:"limit"`
	Start      int    `form:"start" json:"start" binding:"required"`
}

type DeleteWorkerService struct {
}

type ShowWorkerService struct {
}

func (service *CreateWorkerService) Create() *serializer.Response {
	code := e.SUCCESS
	var worker model.Worker

	var count int64
	model.DB.Model(&model.Worker{}).Where("worker_name=?", service.WorkerName).First(&worker).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	var user model.User
	model.DB.Model(&model.User{}).Where("id=?", service.UserID).First(&user).Count(&count)

	if count == 0 {
		code = e.ErrorNotExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "User不存在",
		}
	}

	worker = model.Worker{
		UserID:       service.UserID,
		WorkerTypeID: service.WorkerTypeID,
		WorkerType:   service.WorkerType,
		WorkerName:   service.WorkerName,
		SkillLevel:   service.SkillLevel,
	}

	if err := model.DB.Create(&worker).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildWorker(worker),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListWorkerService) List() serializer.Response {
	var workers []model.Worker
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.Worker{}).Preload("Workshops").Preload("MachineWorkers").Preload("WorkerAttendances")

	if service.WorkerType != "" {
		tx = tx.Where("worker_type =?", service.WorkerType)
	}
	if service.SkillLevel != 0 {
		tx = tx.Where("skill_level =?", service.SkillLevel)
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&workers)

	var workers2 []model.Worker
	if service.Type == "1" {
		for _, worker := range workers {
			if worker.Workshops == nil || len(worker.Workshops) == 0 {
				workers2 = append(workers2, worker)
			}
		}
		
		return serializer.BuildListResponse(serializer.BuildWorkers(workers2), uint(len(workers2)))
	}

	return serializer.BuildListResponse(serializer.BuildWorkers(workers), uint(total))
}

func (service *ShowWorkerService) Show(id string) serializer.Response {
	var worker model.Worker
	code := e.SUCCESS
	err := model.DB.First(&worker, id).Error
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
		Data:   serializer.BuildWorker(worker),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteWorkerService) Delete(id string) serializer.Response {
	var worker model.Worker
	code := e.SUCCESS
	err := model.DB.First(&worker, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&worker).Error
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

func (service *UpdateWorkerService) Update(id string) serializer.Response {
	var worker model.Worker
	code := e.SUCCESS
	errs := model.DB.Model(model.Worker{}).Where("id = ?", id).First(&worker).Error
	if errs != nil {
		util.LogrusObj.Info(errs)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errs.Error(),
		}
	}

	if service.WorkerType != "" {
		worker.WorkerType = service.WorkerType
	}
	if service.WorkerTypeID != 0 {
		worker.WorkerTypeID = service.WorkerTypeID
	}
	if service.SkillLevel != 0 {
		worker.SkillLevel = service.SkillLevel
	}
	if service.WorkerName != "" {
		worker.WorkerName = service.WorkerName
	}

	err := model.DB.Save(&worker).Error
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
