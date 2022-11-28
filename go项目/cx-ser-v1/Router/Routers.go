package Router

import (
	"cx/Controller"
	"cx/Mapper"
	"cx/Middleware"
	"github.com/gin-gonic/gin"
)

func CreateRoute(re *gin.Engine) *gin.Engine {
	re.Use(Middleware.CORSMiddleware())
	Mapper.AdminMapper(re)
	Mapper.EmployeeMapper(re) //员工 路由
	Mapper.DirectorMapper(re)
	re.GET("/info", Middleware.AuthMiddleware(), Controller.Info)
	return re
}
