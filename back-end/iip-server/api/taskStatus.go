package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateTaskStatus(c *gin.Context) {
	createService := service.CreateTaskStatusService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListTaskStatuss(c *gin.Context) {
	listService := service.ListTaskStatusService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowTaskStatus(c *gin.Context) {
	showTaskStatusService := service.ShowTaskStatusService{}
	res := showTaskStatusService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteTaskStatus(c *gin.Context) {
	deleteTaskStatusService := service.DeleteTaskStatusService{}
	res := deleteTaskStatusService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateTaskStatus(c *gin.Context) {
	updateTaskStatusService := service.UpdateTaskStatusService{}
	if err := c.ShouldBind(&updateTaskStatusService); err == nil {
		res := updateTaskStatusService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
