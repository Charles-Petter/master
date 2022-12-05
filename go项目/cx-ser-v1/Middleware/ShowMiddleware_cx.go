package Middleware

import (
	"github.com/gin-gonic/gin"

	"cx/Global"

	"cx/Model"
	"net/http"
	"strings"
)
//授权中间件
func ShowMiddleware_cx() gin.HandlerFunc {
	return func(c_cx *gin.Context) {
		//GetHeader从请求头返回值
		tokenString_cx := c_cx.GetHeader("Authorization")
		//如果不为空或者 不是以Bearer 开头
		if tokenString_cx == "" || !strings.HasPrefix(tokenString_cx, "Bearer ") { //字符串是否以 prefix开头
			c_cx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c_cx.Abort()
			return
		}
		tokenString_cx = tokenString_cx[7:]
		token_cx, claims, err := Global.ParseToken(tokenString_cx)
		if err != nil || !token_cx.Valid {
			c_cx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c_cx.Abort()
			return
		}
		// 验证通过后获取claim的userid
		userID_cx := claims.UserId_cx
		var loginUser_cx Model.Employee_cx
		Global.Db.First(&loginUser_cx, userID_cx) //First查询一条记录，根据主键ID排序(正序)
		// 用户不存在
		if userID_cx == "" {
			c_cx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c_cx.Abort()
			return
		}
		// 存在
		c_cx.Set("user", loginUser_cx) //set储存新的键值对
		c_cx.Next()//中间件使用的标准用法，表示继续处理请求。
	}
}
