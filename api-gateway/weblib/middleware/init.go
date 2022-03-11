package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//将服务保存在gin.key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["userService"] = service[0]
		context.Keys["taskService"] = service[1]
		context.Next()
	}
}

//错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.JSON(200, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", err),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
