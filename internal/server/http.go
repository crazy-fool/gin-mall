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
		middleware.AesMiddleware(),
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
	//noStrictAuthRouter := rootGroup.Group("/").Use(auth.NoStrictAuth())
	//{
	//	//noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	//}

	categoryHandler := handler.GetCategoryHandler()
	specHandler := handler.GetSpecHandler()
	// Strict permission routing group
	strictAuthRouter := rootGroup.Group("/").Use(auth.StrictAuth())
	{
		strictAuthRouter.POST("v1/category/edit", categoryHandler.Edit)
		strictAuthRouter.GET("v1/category/list", categoryHandler.List)
		strictAuthRouter.GET("v1/category/one-list", categoryHandler.OneList)
		strictAuthRouter.GET("v1/category/son-list", categoryHandler.SonList)

		strictAuthRouter.GET("v1/spec/group-list", specHandler.GroupList)
		strictAuthRouter.POST("v1/spec/group-edit", specHandler.EditSpecGroup)
		strictAuthRouter.GET("v1/spec/list", specHandler.SpecList)
		strictAuthRouter.POST("v1/spec/edit", specHandler.EditSpec)
	}
	return r
}
