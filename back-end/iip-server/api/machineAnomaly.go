package api

import (
	"iip/pkg/util"
	"iip/service"

	"github.com/gin-gonic/gin"
)

func CreateMachineAnomaly2(c *gin.Context) {
	createService := service.CreateMachineAnomalyService2{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create(c.Param("mid"), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func CreateMachineAnomaly(c *gin.Context) {
	util.LogrusObj.Error("===================")
	util.LogrusObj.Info(c.Request.Header)
	util.LogrusObj.Info("===================")

	createService := service.CreateMachineAnomalyService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListMachineAnomalys(c *gin.Context) {
	listService := service.ListMachineAnomalyService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowMachineAnomaly(c *gin.Context) {
	showMachineAnomalyService := service.ShowMachineAnomalyService{}
	res := showMachineAnomalyService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteMachineAnomaly(c *gin.Context) {
	deleteMachineAnomalyService := service.DeleteMachineAnomalyService{}
	res := deleteMachineAnomalyService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateMachineAnomaly(c *gin.Context) {
	updateMachineAnomalyService := service.UpdateMachineAnomalyService{}
	if err := c.ShouldBind(&updateMachineAnomalyService); err == nil {
		res := updateMachineAnomalyService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
