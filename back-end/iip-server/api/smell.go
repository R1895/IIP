package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateSmell(c *gin.Context) {
	createService := service.CreateSmellService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListSmells(c *gin.Context) {
	listService := service.ListSmellService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowSmell(c *gin.Context) {
	showSmellService := service.ShowSmellService{}
	res := showSmellService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteSmell(c *gin.Context) {
	deleteSmellService := service.DeleteSmellService{}
	res := deleteSmellService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateSmell(c *gin.Context) {
	updateSmellService := service.UpdateSmellService{}
	if err := c.ShouldBind(&updateSmellService); err == nil {
		res := updateSmellService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

