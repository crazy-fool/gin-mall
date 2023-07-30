package middleware

import (
	"gin-mall/pkg/config"
	"gin-mall/pkg/helper/resp"
	"gin-mall/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"regexp"
	"time"
)

type JWT struct {
	key []byte
}

type MyCustomClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func init() {
	jwtInstance = &JWT{key: []byte(config.GetConfig().GetString("security.jwt.key"))}
	log.GetLog().Info("=================初始化jwt==================")
}

var jwtInstance *JWT

func GetJwt() *JWT {
	return jwtInstance
}

func (j *JWT) GenToken(userId string, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			ID:        "",
			Audience:  []string{},
		},
	})

	// Sign and get the complete encoded token as a string using the key
	tokenString, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) ParseToken(tokenString string) (*MyCustomClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func StrictAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		logger := log.GetLog()
		if tokenString == "" {
			logger.WithContext(ctx).Warn("请求未携带token，无权限访问", zap.Any("data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}))
			resp.HandleError(ctx, 1, "no token")
			ctx.Abort()
			return
		}

		claims, err := GetJwt().ParseToken(tokenString)
		if err != nil {
			logger.WithContext(ctx).Error("token error", zap.Any("data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}))
			resp.HandleError(ctx, 1, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func NoStrictAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("accessToken")
		}
		if tokenString == "" {
			tokenString = ctx.Query("accessToken")
		}
		if tokenString == "" {
			ctx.Next()
			return
		}

		claims, err := GetJwt().ParseToken(tokenString)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, log.GetLog())
		ctx.Next()
	}
}

func recoveryLoggerFunc(ctx *gin.Context, logger *log.Logger) {
	userInfo := ctx.MustGet("claims").(*MyCustomClaims)
	logger.NewContext(ctx, zap.String("UserId", userInfo.UserId))
}
