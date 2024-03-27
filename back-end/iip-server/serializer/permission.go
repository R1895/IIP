package serializer

import (
	"iip/model"
)

type Permission struct {
	PermissionID   uint   `json:"permission_id"`
	PermissionName string `json:"permission_name"`
	Path           string `json:"path"`
	Method         string `json:"method"`
	Description    string `json:"description"`
	Roles          []Role `json:"roles"`
}

func BuildPermission(permission model.Permission) Permission {
	return Permission{
		PermissionID:   permission.ID,
		PermissionName: permission.PermissionName,
		Description:    permission.Description,
		Path:           permission.Path,
		Method:         permission.Method,
		Roles:          BuildRoles(permission.Roles),
	}
}

func BuildPermissions(items []model.Permission) (permissions []Permission) {
	for _, item := range items {
		permission := BuildPermission(item)
		permissions = append(permissions, permission)
	}
	return permissions
}
