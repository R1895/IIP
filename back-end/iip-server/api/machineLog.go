package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateMachineLog(c *gin.Context) {
	createService := service.CreateMachineLogService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListMachineLogs(c *gin.Context) {
	listService := service.ListMachineLogService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowMachineLog(c *gin.Context) {
	showMachineLogService := service.ShowMachineLogService{}
	res := showMachineLogService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteMachineLog(c *gin.Context) {
	deleteMachineLogService := service.DeleteMachineLogService{}
	res := deleteMachineLogService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateMachineLog(c *gin.Context) {
	updateMachineLogService := service.UpdateMachineLogService{}
	if err := c.ShouldBind(&updateMachineLogService); err == nil {
		res := updateMachineLogService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

