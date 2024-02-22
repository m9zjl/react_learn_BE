package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域访问：cross  origin resource share
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		context.Header("Access-Control-Allow-Origin", origin)
		context.Header("Access-Control-Allow-Headers", "Content-Type, cache-control, Origin, Access-Control-Allow-Origin, AccessToken, X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "PUT, POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		// 允许放行OPTIONS请求
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
