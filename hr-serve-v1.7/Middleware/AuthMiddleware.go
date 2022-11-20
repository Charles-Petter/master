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
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":401,
				"msg":"权限不足",
			})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := Global.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":401,
				"msg":"权限不足",
			})
			context.Abort()
			return
		}

		// 验证通过后获取claim的userid
		userID := claims.UserId
		var admin Model.Employee
		Global.Db.First(&admin, userID)

		// 用户不存在
		if userID == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":401,
				"msg":"权限不足",
			})
			context.Abort()
			return
		}

		// 存在
		context.Set("user", admin)
		context.Next()
	}
}
