package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateWorker(c *gin.Context) {
	createService := service.CreateWorkerService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListWorkers(c *gin.Context) {
	listService := service.ListWorkerService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowWorker(c *gin.Context) {
	showWorkerService := service.ShowWorkerService{}
	res := showWorkerService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteWorker(c *gin.Context) {
	deleteWorkerService := service.DeleteWorkerService{}
	res := deleteWorkerService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateWorker(c *gin.Context) {
	updateWorkerService := service.UpdateWorkerService{}
	if err := c.ShouldBind(&updateWorkerService); err == nil {
		res := updateWorkerService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
