package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)
//员工路由
func EmployeeMapper(re *gin.Engine)  {
	re.POST("/EmployeeBasicByEmployee", Controller.EmployeeBasicByEmployee)
	//re.POST("/SalaryBasicByEmp", Controller.SalaryBasicByEmp)

	re.POST("/EmployeeBasic/SearchByEmployee", Controller.EmployeeBasicSearchByEmployee)
}
