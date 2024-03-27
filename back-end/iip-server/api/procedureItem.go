package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateProcedureItem(c *gin.Context) {
	createService := service.CreateProcedureItemService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListProcedureItems(c *gin.Context) {
	listService := service.ListProcedureItemService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowProcedureItem(c *gin.Context) {
	showProcedureItemService := service.ShowProcedureItemService{}
	res := showProcedureItemService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteProcedureItem(c *gin.Context) {
	deleteProcedureItemService := service.DeleteProcedureItemService{}
	res := deleteProcedureItemService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateProcedureItem(c *gin.Context) {
	updateProcedureItemService := service.UpdateProcedureItemService{}
	if err := c.ShouldBind(&updateProcedureItemService); err == nil {
		res := updateProcedureItemService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

