package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)

func DirectorMapper(re *gin.Engine)  {
	re.POST("/DirectorDepartment", Controller.DirectorDepartment)
	re.POST("/DirectorInitExamine", Controller.DirectorInitExamine)
	re.POST("/DirectorAgreeApply", Controller.DirectorAgreeApply)
	re.POST("/DirectorRefuseApply", Controller.DirectorRefuseApply)
	re.POST("/EmployeeBasicByDirector", Controller.EmployeeBasicByDirector)
	re.POST("/EmployeeBasic/SearchByDirector", Controller.EmployeeBasicSearchByDirector)
	//re.POST("/EmployeeBasic/SearchDateByDirector", Controller.EmployeeBasicSearchDateByDirector)
	re.POST("/EmployeeQuitByDirector", Controller.EmployeeQuitByDirector)
	re.POST("/EmployeeQuitByDirector/SearchByDate", Controller.EmployeeBasicByDirectorBySearchByDate)
	re.POST("/GetDepartmentName", Controller.GetDepartmentName)
	re.POST("/TalentApplyPart", Controller.TalentApplyPart)

}
