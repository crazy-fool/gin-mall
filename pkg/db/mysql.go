package db

import (
	"gin-mall/pkg/config"
	"gin-mall/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func init() {
	db = newDB()
	newRedis()
}

var db *gorm.DB

func newDB() *gorm.DB {

	tdb, err := gorm.Open(mysql.Open(config.GetConfig().GetString("data.mysql.user")), getGormConfig())
	if err != nil {
		log.GetLog().Info("==================数据库初始化失败=======================")
		panic(err)
	}
	c, err := tdb.DB()
	if err != nil {
		log.GetLog().Info("==================数据库设置连接池失败=======================")
	}
	c.SetMaxIdleConns(10)
	c.SetMaxIdleConns(100)
	log.GetLog().Info("==================数据库初始化完成=======================")
	return tdb
}

func GetDB() *gorm.DB {
	return db
}

func getGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "hmh_",
			SingularTable: true,
		},
	}
}
