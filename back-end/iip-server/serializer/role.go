package serializer

import (
	"iip/model"
)

type Role struct {
	RoleId      uint         `json:"role_id"`
	RoleName    string       `json:"role_name"`
	Description string       `json:"description"`
	Users       []User       `json:"users"`
	Permissions []Permission `json:"permissions"`
}

func BuildRole(role model.Role) Role {
	return Role{
		RoleId:      role.ID,
		RoleName:    role.RoleName,
		Description: role.Description,
		Users:       BuildUsers(role.Users),
		Permissions: BuildPermissions(role.Permissions),
	}
}

func BuildRoles(items []model.Role) (roles []Role) {
	for _, item := range items {
		role := BuildRole(item)
		roles = append(roles, role)
	}
	return roles
}
