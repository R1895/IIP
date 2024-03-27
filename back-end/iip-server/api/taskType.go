package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateTaskType(c *gin.Context) {
	createService := service.CreateTaskTypeService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListTaskTypes(c *gin.Context) {
	listService := service.ListTaskTypeService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowTaskType(c *gin.Context) {
	showTaskTypeService := service.ShowTaskTypeService{}
	res := showTaskTypeService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteTaskType(c *gin.Context) {
	deleteTaskTypeService := service.DeleteTaskTypeService{}
	res := deleteTaskTypeService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateTaskType(c *gin.Context) {
	updateTaskTypeService := service.UpdateTaskTypeService{}
	if err := c.ShouldBind(&updateTaskTypeService); err == nil {
		res := updateTaskTypeService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
