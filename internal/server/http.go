package server

import (
	"gin-mall/internal/handler"
	"gin-mall/internal/middleware"
	"gin-mall/pkg/helper/resp"
	"gin-mall/pkg/log"
	"github.com/gin-gonic/gin"
)

func NewServerHTTP() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(),
		middleware.RequestLogMiddleware(),
		//middleware.SignMiddleware(),
	)
	userHandler := handler.GetUserHandler()
	// No route group has permission
	noAuthRouter := r.Group("/")
	{
		noAuthRouter.GET("/", func(ctx *gin.Context) {
			log.GetLog().WithContext(ctx).Info("hello")
			resp.HandleSuccess(ctx, map[string]interface{}{
				"message": "Hi mall",
			})
		})
		noAuthRouter.POST("/register", userHandler.Register)
		noAuthRouter.POST("/login", userHandler.Login)
	}
	// Non-strict permission routing group
	noStrictAuthRouter := r.Group("/").Use(middleware.NoStrictAuth())
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// Strict permission routing group
	strictAuthRouter := r.Group("/").Use(middleware.StrictAuth())
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}
	return r
}
