package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"microservice/user/rpc/internal/config"
	"microservice/user/rpc/models"
)

type ServiceContext struct {
	Config   config.Config
	DbEngine *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "douyin_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	// 根据结构体自动创建表
	db.AutoMigrate(&models.User{})

	return &ServiceContext{
		Config:   c,
		DbEngine: db,
	}
}
