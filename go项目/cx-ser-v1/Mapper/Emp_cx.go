package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)
//员工路由
func EmployeeMapper_cx(re *gin.Engine)  {
	re.POST("/EmployeeBasic_cx", Controller.EmployeeBasic_cx)
	re.POST("/Employee/SearchEmp_cx", Controller.EmployeeSearchEmp_cx)
}
