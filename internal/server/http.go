package server

import (
	"gin-mall/internal/handler"
	"gin-mall/internal/middleware"
	"gin-mall/pkg/auth"
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
	rootGroup := r.Group("/api/")
	userHandler := handler.GetUserHandler()
	// No route group has permission
	noAuthRouter := rootGroup.Group("/")
	{
		noAuthRouter.GET("/", func(ctx *gin.Context) {
			log.GetLog().WithContext(ctx).Info("hello")
			resp.HandleSuccess(ctx, map[string]interface{}{
				"message": "Hi mall",
			})
		})
		noAuthRouter.POST("v1/user/register", userHandler.Register)
		noAuthRouter.POST("v1/user/login", userHandler.Login)
	}
	// Non-strict permission routing group
	noStrictAuthRouter := rootGroup.Group("/").Use(auth.NoStrictAuth())
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// Strict permission routing group
	strictAuthRouter := rootGroup.Group("/").Use(auth.StrictAuth())
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}
	return r
}
