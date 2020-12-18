package main

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/zepeng-jiang/Go-000/Week04/user/api/v1"
	"github.com/zepeng-jiang/Go-000/Week04/user/configs"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	// 初始化配置信息
	conf, err := configs.InitConfig()
	if err != nil {
		log.Fatalf("Init config failed. err: %s\n", err.Error())
	}

	// 初始化数据库
	db, err1 := initDB(conf)
	if err1 != nil {
		log.Fatalf("Init database failed. err: %+v\n", err)
	}

	// 初始化缓存
	cache := initCache(conf)

	// 依赖注入
	service := InitializeUser(db, cache)

	listen, err2 := net.Listen("tcp", "127.0.0.1:"+conf.Server.Port)
	if err2 != nil {
		os.Exit(1)
	}

	server := grpc.NewServer()
	v1.RegisterUserServer(server, service)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("RPC server listen failed. err: %s\n", err.Error())
	}
}

func initDB(conf *configs.Config) (*gorm.DB, error) {
	db, err := configs.NewDBEngine(conf)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initCache(conf *configs.Config) *redis.Client {
	return nil
}
