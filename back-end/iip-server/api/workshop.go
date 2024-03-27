package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateWorkshop(c *gin.Context) {
	createService := service.CreateWorkshopService{}
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListWorkshops(c *gin.Context) {
	listService := service.ListWorkshopService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowWorkshop(c *gin.Context) {
	showWorkshopService := service.ShowWorkshopService{}
	res := showWorkshopService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteWorkshop(c *gin.Context) {
	deleteWorkshopService := service.DeleteWorkshopService{}
	res := deleteWorkshopService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateWorkshop(c *gin.Context) {
	updateWorkshopService := service.UpdateWorkshopService{}
	if err := c.ShouldBind(&updateWorkshopService); err == nil {
		res := updateWorkshopService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func AddWorker(c *gin.Context) {
	addWorkerService := service.AddWorkerService{}
	if err := c.ShouldBind(&addWorkerService); err == nil {
		res := addWorkerService.AddWorker()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
func RemoveWorker(c *gin.Context) {
	removeWorkerService := service.RemoveWorkerService{}
	if err := c.ShouldBind(&removeWorkerService); err == nil {
		res := removeWorkerService.RemoveWorker()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

