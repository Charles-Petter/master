package Middleware

import (
	"github.com/gin-gonic/gin"
)
//跨域资源共享中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c_cx *gin.Context) {
		method := c_cx.Request.Method
		c_cx.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		c_cx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c_cx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c_cx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c_cx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//if c_cx.Request.Method == http.MethodOptions {
		if method == "OPTIONS" {
			c_cx.AbortWithStatus(200)
		} else {
			c_cx.Next()
		}
	}
}