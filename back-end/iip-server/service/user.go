package service

import (
	"fmt"
	"iip/model"
	"iip/pkg/e"
	"iip/pkg/util"
	"iip/serializer"
	"time"

	"github.com/jinzhu/gorm"
)

// UserService 用户注册服务
type UserService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" `
	Password  string `form:"password" json:"password"`
	Email     string `form:"email" json:"email"`
	Telephone string `form:"telephone" json:"telephone"`
}

func (service *UserService) Register() *serializer.Response {
	code := e.SUCCESS
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	//表单验证
	if count == 1 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.UserName = service.UserName
	user.Email = service.Email
	user.Telephone = service.Telephone

	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登陆函数
func (service *UserService) Login() serializer.Response {
	var user model.User
	code := e.SUCCESS
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			util.LogrusObj.Info(err)
			code = e.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

type ShowUserService struct {
}

type DeleteUserService struct {
}

type UpdateUserService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=2,max=100"`
	Email     string `form:"email" json:"email"`
	Telephone string `form:"telephone" json:"telephone"`
}

type ListUserService struct {
	Limit    int    `form:"limit" json:"limit" binding:"required"`
	Start    int    `form:"start" json:"start"`
	UserName string `form:"user_name" json:"user_name"`
}

type CreateUserService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=2,max=100"`
	Password  string `form:"password" json:"password"`
	Email     string `form:"email" json:"email"`
	Telephone string `form:"telephone" json:"telephone"`
}

func (service *CreateUserService) Create() *serializer.Response {
	var user model.User
	code := e.SUCCESS

	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)

	if count == 1 {
		code = e.ErrorExist
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 添加角色
	user = model.User{
		UserName:         service.UserName,
		Email:            service.Email,
		Telephone:        service.Telephone,
		LastLogin:        time.Now(),
		RegistrationDate: time.Now(),
	}

	if err := model.DB.Create(&user).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

func (service *ShowUserService) Show(id string) serializer.Response {
	var user model.User
	code := e.SUCCESS
	err := model.DB.First(&user, id).Error
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
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteUserService) Delete(id string) serializer.Response {
	var user model.User
	code := e.SUCCESS
	err := model.DB.First(&user, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&user).Error
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

func (service *UpdateUserService) Update(id string) serializer.Response {
	var user model.User
	model.DB.Model(model.User{}).Where("id = ?", id).First(&user)
	if service.UserName != "" {
		user.UserName = service.UserName
	}
	if service.Email != "" {
		user.Email = service.Email
	}
	if service.Telephone != "" {
		user.Telephone = service.Telephone
	}
	code := e.SUCCESS
	err := model.DB.Save(&user).Error
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

func (service *ListUserService) List() serializer.Response {
	var users []model.User
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	// 分页
	tx := model.DB.Model(model.User{})
	if service.UserName != "" {
		tx = tx.Where(fmt.Sprintf("user_name like '%%%s%%'", service.UserName))
	}

	tx.Count(&total).Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&users)

	return serializer.BuildListResponse(serializer.BuildUsers(users), uint(total))
}

// 修改用户角色
