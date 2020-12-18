// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/biz"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/data"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/service"
)

// InitializeEvent 声明injector的函数签名
func InitializeUser(db *gorm.DB, cache *redis.Client) *service.UserService {
	wire.Build(service.NewUserService, biz.NewUserBIZ, data.NewUserRepository)
	return &service.UserService{}
}