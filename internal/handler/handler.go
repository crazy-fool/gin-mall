package handler

import (
	"gin-mall/pkg/auth"
	"github.com/gin-gonic/gin"
)

func init() {
	hdl = &handler{}
	userHdl = &userHandler{handler: hdl}
	categoryHdl = &categoryHandler{handler: hdl}
	spuHdl = &spuHandler{handler: hdl}
}

type handler struct {
}

var hdl *handler
var userHdl *userHandler
var categoryHdl *categoryHandler
var spuHdl *spuHandler

func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*auth.MyCustomClaims).UserId
}
