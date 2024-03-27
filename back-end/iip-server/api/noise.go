package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateNoise(c *gin.Context) {
	createService := service.CreateNoiseService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListNoises(c *gin.Context) {
	listService := service.ListNoiseService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowNoise(c *gin.Context) {
	showNoiseService := service.ShowNoiseService{}
	res := showNoiseService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteNoise(c *gin.Context) {
	deleteNoiseService := service.DeleteNoiseService{}
	res := deleteNoiseService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateNoise(c *gin.Context) {
	updateNoiseService := service.UpdateNoiseService{}
	if err := c.ShouldBind(&updateNoiseService); err == nil {
		res := updateNoiseService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

