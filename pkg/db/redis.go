package db

import (
	"context"
	"fmt"
	"gin-mall/pkg/config"
	"gin-mall/pkg/log"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb *redis.Client

func newRedis() *redis.Client {
	conf := config.GetConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("data.redis.addr"),
		Password: conf.GetString("data.redis.password"),
		DB:       conf.GetInt("data.redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.GetLog().Info("==================redis数据库初始化失败=======================")
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}
	log.GetLog().Info("==================redis数据库初始化成功=======================")
	return rdb
}
func GetRedisClient() *redis.Client {
	return rdb
}
