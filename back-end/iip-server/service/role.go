package service

import (
	"fmt"
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
)

// 创建角色
type CreateRoleService struct {
	RoleName    string `form:"role_name" json:"role_name" binding:"required,min=2,max=100"`
	Description string `form:"description" json:"description"`
}

// 更新角色
type UpdateRoleService struct {
	ID          uint   `form:"id" json:"id"`
	RoleName    string `form:"role_name" json:"role_name" binding:"required,min=2,max=100"`
	Description string `form:"description" json:"description"`
}

// 查询角色，分页
type ListRoleService struct {
	Limit       int    `form:"limit" json:"limit"`
	Start       int    `form:"start" json:"start" binding:"required"`
	RoleName    string `form:"role_name" json:"role_name"`
	RoleId      string `form:"role_id" json:"role_id"`
}

// 删除角色
type DeleteRoleService struct {
}

// 展示角色详情
type ShowRoleService struct {
}

func (service *CreateRoleService) Create() *serializer.Response {
	code := e.SUCCESS
	var role model.Role
	var count int64
	model.DB.Model(&model.Role{}).Where("role_name=?", service.RoleName).First(&role).Count(&count)
	//表单验证
	if count == 1 {
		code = e.ErrorExistRole
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	role = model.Role{
		RoleName:    service.RoleName,
		Description: service.Description,
	}

	if err := model.DB.Create(&role).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildRole(role),
		Msg:    e.GetMsg(code),
	}
}
func (service *ListRoleService) List() serializer.Response {
	var roles []model.Role
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.Role{}).Preload("Users").Preload("Permissions")
	if service.RoleName != "" {
		tx = tx.Where(fmt.Sprintf("role_name like '%%%s%%'", service.RoleName))
	}
	if service.RoleId != "" {
		tx = tx.Where("role_id =?", service.RoleId)
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&roles)

	return serializer.BuildListResponse(serializer.BuildRoles(roles), uint(total))
}

func (service *ShowRoleService) Show(id string) serializer.Response {
	// TODO 展示网络详情
	var role model.Role
	code := e.SUCCESS
	err := model.DB.Preload("Users").Preload("Permissions").First(&role, id).Error
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
		Data:   serializer.BuildRole(role),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteRoleService) Delete(id string) serializer.Response {
	var role model.Role
	code := e.SUCCESS
	err := model.DB.First(&role, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&role).Error
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

func (service *UpdateRoleService) Update(id string) serializer.Response {
	var role model.Role
	model.DB.Model(model.Role{}).Where("id = ?", id).First(&role)
	if service.RoleName != "" {
		role.RoleName = service.RoleName
	}
	if service.Description != "" {
		role.Description = service.Description
	}

	code := e.SUCCESS
	err := model.DB.Save(&role).Error
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
