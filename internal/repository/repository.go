package repository

import (
	"gin-mall/pkg/db"
	"gin-mall/pkg/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func init() {
	newRepo()
}

func newRepo() {
	log.GetLog().Info("===========初始化repository============")
	repo = &Repository{
		db:  db.GetDB(),
		rdb: db.GetRedisClient(),
	}
	userRepo = &userRepository{Repository: repo}
}

var repo *Repository
var userRepo *userRepository

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}
