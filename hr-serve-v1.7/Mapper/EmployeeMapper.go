package Mapper

import (
	"github.com/gin-gonic/gin"
	"m/Controller"
)

func EmployeeMapper(re *gin.Engine)  {
	re.POST("/ApplyPostInformation", Controller.ApplyPostInformation)
	re.POST("/ApplyNewPost", Controller.ApplyNewPost)
	re.POST("/ApplyResignedInformation", Controller.ApplyResignedInformation)
	re.POST("/ApplyResigned", Controller.ApplyResigned)
	re.POST("/EmployeeBasicByEmployee", Controller.EmployeeBasicByEmployee)
	re.POST("/EmployeeBasic/SearchByEmployee", Controller.EmployeeBasicSearchByEmployee)
	re.POST("/EmployeeBasic/SearchDateByEmployee", Controller.EmployeeBasicSearchDateByEmployee)
}
