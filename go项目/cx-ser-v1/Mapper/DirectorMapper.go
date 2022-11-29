package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)

func DirectorMapper(re *gin.Engine)  {
	re.POST("/DirectorDepartment", Controller.DirectorDepartment)
	//re.POST("/DirectorInitExamine", Controller.DirectorInitExamine)
	re.POST("/EmployeeBasicByDirector", Controller.EmployeeBasicByDirector)
	re.POST("/EmployeeBasicSearchByDirector/", Controller.EmployeeBasicSearchByDirector)
	//re.POST("/EmployeeBasic/SearchDateByDirector", Controller.EmployeeBasicSearchDateByDirector)
	re.POST("/GetDepartmentName", Controller.GetDepartmentName)

	//部门搜索
	re.POST("/DepartmentSearch", Controller.DepartmentSearch)


}
