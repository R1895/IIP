package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateProcedureType(c *gin.Context) {
	createService := service.CreateProcedureTypeService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListProcedureTypes(c *gin.Context) {
	listService := service.ListProcedureTypeService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowProcedureType(c *gin.Context) {
	showProcedureTypeService := service.ShowProcedureTypeService{}
	res := showProcedureTypeService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteProcedureType(c *gin.Context) {
	deleteProcedureTypeService := service.DeleteProcedureTypeService{}
	res := deleteProcedureTypeService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateProcedureType(c *gin.Context) {
	updateProcedureTypeService := service.UpdateProcedureTypeService{}
	if err := c.ShouldBind(&updateProcedureTypeService); err == nil {
		res := updateProcedureTypeService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
