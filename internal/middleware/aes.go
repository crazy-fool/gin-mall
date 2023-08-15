package middleware

import (
	"bytes"
	"gin-mall/pkg/config"
	"gin-mall/pkg/helper/aesutil"
	"gin-mall/pkg/helper/resp"
	"gin-mall/pkg/log"
	"github.com/gin-gonic/gin"
	"io"
)

var aesRequestUtil aesutil.UtilAes

func init() {
	key := "12345678901234567890123456789012"
	aesRequestUtil = aesutil.NewPkcs7Cbc([]byte(key))
}

// AesMiddleware 请求参数加密
func AesMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.GetConfig().GetString("env") != "prod" {
			c.Next()
			return
		}
		data, err := c.GetRawData()
		if err != nil {
			log.GetLog().Error("[aes数据加密] error:" + err.Error())
			log.GetLog().Error("[aes数据加密] error data:" + string(data))
			resp.ResponseError(c, resp.DecryptedFailed)
			c.Abort()
			return
		}
		log.GetLog().Debug("[aes数据加密] data" + string(data))
		decrypted, err := aesRequestUtil.Base64Decrypted(string(data))
		if err != nil {
			log.GetLog().Error("[aes数据加密] error" + err.Error())
			resp.ResponseError(c, resp.DecryptedFailed)
			c.Abort()
			return
		}
		log.GetLog().Debug("[aes数据加密] 解密data" + decrypted)
		c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(decrypted)))
		c.Next()
	}
}
