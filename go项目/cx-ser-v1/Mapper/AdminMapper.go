package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)

func AdminMapper(re *gin.Engine) {
	//POST请求路由              路径         hander
	re.POST("/AdminLogin", Controller.AdminLogin)
	re.POST("/EmployeeBasic", Controller.EmployeeBasic)
	re.POST("/EmployeeBasic/Update", Controller.EmployeeBasicUpdate)
	re.POST("/EmployeeBasic/Add", Controller.EmployeeBasicAdd)
	re.POST("/Edit", Controller.Edit)
	re.POST("/Department/Basic", Controller.DepartmentBasic)
	re.POST("/SearchEmpSalary", Controller.SearchEmpSalary)
	re.POST("/Department/Init", Controller.DepartmentInit)
	re.POST("/EmpSalary/Init", Controller.EmpSalaryInit)
	re.POST("/Department/Update", Controller.DepartmentUpdate)
	re.POST("/SearchByDepartmentNumber", Controller.SearchByDepartmentNumber)
}
