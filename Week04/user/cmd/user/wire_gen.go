// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/biz"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/data"
	"github.com/zepeng-jiang/Go-000/Week04/user/internal/service"
)

// Injectors from wire.go:

func InitializeUser(db *gorm.DB, cache *redis.Client) *service.UserService {
	userRepository := data.NewUserRepository(db, cache)
	userCase := biz.NewUserBIZ(userRepository)
	userService := service.NewUserService(userCase)
	return userService
}