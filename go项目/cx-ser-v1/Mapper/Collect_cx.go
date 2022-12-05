package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)
//主管路由
func DirectorMapper_cx(re *gin.Engine)  {
	//查找部门员工
	re.POST("/EmployeeBasicByDirector", Controller.EmpByDirector_cx)
	//部门搜索
	re.POST("/DepartmentSearch_cx", Controller.DepartmentSearch_cx)
}
