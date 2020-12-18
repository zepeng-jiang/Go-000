package biz

import (
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/data"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/entity"
)

type UserRepo interface {
	SaveUser(user *entity.User) error
}

type UserCase struct {
	repo UserRepo
}

func NewUserBIZ(repo *data.UserRepository) *UserCase {
	return &UserCase{repo: repo}
}

func (b *UserCase) CreateUser(user *entity.User) error {
	// 检查user字段
	if err := user.Check(); err != nil {
		return err
	}
	return b.repo.SaveUser(user)
}
