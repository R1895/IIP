package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateSafe(c *gin.Context) {
	createService := service.CreateSafeService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListSafes(c *gin.Context) {
	listService := service.ListSafeService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowSafe(c *gin.Context) {
	showSafeService := service.ShowSafeService{}
	res := showSafeService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteSafe(c *gin.Context) {
	deleteSafeService := service.DeleteSafeService{}
	res := deleteSafeService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateSafe(c *gin.Context) {
	updateSafeService := service.UpdateSafeService{}
	if err := c.ShouldBind(&updateSafeService); err == nil {
		res := updateSafeService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
