package Middleware

import (
	"github.com/gin-gonic/gin"
	"m/Global"
	"m/Model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//GetHeader从请求头返回值
		tokenString := context.GetHeader("Authorization")
		//如果不为空或者 不是以Bearer 开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") { //字符串是否以 prefix开头
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			context.Abort()
			return
		}
		tokenString = tokenString[7:] //从下标 7 到 endIndex-1 下的元素创建为一个新的切片。

		token, claims, err := Global.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			context.Abort()
			return
		}

		// 验证通过后获取claim的userid
		userID := claims.UserId
		var admin Model.Employee
		Global.Db.First(&admin, userID) //First查询一条记录，根据主键ID排序(正序)

		// 用户不存在
		if userID == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			context.Abort()
			return
		}

		// 存在
		context.Set("user", admin) //set储存新的键值对
		context.Next()
	}
}
