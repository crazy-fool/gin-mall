package middleware

import (
	"gin-mall/pkg/config"
	"gin-mall/pkg/helper/md5"
	"gin-mall/pkg/helper/resp"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
)

func SignMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requiredHeaders := []string{"Timestamp", "Nonce", "Sign", "App-Version"}

		for _, header := range requiredHeaders {
			value, ok := ctx.Request.Header[header]
			if !ok || len(value) == 0 {
				resp.ResponseError(ctx, resp.SignFailed)
				ctx.Abort()
				return
			}
		}
		conf := config.GetConfig()
		data := map[string]string{
			"AppKey":     conf.GetString("security.api_sign.app_key"),
			"Timestamp":  ctx.Request.Header.Get("Timestamp"),
			"Nonce":      ctx.Request.Header.Get("Nonce"),
			"AppVersion": ctx.Request.Header.Get("App-Version"),
		}

		var keys []string
		for k := range data {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

		var str string
		for _, k := range keys {
			str += k + data[k]
		}
		str += conf.GetString("security.api_sign.app_security")

		if ctx.Request.Header.Get("Sign") != strings.ToUpper(md5.Md5(str)) {
			resp.ResponseError(ctx, resp.SignFailed)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
