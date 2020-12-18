package service

import (
	"context"
	"github.com/pkg/errors"
	v1 "github.com/zepeng-jiang/Go-000/Week04/user/api/v1"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/biz"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/entity"
)

type UserService struct {
	biz *biz.UserCase
}

func NewUserService(biz *biz.UserCase) *UserService {
	return &UserService{biz: biz}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.UserRequest) (*v1.UserResponse, error) {
	// DTO => DO
	user := &entity.User{}
	user.Name = req.Name
	user.Password = req.Password

	err := s.biz.CreateUser(user)
	if err != nil {
		return nil, errors.Wrap(err, "create user failed")
	}
	return &v1.UserResponse{Message: "create user success"}, nil
}
