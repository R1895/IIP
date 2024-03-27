package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateLog(c *gin.Context) {
	createService := service.CreateLogService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListLogs(c *gin.Context) {
	listService := service.ListLogService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowLog(c *gin.Context) {
	showLogService := service.ShowLogService{}
	res := showLogService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteLog(c *gin.Context) {
	deleteLogService := service.DeleteLogService{}
	res := deleteLogService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateLog(c *gin.Context) {
	updateLogService := service.UpdateLogService{}
	if err := c.ShouldBind(&updateLogService); err == nil {
		res := updateLogService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

