package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateWorkerAttendance(c *gin.Context) {
	createService := service.CreateWorkerAttendanceService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListWorkerAttendances(c *gin.Context) {
	listService := service.ListWorkerAttendanceService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowWorkerAttendance(c *gin.Context) {
	showWorkerAttendanceService := service.ShowWorkerAttendanceService{}
	res := showWorkerAttendanceService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteWorkerAttendance(c *gin.Context) {
	deleteWorkerAttendanceService := service.DeleteWorkerAttendanceService{}
	res := deleteWorkerAttendanceService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateWorkerAttendance(c *gin.Context) {
	updateWorkerAttendanceService := service.UpdateWorkerAttendanceService{}
	if err := c.ShouldBind(&updateWorkerAttendanceService); err == nil {
		res := updateWorkerAttendanceService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

