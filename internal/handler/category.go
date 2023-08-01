package handler

import (
	"gin-mall/internal/params"
	"gin-mall/internal/service"
	"gin-mall/pkg/helper/resp"
	"gin-mall/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryHandler interface {
	Edit(ctx *gin.Context)
	List(ctx *gin.Context)
}
type categoryHandler struct {
	*handler
}

func GetCategoryHandler() CategoryHandler {
	return categoryHdl
}

type Test struct {
	Name string
	Son
}
type Son struct {
	Id int
}

func (h *categoryHandler) Edit(ctx *gin.Context) {
	var param params.CategoryEditParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}
	if err := service.GetCategoryService().Edit(ctx, &param); err != nil {
		log.GetLog().Info("操作失败", zap.Error(err))
		resp.ResponseError(ctx, resp.OpFailed)
		return
	}
	resp.HandleSuccess(ctx, nil)
}

func (h *categoryHandler) List(ctx *gin.Context) {
	var param params.CategoryListParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}
	ret := service.GetCategoryService().GetList(ctx, &param)
	resp.HandleSuccess(ctx, ret)
}
