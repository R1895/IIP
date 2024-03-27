package serializer

import (
	"iip/model"
	"time"
)

type User struct {
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Avatar    string `json:"avatar"`
	// PasswordDigest   string    `json:"password_digest"`
	LastLogin        time.Time `json:"last_login"`
	IsActive         uint      `json:"is_active"`
	RegistrationDate time.Time `json:"registration_date"`
	Roles            []Role    `json:"roles"`
	Messages         []Message `json:"messages"`
}


func BuildUser(user model.User) User {
	return User{
		UserID:    user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		Telephone: user.Telephone,
		// PasswordDigest:   user.PasswordDigest,
		LastLogin:        user.LastLogin,
		IsActive:         user.IsActive,
		RegistrationDate: user.RegistrationDate,
		Roles:            BuildRoles(user.Roles),
		Messages:         BuildMessages(user.Messages),
	}
}

func BuildUsers(items []model.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
