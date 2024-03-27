package model

import (
	"github.com/jinzhu/gorm"
)

// 权限表
type Permission struct {
	gorm.Model
	PermissionName string
	Path           string
	Method         string
	Description    string
	Roles          []Role `gorm:"many2many:role_permissions;"`
	Field1         string
	Field2         string
}
