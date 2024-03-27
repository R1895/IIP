package service

import (
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"strconv"
	"strings"
	"time"
)

type CreateWorkerAttendanceService struct {
	WorkerID      uint      `form:"worker_id" json:"worker_id"`
	Date          time.Time `form:"date" json:"date"`
	ArrivalTime   time.Time `form:"arrival_time" json:"arrival_time"`
	DepartureTime time.Time `form:"departure_time" json:"departure_time"`
	Type          string    `form:"type" json:"type"`
}

type UpdateWorkerAttendanceService struct {
	Date          time.Time `form:"date" json:"date"`
	ArrivalTime   time.Time `form:"arrival_time" json:"arrival_time"`
	DepartureTime time.Time `form:"departure_time" json:"departure_time"`
	Type          string    `form:"type" json:"type"`
}

type ListWorkerAttendanceService struct {
	WorkerID uint   `form:"worker_id" json:"worker_id"`
	Date     string `form:"date" json:"date" `
	Limit    int    `form:"limit" json:"limit"`
	Start    int    `form:"start" json:"start" binding:"required"`
	Type     string `form:"type" json:"type"`
}

type DeleteWorkerAttendanceService struct {
}

type ShowWorkerAttendanceService struct {
}

func (service *CreateWorkerAttendanceService) Create() *serializer.Response {
	code := e.SUCCESS
	var workerAttendance model.WorkerAttendance
	// TODO 检查时间格式是否正确
	workerAttendance = model.WorkerAttendance{
		Date:          service.Date,
		WorkerID:      service.WorkerID,
		ArrivalTime:   service.ArrivalTime,
		DepartureTime: service.DepartureTime,
		Type:          service.Type,
	}

	if err := model.DB.Create(&workerAttendance).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildWorkerAttendance(workerAttendance),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListWorkerAttendanceService) List() serializer.Response {
	var workerAttendances []model.WorkerAttendance
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.WorkerAttendance{})
	if service.WorkerID != 0 {
		tx = tx.Where("worker_id =?", service.WorkerID)
	}
	date := strings.Split(service.Date, "-")
	if len(date) == 1 {
		tx.Where("YEAR(date) = ?", date[0])
	}
	if len(date) == 2 {
		tx.Where("YEAR(date) = ? AND MONTH(date) = ?", date[0], date[1])
	}
	if len(date) == 3 {
		y, err := strconv.Atoi(date[0])
		if err != nil {
			panic("failed to parse year")
		}
		m, err := strconv.Atoi(date[1])
		if err != nil {
			panic("failed to parse month")
		}

		d, err := strconv.Atoi(date[2])
		if err != nil {
			panic("failed to parse day")
		}
		targetDate := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
		tx.Where("DATE(date) = ?", targetDate.Format("2006-01-02"))
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&workerAttendances)

	return serializer.BuildListResponse(serializer.BuildWorkerAttendances(workerAttendances), uint(total))
}

func (service *ShowWorkerAttendanceService) Show(id string) serializer.Response {
	var workerAttendance model.WorkerAttendance
	code := e.SUCCESS
	err := model.DB.First(&workerAttendance, id).Error
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
		Data:   serializer.BuildWorkerAttendance(workerAttendance),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteWorkerAttendanceService) Delete(id string) serializer.Response {
	var workerAttendance model.WorkerAttendance
	code := e.SUCCESS
	err := model.DB.First(&workerAttendance, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&workerAttendance).Error
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

func (service *UpdateWorkerAttendanceService) Update(id string) serializer.Response {
	var workerAttendance model.WorkerAttendance
	model.DB.Model(model.WorkerAttendance{}).Where("id = ?", id).First(&workerAttendance)

	// TODO 添加约束
	// if service.Date.IsZero() {
	// 	workerAttendance.Date = service.Date
	// }

	workerAttendance.DepartureTime = service.DepartureTime
	workerAttendance.ArrivalTime = service.ArrivalTime

	code := e.SUCCESS
	err := model.DB.Save(&workerAttendance).Error
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
