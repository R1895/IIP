package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateWorkerType(c *gin.Context) {
	createService := service.CreateWorkerTypeService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListWorkerTypes(c *gin.Context) {
	listService := service.ListWorkerTypeService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowWorkerType(c *gin.Context) {
	showWorkerTypeService := service.ShowWorkerTypeService{}
	res := showWorkerTypeService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteWorkerType(c *gin.Context) {
	deleteWorkerTypeService := service.DeleteWorkerTypeService{}
	res := deleteWorkerTypeService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateWorkerType(c *gin.Context) {
	updateWorkerTypeService := service.UpdateWorkerTypeService{}
	if err := c.ShouldBind(&updateWorkerTypeService); err == nil {
		res := updateWorkerTypeService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
