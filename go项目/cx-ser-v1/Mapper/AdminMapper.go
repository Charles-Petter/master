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
	//re.POST("/EmployeeBasic", Controller.EmployeeAll_cx)
	//修改
	re.POST("/EmployeeBasic/Update_cx", Controller.EmployeeBasicUpdate_cx)
	//添加
	re.POST("/EmployeeBasic/Add_cx", Controller.EmployeeBasicAdd_cx)
	//编辑
	//re.POST("/Edit", Controller.EditMessage_cx)
	re.POST("/Department/Message_cx", Controller.DepartmentMessage_cx)
	re.POST("/SearchEmpSalary_cx", Controller.SearchEmpSalary_cx)
	re.POST("/Department/Init_cx", Controller.DepartmentInit_cx)
	re.POST("/EmpSalary/Init_cx", Controller.EmpSalaryInit_cx)
    //员工工资
	re.POST("/Employee/Salary_cx", Controller.EmployeeSalary_cx)
	//re.POST("/Department/Update", Controller.DepartmentUpdate)
	//re.POST("/SearchByDepartmentNumber", Controller.SearchByDepartmentNumber_cx)
}
