package Router

import (
	"cx/Controller"
	"cx/Mapper"
	"cx/Middleware"
	"github.com/gin-gonic/gin"
)
//框架入口
func CreateRoute(r_cx *gin.Engine) *gin.Engine {
	//定义服务路由信息
	r_cx.Use(Middleware.CORS_cx())
	//登录路由
	Mapper.AdminMapper(r_cx)
	//员工 路由
	Mapper.EmployeeMapper(r_cx)
	//主管路由
	Mapper.DirectorMapper(r_cx)
	r_cx.GET("/info", Middleware.ShowMiddleware_cx(), Controller.Info)
	return r_cx
}
