package api

import (
	"iip/pkg/util"
	"iip/service"

	"github.com/gin-gonic/gin"
)

func CreateEnvironment(c *gin.Context) {
	createService := service.CreateEnvironmentService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListEnvironments(c *gin.Context) {
	listService := service.ListEnvironmentService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowEnvironment(c *gin.Context) {
	showEnvironmentService := service.ShowEnvironmentService{}
	res := showEnvironmentService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteEnvironment(c *gin.Context) {
	deleteEnvironmentService := service.DeleteEnvironmentService{}
	res := deleteEnvironmentService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateEnvironment(c *gin.Context) {
	updateEnvironmentService := service.UpdateEnvironmentService{}
	if err := c.ShouldBind(&updateEnvironmentService); err == nil {
		res := updateEnvironmentService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
