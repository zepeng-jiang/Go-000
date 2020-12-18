package entity

import (
	"errors"
	"github.com/jinzhu/gorm"
)

// User
type User struct {
	gorm.Model
	Name     string `gorm:"column:name" json:"name" form:"name"`
	Password string `gorm:"column:password" json:"password" form:"password"`
}

// CheckName 检查字段
func (u *User) Check() error {
	if u.Name == "" {
		return errors.New("user name must be not empty")
	}
	if u.Password == "" {
		return errors.New("user password must be not empty")
	} else if len(u.Password) < 8 {
		return errors.New("user password length must be greater than zero")
	}
	return nil
}