package handler

import (
	"github.com/gin-gonic/gin"
)

type SpuHandler interface {
	Edit(ctx *gin.Context)
	List(ctx *gin.Context)
}
type spuHandler struct {
	*handler
}

func GetSpuHandler() SpuHandler {
	return spuHdl
}

func (s spuHandler) Edit(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s spuHandler) List(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
