package Router

import (
	"github.com/gin-gonic/gin"
	"m/Controller"
	"m/Mapper"
	"m/Middleware"
)

func CreateRoute(re *gin.Engine) *gin.Engine {
	re.Use(Middleware.CORSMiddleware())
	Mapper.AdminMapper(re)
	Mapper.EmployeeMapper(re) //员工 路由
	Mapper.DirectorMapper(re)
	//Mapper.AdministratorMapper(re)
	//Mapper.DesignerMapper(re)
	//Mapper.OperatorMapper(re)
	//Mapper.ComponentMapper(re)
	re.GET("/info", Middleware.AuthMiddleware(), Controller.Info)
	return re
}
