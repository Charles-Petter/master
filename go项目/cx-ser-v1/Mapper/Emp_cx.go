package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)
//员工路由
func EmployeeMapper_cx(re_cx *gin.Engine)  {
	re_cx.POST("/EmployeeBasic_cx", Controller.EmployeeBasic_cx)
	re_cx.POST("/Employee/SearchEmp_cx", Controller.EmployeeSearchEmp_cx)
}
