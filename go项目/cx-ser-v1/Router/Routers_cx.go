package Router

import (
	"cx/Controller"
	"cx/Mapper"
	"cx/Middleware"
	"github.com/gin-gonic/gin"
)

func CreateRoute(r_cx *gin.Engine) *gin.Engine {
	r_cx.Use(Middleware.CORSMiddleware())
	Mapper.AdminMapper(r_cx)
	Mapper.EmployeeMapper(r_cx) //员工 路由
	Mapper.DirectorMapper(r_cx)
	r_cx.GET("/info", Middleware.ShowMiddleware_cx(), Controller.Info)
	return r_cx
}
