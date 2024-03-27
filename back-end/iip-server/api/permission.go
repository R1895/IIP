package api

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/util"
	"iip/service"
)

func CreatePermission(c *gin.Context) {
	createService := service.CreatePermissionService{}
	// 鉴权
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ListPermissions(c *gin.Context) {
	listService := service.ListPermissionService{}

	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowPermission(c *gin.Context) {
	showPermissionService := service.ShowPermissionService{}
	res := showPermissionService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeletePermission(c *gin.Context) {
	deletePermissionService := service.DeletePermissionService{}
	res := deletePermissionService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdatePermission(c *gin.Context) {
	updatePermissionService := service.UpdatePermissionService{}
	if err := c.ShouldBind(&updatePermissionService); err == nil {
		res := updatePermissionService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
