package Mapper

import (
	"github.com/gin-gonic/gin"
	"cx/Controller"
)

func AdminMapper(re *gin.Engine) {
	//POST请求路由              路径         hander
	re.POST("/AdminLogin", Controller.AdminLogin)
	//菜单 不知道干嘛的  先不删
	//re.POST("/AllMenu", Controller.AllMenu)
	re.POST("/EmployeeBasic", Controller.EmployeeBasic)
	//re.POST("/SalaryBasic_cx", Controller.SalaryBasicByEmp)

	re.POST("/EmployeeBasic/Update", Controller.EmployeeBasicUpdate)
	//re.POST("/EmployeeBasic/Delete", Controller.EmployeeBasicDelete)
	//re.POST("/EmployeeBasic/MultiDelete", Controller.EmployeeBasicMultiDelete)
	re.POST("/EmployeeBasic/Add", Controller.EmployeeBasicAdd)
	re.POST("/EmployeeBasic/Search", Controller.EmployeeBasicSearch)

	//re.POST("/EmployeeBasic/SearchDate", Controller.EmployeeBasicSearchDate)
	//re.POST("/EmployeeBasic/SearchAdvance", Controller.EmployeeBasicSearchAdvance)
	//re.POST("/EmployeeBasic/Import", Controller.EmployeeBasicImport)
	re.POST("/EmployeeBasic/Export", Controller.EmployeeBasicExport)
	//re.POST("/DepartmentTransfer", Controller.DepartmentTransfer)
	//re.POST("/DepartmentTransfer/Search", Controller.DepartmentTransferSearch)
	//re.POST("/PostTransfer", Controller.PostTransfer)
	//re.POST("/PostTransfer/Search", Controller.PostTransferSearch)
	//re.POST("/EmployeeQuit", Controller.EmployeeQuit)
	//re.POST("/EmployeeQuit/SearchByDate", Controller.EmployeeQuitSearchByDate)
	//re.POST("/ExpelEmployee", Controller.ExpelEmployee)
	//re.POST("/AdminInitExamine", Controller.AdminInitExamine)
	//re.POST("/TalentPool/Add", Controller.TalentPoolAdd)
	re.POST("/Edit", Controller.Edit)
	//re.POST("/TalentApply", Controller.TalentApply)
	//re.POST("/TalentAgreeApply", Controller.TalentAgreeApply)
	//re.POST("/TalentRefuseApply", Controller.TalentRefuseApply)
	re.POST("/Department/Basic", Controller.DepartmentBasic)
	re.POST("/SearchEmpSalary", Controller.SearchEmpSalary)

	re.POST("/Department/Init", Controller.DepartmentInit)
	re.POST("/EmpSalary/Init", Controller.EmpSalaryInit)

	//re.POST("/Department/Post", Controller.DepartmentPost)
	re.POST("/Department/Update", Controller.DepartmentUpdate)
	//re.POST("/Department/Delete", Controller.DepartmentDelete)
	//re.POST("/Department/Add", Controller.DepartmentAdd)
	//re.POST("/Post/Init", Controller.PostInit)
	//re.POST("/Post/Add", Controller.PostAdd)
	//re.POST("/Post/Update", Controller.PostUpdate)
	//re.POST("/Post/Delete", Controller.PostDelete)
	re.POST("/SearchByDepartmentNumber", Controller.SearchByDepartmentNumber)
	//re.POST("/SearchByPostNumber", Controller.SearchByPostNumber)
	//re.POST("/DepartmentTransfer/SearchByDate", Controller.DepartmentTransferSearchByDate)
	//re.POST("/PostTransfer/SearchByDate", Controller.PostTransferSearchByDate)
}
