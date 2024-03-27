package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName         string `gorm:"unique"`
	Email            string `gorm:"unique"`
	Telephone        string `gorm:"unique"`
	Avatar           string
	PasswordDigest   string
	LastLogin        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	IsActive         uint
	RegistrationDate time.Time `gorm:"default:NULL"`
	Roles            []Role    `gorm:"many2many:user_roles;"` // 定义多对多关联
	Messages         []Message
	Field1           string
	Field2           string
	Field3           string
	Field4           string
}

const (
	PassWordCost = 12 //密码加密难度
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
