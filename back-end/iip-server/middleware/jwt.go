package middleware

import (
	"github.com/gin-gonic/gin"
	"iip/pkg/e"
	"iip/pkg/util"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
