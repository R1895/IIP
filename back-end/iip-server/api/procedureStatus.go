package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateProcedureStatus(c *gin.Context) {
	createService := service.CreateProcedureStatusService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListProcedureStatuss(c *gin.Context) {
	listService := service.ListProcedureStatusService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowProcedureStatus(c *gin.Context) {
	showProcedureStatusService := service.ShowProcedureStatusService{}
	res := showProcedureStatusService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteProcedureStatus(c *gin.Context) {
	deleteProcedureStatusService := service.DeleteProcedureStatusService{}
	res := deleteProcedureStatusService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateProcedureStatus(c *gin.Context) {
	updateProcedureStatusService := service.UpdateProcedureStatusService{}
	if err := c.ShouldBind(&updateProcedureStatusService); err == nil {
		res := updateProcedureStatusService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
