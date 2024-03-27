package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateMessage(c *gin.Context) {
	createService := service.CreateMessageService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListMessages(c *gin.Context) {
	listService := service.ListMessageService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowMessage(c *gin.Context) {
	showMessageService := service.ShowMessageService{}
	res := showMessageService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteMessage(c *gin.Context) {
	deleteMessageService := service.DeleteMessageService{}
	res := deleteMessageService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateMessage(c *gin.Context) {
	updateMessageService := service.UpdateMessageService{}
	if err := c.ShouldBind(&updateMessageService); err == nil {
		res := updateMessageService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
