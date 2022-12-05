package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//响应
func Response(c_cx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c_cx.JSON(httpStatus, gin.H{"code":code, "data":data, "msg":msg})
}
//登陆成功
func Success(c_cx *gin.Context, data gin.H, msg string) {
	Response(c_cx, http.StatusOK, 200, data, msg)
}

//func Fail(context *gin.Context, msg string, data gin.H) {
//	Response(context, http.StatusOK, 400, data,msg)
//}