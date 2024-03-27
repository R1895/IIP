package model

import (
	"github.com/jinzhu/gorm"
)

// 角色表
type Role struct {
	gorm.Model
	RoleName    string
	Description string
	Users       []User       `gorm:"many2many:user_roles;"`       // 定义user和role多对多关联
	Permissions []Permission `gorm:"many2many:role_permissions;"` // 定义permission和role多对多关联
	Field1      string
	Field2      string
}
