package handler

import "github.com/gin-gonic/gin"

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
	//TODO implement me
	panic("implement me")
}

func (s specHandler) GroupList(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s specHandler) EditSpec(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s specHandler) SpecList(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
