package Router

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
	"cx/Mapper"
	"cx/Middleware"

)

func CreateRoute(re *gin.Engine) *gin.Engine {
	//r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "OK")
	//})
	re.Use(Middleware.CORSMiddleware())
	Mapper.AdminMapper(re) //登录 路由
	Mapper.EmployeeMapper(re) //员工 路由
	Mapper.DirectorMapper(re)







	//Mapper.AdministratorMapper(re)
	//Mapper.DesignerMapper(re)
	//Mapper.OperatorMapper(re)
	//Mapper.ComponentMapper(re)
	re.GET("/info", Middleware.AuthMiddleware(), Controller.Info)
	return re
}
