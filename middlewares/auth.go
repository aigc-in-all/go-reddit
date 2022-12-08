package middlewares

import (
	"github.com/gin-gonic/gin"
	"goreddit/controller"
	"goreddit/pkg/jwt"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Authorization: Bearer xxxxxx.xxxxxx
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userId信息保存到请求的上下文c上
		c.Set(controller.CtxUserIDKey, mc.UserId)
		c.Next()
	}
}
