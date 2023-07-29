package handler

import (
	"gin-mall/internal/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	hdl = &handler{}
	userHdl = &userHandler{handler: hdl}
}

type handler struct {
}

var hdl *handler
var userHdl *userHandler

func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*middleware.MyCustomClaims).UserId
}
