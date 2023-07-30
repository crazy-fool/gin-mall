package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: 0, Message: "success", Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func ResponseError(ctx *gin.Context, code int) {
	msg, ok := responseMessage[code]
	if !ok {
		msg = CommonErrorMessage
	}
	resp := response{Code: code, Message: msg, Data: map[string]string{}}
	ctx.JSON(http.StatusOK, resp)
}

func ResponseErrorWithMsg(ctx *gin.Context, code int, ext string) {
	msg, ok := responseMessage[code]
	if !ok {
		msg = CommonErrorMessage
	}
	resp := response{Code: code, Message: msg + ":" + ext, Data: map[string]string{}}
	ctx.JSON(http.StatusOK, resp)
}
