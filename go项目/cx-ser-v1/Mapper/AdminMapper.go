package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)

func AdminMapper(re *gin.Engine) {
	//POST请求路由              路径         hander
	//身份信息验证
	re.POST("/AdminLogin", Controller.VerifyLogin_cx)
	//查询员工信息
	re.POST("/EmployeeBasic", Controller.EmployeeAll_cx)
	//修改
	re.POST("/EmployeeBasic/Update", Controller.EmployeeBasicUpdate)
	//添加
	re.POST("/EmployeeBasic/Add", Controller.EmployeeBasicAdd)
	//编辑
	re.POST("/Edit", Controller.EditMessage_cx)
	re.POST("/Department/Basic", Controller.DepartmentBasic)
	re.POST("/SearchEmpSalary", Controller.SearchEmpSalary)
	re.POST("/Department/Init", Controller.DepartmentInit)
	re.POST("/EmpSalary/Init", Controller.EmpSalaryInit)
	re.POST("/Department/Update", Controller.DepartmentUpdate)
	re.POST("/SearchByDepartmentNumber", Controller.SearchByDepartmentNumber)
}
