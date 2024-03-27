package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateDictionaryItem(c *gin.Context) {
	createService := service.CreateDictionaryItemService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListDictionaryItems(c *gin.Context) {
	listService := service.ListDictionaryItemService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowDictionaryItem(c *gin.Context) {
	showDictionaryItemService := service.ShowDictionaryItemService{}
	res := showDictionaryItemService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteDictionaryItem(c *gin.Context) {
	deleteDictionaryItemService := service.DeleteDictionaryItemService{}
	res := deleteDictionaryItemService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateDictionaryItem(c *gin.Context) {
	updateDictionaryItemService := service.UpdateDictionaryItemService{}
	if err := c.ShouldBind(&updateDictionaryItemService); err == nil {
		res := updateDictionaryItemService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
