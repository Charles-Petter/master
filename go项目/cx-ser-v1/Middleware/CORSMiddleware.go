package Middleware

import (
	"github.com/gin-gonic/gin"
)
//跨域资源共享中间件
func CORS_cx() gin.HandlerFunc {
	return func(c_cx *gin.Context) {
		//请求方法赋值给变量
		method_cx := c_cx.Request.Method
		c_cx.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		c_cx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c_cx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c_cx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c_cx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")//设置为true，允许ajax异步请求带cookie信息
		//放行所有OPTIONS方法
		if method_cx == "OPTIONS" {
			c_cx.AbortWithStatus(200)
		} else {
			c_cx.Next()
		}
	}
}