package api

import (
	"iip/pkg/util"
	"iip/service"

	"github.com/gin-gonic/gin"
)

func CreateIventory(c *gin.Context) {
	createService := service.CreateIventoryService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListIventorys(c *gin.Context) {
	listService := service.ListIventoryService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowIventory(c *gin.Context) {
	showIventoryService := service.ShowIventoryService{}
	res := showIventoryService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteIventory(c *gin.Context) {
	deleteIventoryService := service.DeleteIventoryService{}
	res := deleteIventoryService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateIventory(c *gin.Context) {
	updateIventoryService := service.UpdateIventoryService{}
	if err := c.ShouldBind(&updateIventoryService); err == nil {
		res := updateIventoryService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

