package handler

import (
	"gin-mall/internal/params"
	"gin-mall/internal/service"
	"gin-mall/pkg/helper/resp"
	"gin-mall/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SpecHandler interface {
	EditSpecGroup(ctx *gin.Context)
	GroupList(ctx *gin.Context)
	EditSpec(ctx *gin.Context)
	SpecList(ctx *gin.Context)
}

type specHandler struct {
	*handler
}

func GetSpecHandler() SpecHandler {
	return specHdl
}

func (s specHandler) EditSpecGroup(ctx *gin.Context) {
	var param params.GroupEditParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}
	if err := service.GetSpecService().EditGroup(ctx.Request.Context(), &param); err != nil {
		log.GetLog().Info("操作失败", zap.Error(err))
		resp.ResponseError(ctx, resp.OpFailed)
		return
	}
	resp.HandleSuccess(ctx, nil)
}

func (s specHandler) GroupList(ctx *gin.Context) {
	param := new(params.GroupListParam)
	if err := ctx.ShouldBind(param); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}
	ret := service.GetSpecService().GetGroupList(ctx.Request.Context(), param)
	resp.HandleSuccess(ctx, ret)
}

func (s specHandler) EditSpec(ctx *gin.Context) {
	var param params.SpecEditParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}
	if err := service.GetSpecService().EditSpec(ctx.Request.Context(), &param); err != nil {
		log.GetLog().Info("操作失败", zap.Error(err))
		resp.ResponseError(ctx, resp.OpFailed)
		return
	}
	resp.HandleSuccess(ctx, nil)
}

func (s specHandler) SpecList(ctx *gin.Context) {
	param := new(params.SpecListParam)
	if err := ctx.ShouldBind(param); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}
	ret := service.GetSpecService().GetSpecList(ctx.Request.Context(), param)
	resp.HandleSuccess(ctx, ret)
}
