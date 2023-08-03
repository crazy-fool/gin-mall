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
	OneList(ctx *gin.Context)
	SonList(ctx *gin.Context)
}
type categoryHandler struct {
	*handler
}

func GetCategoryHandler() CategoryHandler {
	return categoryHdl
}

// Edit 分类编辑
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

// List 带分页的分类列表
func (h *categoryHandler) List(ctx *gin.Context) {
	var param params.CategoryListParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}
	ret := service.GetCategoryService().GetPageList(ctx, &param)
	resp.HandleSuccess(ctx, ret)
}

// OneList 一级分类列表
func (h *categoryHandler) OneList(ctx *gin.Context) {
	param := &params.CategoryListParam{
		ParentId: new(uint),
	}
	ret := service.GetCategoryService().GetList(ctx, param)
	resp.HandleSuccess(ctx, ret)
}

// SonList 一级分类列表
func (h *categoryHandler) SonList(ctx *gin.Context) {
	var param params.CategoryListParam
	if err := ctx.ShouldBind(&param); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}

	log.GetLog().Info("查询分类", zap.Any("param", param))

	ret := service.GetCategoryService().GetList(ctx, &param)
	resp.HandleSuccess(ctx, ret)
}
