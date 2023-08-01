package repository

import (
	"fmt"
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
	spuRepo = &spuRepository{Repository: repo}
	fmt.Println(repo.db)
	fmt.Printf("%p", repo.db)
}

var repo *Repository
var userRepo *userRepository
var spuRepo *spuRepository

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}
