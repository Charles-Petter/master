package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)
//主管路由
func DirectorMapper(re *gin.Engine)  {
	//re.POST("/DirectorDepartment", Controller.DirectorDepartment_cx)
	re.POST("/EmployeeBasicByDirector", Controller.EmployeeByDirector_cx)
	//查询主管
	//re.POST("/EmployeeBasicSearchByDirector/", Controller.EmployeeSearchByDirector_cx)
	//获取部门名称
	re.POST("/GetDepartmentName", Controller.GetDepartmentName_cx)
	//部门搜索
	re.POST("/DepartmentSearch_cx", Controller.DepartmentSearch_cx)
}
