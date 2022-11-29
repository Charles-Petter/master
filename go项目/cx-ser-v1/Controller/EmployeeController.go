package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"cx/Global"

	"cx/Model"
	"net/http"
)
//申请调动员工的信息
//func ApplyPostInformation(context *gin.Context)  {
//	var requestEmployee Model.Employee
//	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
//	fmt.Println("申请调动的员工信息：", requestEmployee)
//	var employee []Model.Employee
//	err := Global.Db.Where("id = ?", requestEmployee.Id).First(&employee).Error
//	if err != nil {
//		fmt.Println("错误：", err)
//		context.JSON(http.StatusOK, gin.H{
//			"code" : 200,
//			"msg" : "查询员工出错",
//		})
//		return
//	}
//	fmt.Println("查询到的员工：", employee)
//	context.JSON(http.StatusOK, employee)
//}

//申请离职的员工信息
//func ApplyResignedInformation(context *gin.Context)  {
//	var requestEmployee Model.Employee
//	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
//	fmt.Println("申请离职的员工信息：", requestEmployee)
//	var employee []Model.Employee
//	err := Global.Db.Where("id = ?", requestEmployee.Id).First(&employee).Error
//	if err != nil {
//		fmt.Println("错误：", err)
//		context.JSON(http.StatusOK, gin.H{
//			"code" : 200,
//			"msg" : "查询员工出错",
//		})
//		return
//	}
//	fmt.Println("查询到的员工：", employee)
//	context.JSON(http.StatusOK, employee)
//}


func EmployeeBasicByEmployee(context *gin.Context)  {
	request := make(map[string]interface{})
	context.ShouldBind(&request)
	requestId := request["id"].(string)
	fmt.Println("登录的员工ID：", requestId)

	//根据员工id查找部门，返回此部门的所有同事
	var employees []Model.Employee
	var err error
	var employee Model.Employee
	err = Global.Db.Where("id = ?", requestId).First(&employee).Error
	if err != nil {
		fmt.Println("查找员工ID出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "查找员工ID出错",
		})
		return
	}
	err = Global.Db.Where("department_name = ? and employee_type != ?", employee.Department_name, "总裁").Find(&employees).Error
	if err != nil {
		fmt.Println("查找员工部门的同事出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "查找员工部门的同事出错",
		})
		return
	}
	context.JSON(http.StatusOK, employees)
}

func EmployeeBasicSearchByEmployee(context *gin.Context)  {
	request := make(map[string]interface{})
	context.ShouldBind(&request)
	requestId := request["id"].(string)//登录员工的id
	requestName := request["name"].(string)//要查找的姓名
	fmt.Println("id = ", requestId, ";name = ", requestName)
	var tempDepartment Model.Employee
	var err error
	err = Global.Db.Where("id = ?", requestId).First(&tempDepartment).Error
	if err != nil {
		fmt.Println("查找员工信息出错", err)
	}
	fmt.Println("员工信息：", tempDepartment)
	var employee []Model.Employee
	err = Global.Db.Where("name = ? and department_name = ? and employee_type != ?", requestName, tempDepartment.Department_name, "总裁").First(&employee).Error
	if err != nil {
		fmt.Println("未找到此员工")
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "未找到此员工",
		})
		return
	}
	fmt.Println("查询结果：", employee)
	context.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "查询成功",
		"data" : employee,
	})
}

//func EmployeeBasicSearchDateByEmployee(context *gin.Context)  {
//	requestDate := make(map[string]interface{})
//	context.ShouldBind(&requestDate)
//	fmt.Println(requestDate)
//	requestId := requestDate["id"].(string)
//	startDate := requestDate["start"].(string)
//	endDate := requestDate["end"].(string)
//	fmt.Println("开始日期：", startDate)
//	fmt.Println("结束日期：", endDate)
//	tempStart, _ := time.Parse("2006-01-02", startDate)
//	tempEnd, _ := time.Parse("2006-01-02", endDate)
//	fmt.Println(tempStart, tempEnd)
//	start := pgtype.Date{tempStart, 2, 0}
//	end := pgtype.Date{tempEnd, 2, 0}
//	fmt.Println("开始日期2：", start)
//	fmt.Println("结束日期2：", end)
//	var err error
//	var temp Model.Employee
//	err = Global.Db.Where("id = ?", requestId).First(&temp).Error
//	var employee []Model.Employee
//	err = Global.Db.Where("entry_date >= ? and entry_date <= ? and department_name = ? and employee_type != ?", start, end, temp.Department_name, "总裁").Find(&employee).Error
//	if err != nil {
//		fmt.Println("查找指定日期区间的员工出错", err)
//		context.JSON(http.StatusOK, gin.H{
//			"code" : 200,
//			"msg" : "查找指定日期区间的员工出错",
//		})
//		return
//	}
//	fmt.Println("查找结果：", employee)
//	context.JSON(http.StatusOK, gin.H{
//		"code" : 200,
//		"msg" : "查找成功",
//		"data" : employee,
//	})
//}