package main

import (
	"fmt"
	_ "gin-mall/internal/repository"
	"gin-mall/internal/server"
	_ "gin-mall/internal/service"
	"gin-mall/pkg/config"
	_ "gin-mall/pkg/db"
	"gin-mall/pkg/http"
	"gin-mall/pkg/log"
	"go.uber.org/zap"
)

func main() {
	logger := log.GetLog()
	conf := config.GetConfig()
	app := server.NewServerHTTP()
	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
}
