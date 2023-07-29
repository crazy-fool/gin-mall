package db

import (
	"gin-mall/pkg/config"
	"gin-mall/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	newDB()
	newRedis()
}

var db *gorm.DB

func newDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.GetConfig().GetString("data.mysql.user")), &gorm.Config{})
	if err != nil {
		log.GetLog().Info("==================数据库初始化失败=======================")
		panic(err)
	}
	log.GetLog().Info("==================数据库初始化完成=======================")
	return db
}

func GetDB() *gorm.DB {
	return db
}
