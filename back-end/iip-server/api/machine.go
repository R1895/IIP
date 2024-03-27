package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateMachine(c *gin.Context) {
	createService := service.CreateMachineService{}
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListMachines(c *gin.Context) {
	listService := service.ListMachineService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ModifyMachineStatus(c *gin.Context) {
	modifyMachineStatus := service.ModifyMachineStatusService{}
	if err := c.ShouldBind(&modifyMachineStatus); err == nil {
		res := modifyMachineStatus.ModifyMachineStatus(c.Param("mid"),c.Param("sid"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
func ShowMachine(c *gin.Context) {
	showMachineService := service.ShowMachineService{}
	res := showMachineService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteMachine(c *gin.Context) {
	deleteMachineService := service.DeleteMachineService{}
	res := deleteMachineService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateMachine(c *gin.Context) {
	updateMachineService := service.UpdateMachineService{}
	if err := c.ShouldBind(&updateMachineService); err == nil {
		res := updateMachineService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// TODO 批量创建 接口
