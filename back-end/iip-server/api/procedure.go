package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateProcedure(c *gin.Context) {
	createService := service.CreateProcedureService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListProcedures(c *gin.Context) {
	listService := service.ListProcedureService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowProcedure(c *gin.Context) {
	showProcedureService := service.ShowProcedureService{}
	res := showProcedureService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteProcedure(c *gin.Context) {
	deleteProcedureService := service.DeleteProcedureService{}
	res := deleteProcedureService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateProcedure(c *gin.Context) {
	updateProcedureService := service.UpdateProcedureService{}
	if err := c.ShouldBind(&updateProcedureService); err == nil {
		res := updateProcedureService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
