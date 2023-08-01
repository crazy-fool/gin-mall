package handler

import (
	"gin-mall/internal/params"
	"gin-mall/pkg/helper/resp"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	Edit(ctx *gin.Context)
}
type categoryHandler struct {
	*handler
}

func GetCategoryHandler() CategoryHandler {
	return categoryHdl
}

func (h *categoryHandler) Edit(ctx *gin.Context) {
	var param params.EditCategoryParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}
}
