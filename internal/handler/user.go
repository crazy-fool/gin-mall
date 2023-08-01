package handler

import (
	"gin-mall/internal/params"
	"gin-mall/internal/service"
	"gin-mall/pkg/helper/resp"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type userHandler struct {
	*handler
}

func GetUserHandler() UserHandler {
	return userHdl
}

func (h *userHandler) Register(ctx *gin.Context) {
	req := new(params.RegisterParam)
	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.ResponseError(ctx, resp.ParamError)
		return
	}

	if err := service.GetUserService().Register(ctx, req); err != nil {
		resp.ResponseError(ctx, resp.OpFailed)
		return
	}

	resp.HandleSuccess(ctx, nil)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var req params.LoginParam
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.ResponseErrorWithMsg(ctx, resp.CommonError, errors.Wrap(err, "invalid request").Error())
		return
	}

	token, err := service.GetUserService().Login(ctx, &req)
	if err != nil {
		resp.ResponseErrorWithMsg(ctx, resp.CommonError, err.Error())
		return
	}

	resp.HandleSuccess(ctx, gin.H{
		"accessToken": token,
	})
}

func (h *userHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		resp.ResponseError(ctx, resp.CustomerNotLogin)
		return
	}

	user, err := service.GetUserService().GetProfile(ctx, userId)
	if err != nil {
		resp.ResponseError(ctx, resp.ResourceNotFound)
		return
	}

	resp.HandleSuccess(ctx, user)
}

func (h *userHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req params.UpdateProfileParam
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.ResponseErrorWithMsg(ctx, resp.ParamError, err.Error())
		return
	}

	if err := service.GetUserService().UpdateProfile(ctx, userId, &req); err != nil {
		resp.ResponseErrorWithMsg(ctx, resp.OpFailed, err.Error())
		return
	}

	resp.HandleSuccess(ctx, nil)
}
