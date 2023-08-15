package middleware

import (
	"github.com/gin-gonic/gin"
)

func AesMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//		body := []byte(`{
		//    "account" :"15838333913",
		//    "password":"12345678"
		//}`)
		//		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		c.Next()
	}
}
