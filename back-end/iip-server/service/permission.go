package service

import (
 
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
)

// 创建权限
type CreatePermissionService struct {
	PermissionName string `form:"permission_name" json:"permission_name" binding:"required,min=2,max=100"`
	Description    string `form:"description" json:"description"`
	Roles          []int  `form:"roles" json:"roles"`
}

// 更新权限
type UpdatePermissionService struct {
	ID             uint   `form:"id" json:"id"`
	PermissionName string `form:"permission_name" json:"permission_name" binding:"required,min=2,max=100"`
	Description    string `form:"description" json:"description"`
	Roles          []int  `form:"roles" json:"roles"`
}

// 查询权限，分页
type ListPermissionService struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start" binding:"required"`
}

// 删除权限
type DeletePermissionService struct {
}

// 展示权限详情
type ShowPermissionService struct {
}

func (service *CreatePermissionService) Create() *serializer.Response {
	code := e.SUCCESS
	var permission model.Permission
	var count int64
	model.DB.Model(&model.Workshop{}).Where("permission_name=?", service.PermissionName).First(&permission).Count(&count)
	permission = model.Permission{
		PermissionName: service.PermissionName,
		Description:    service.Description,
	}
	if err := model.DB.Create(&permission).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildPermission(permission),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListPermissionService) List() serializer.Response {
	var permissions []model.Permission
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 预加载和分页
	tx := model.DB.Model(model.Permission{})

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&permissions)

	return serializer.BuildListResponse(
		serializer.BuildPermissions(permissions), uint(total))
}

func (service *ShowPermissionService) Show(id string) serializer.Response {
	var permission model.Permission
	code := e.SUCCESS
	err := model.DB.Model(model.Permission{}).First(&permission, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildPermission(permission),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeletePermissionService) Delete(id string) serializer.Response {
	var permission model.Permission
	code := e.SUCCESS
	err := model.DB.First(&permission, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&permission).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *UpdatePermissionService) Update(id string) serializer.Response {
	var permission model.Permission
	model.DB.Model(model.Permission{}).Where("id = ?", id).First(&permission)
	// permission.Roles = service.Roles
	if service.PermissionName != "" {
		permission.PermissionName = service.PermissionName
	}

	code := e.SUCCESS
	err := model.DB.Save(&permission).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "修改成功",
	}
}
