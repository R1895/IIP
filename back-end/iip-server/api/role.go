package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreateRole(c *gin.Context) {
	createService := service.CreateRoleService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListRoles(c *gin.Context) {
	listService := service.ListRoleService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowRole(c *gin.Context) {
	showRoleService := service.ShowRoleService{}
	res := showRoleService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteRole(c *gin.Context) {
	deleteRoleService := service.DeleteRoleService{}
	res := deleteRoleService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateRole(c *gin.Context) {
	updateRoleService := service.UpdateRoleService{}
	if err := c.ShouldBind(&updateRoleService); err == nil {
		res := updateRoleService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
