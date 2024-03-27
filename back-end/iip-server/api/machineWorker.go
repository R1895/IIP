package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateMachineWorker(c *gin.Context) {
	createService := service.CreateMachineWorkerService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListMachineWorkers(c *gin.Context) {
	listService := service.ListMachineWorkerService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowMachineWorker(c *gin.Context) {
	showMachineWorkerService := service.ShowMachineWorkerService{}
	res := showMachineWorkerService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteMachineWorker(c *gin.Context) {
	deleteMachineWorkerService := service.DeleteMachineWorkerService{}
	res := deleteMachineWorkerService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateMachineWorker(c *gin.Context) {
	updateMachineWorkerService := service.UpdateMachineWorkerService{}
	if err := c.ShouldBind(&updateMachineWorkerService); err == nil {
		res := updateMachineWorkerService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}