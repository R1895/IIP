package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateMachineStatus(c *gin.Context) {
	createService := service.CreateMachineStatusService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListMachineStatuss(c *gin.Context) {
	listService := service.ListMachineStatusService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowMachineStatus(c *gin.Context) {
	showMachineStatusService := service.ShowMachineStatusService{}
	res := showMachineStatusService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteMachineStatus(c *gin.Context) {
	deleteMachineStatusService := service.DeleteMachineStatusService{}
	res := deleteMachineStatusService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateMachineStatus(c *gin.Context) {
	updateMachineStatusService := service.UpdateMachineStatusService{}
	if err := c.ShouldBind(&updateMachineStatusService); err == nil {
		res := updateMachineStatusService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}