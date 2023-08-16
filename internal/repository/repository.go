package repository

import (
	"gin-mall/pkg/db"
	"gin-mall/pkg/log"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func init() {
	newRepo()
}

func newRepo() {
	repo = &Repository{
		db:  db.GetDB(),
		rdb: db.GetRedisClient(),
	}
	log.GetLog().Info("===========初始化repository============")
	userRepo = &userRepository{Repository: repo}
	spuRepo = &spuRepository{Repository: repo}
	categoryRepo = &categoryRepository{Repository: repo}
	specRepo = &specRepository{Repository: repo}
}

var repo *Repository
var userRepo *userRepository
var spuRepo *spuRepository
var categoryRepo *categoryRepository
var specRepo *specRepository

type Repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (r Repository) recordError(msg string, err error) {
	if err != nil {
		log.GetLog().Error("[数据库操作error]"+msg, zap.Error(err))
	}
}
