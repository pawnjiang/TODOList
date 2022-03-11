package middleware

import (
	"api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code uint32
		code = 200
		token := context.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				code = 401
			}
		}
		if code != 200 {
			context.JSON(500, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
