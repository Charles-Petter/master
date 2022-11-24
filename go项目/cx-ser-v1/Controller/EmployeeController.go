package Controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgtype"

	"cx/Global"

	"cx/Model"
	"net/http"

	"strconv"
	"time"
)
//申请调动员工的信息
func ApplyPostInformation(context *gin.Context)  {
	var requestEmployee Model.Employee
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
	fmt.Println("申请调动的员工信息：", requestEmployee)
	var employee []Model.Employee
	err := Global.Db.Where("id = ?", requestEmployee.Id).First(&employee).Error
	if err != nil {
		fmt.Println("错误：", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "查询员工出错",
		})
		return
	}
	fmt.Println("查询到的员工：", employee)
	context.JSON(http.StatusOK, employee)
}
//请求新的调动信息
func ApplyNewPost(context *gin.Context)  {
	var requestApplication Model.PostApplications
	json.NewDecoder(context.Request.Body).Decode(&requestApplication)
	fmt.Println("请求的调动信息：", requestApplication)

	var err error
	var tempEmployee Model.Employee
	err = Global.Db.Where("id = ?", requestApplication.Id).First(&tempEmployee).Error
	if err != nil {
		fmt.Println("出错了", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "查找员工出错",
		})
		return
	}
	requestApplication.Name = tempEmployee.Name
	requestApplication.Original_department_name = tempEmployee.Department_name
	requestApplication.Original_post_name = tempEmployee.Post_name

	var tempCount int64
	err = Global.Db.Table("post_applications").Count(&tempCount).Error
	if err != nil {
		fmt.Println("计算记录数量出错！", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "计算记录数量出错",
		})
		return
	}
	fmt.Println("记录数量：", tempCount)
	strInt64 := strconv.FormatInt(tempCount, 10)
	count, _ := strconv.Atoi(strInt64)
	requestApplication.Application_id = count

	fmt.Println("即将写入的申请信息：", requestApplication)

	var tempApplication Model.PostApplications
	err = Global.Db.Where("id = ? and original_department_name = ? and new_department_name = ? and original_post_name = ? and new_post_name = ? and is_agree = ?", requestApplication.Id, requestApplication.Original_department_name, requestApplication.New_department_name, requestApplication.Original_post_name, requestApplication.New_post_name, 0).First(&tempApplication).Error
	if err == nil {
		fmt.Println("请勿重复提交！", tempApplication)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "请勿重复提交",
		})
		return
	}

	err = Global.Db.Create(&requestApplication).Error
	if err != nil {
		fmt.Println("写入申请信息出错！", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "写入申请信息出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "申请成功",
	})

}
//申请离职的员工信息
func ApplyResignedInformation(context *gin.Context)  {
	var requestEmployee Model.Employee
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
	fmt.Println("申请离职的员工信息：", requestEmployee)
	var employee []Model.Employee
	err := Global.Db.Where("id = ?", requestEmployee.Id).First(&employee).Error
	if err != nil {
		fmt.Println("错误：", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "查询员工出错",
		})
		return
	}
	fmt.Println("查询到的员工：", employee)
	context.JSON(http.StatusOK, employee)
}
//申请离职
func ApplyResigned(context *gin.Context)  {
	var requestResigned Model.ResignedEmployee
	json.NewDecoder(context.Request.Body).Decode(&requestResigned)
	fmt.Println("收到的离职申请：", requestResigned)

	var err error
	err = Global.Db.Table("employees").Where("id = ?", requestResigned.Id).Update("is_quit", "是").Error
	if err != nil {
		fmt.Println("更新员工信息表出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "更新员工信息表出错",
		})
		return
	}
	requestResigned.Resignation_date = pgtype.Date{time.Now(), 2, 0}
	err = Global.Db.Create(&requestResigned).Error
	if err != nil {
		fmt.Println("写入离职信息表出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "写入离职信息表出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "离职成功",
	})
}

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

func EmployeeBasicSearchDateByEmployee(context *gin.Context)  {
	requestDate := make(map[string]interface{})
	context.ShouldBind(&requestDate)
	fmt.Println(requestDate)
	requestId := requestDate["id"].(string)
	startDate := requestDate["start"].(string)
	endDate := requestDate["end"].(string)
	fmt.Println("开始日期：", startDate)
	fmt.Println("结束日期：", endDate)
	tempStart, _ := time.Parse("2006-01-02", startDate)
	tempEnd, _ := time.Parse("2006-01-02", endDate)
	fmt.Println(tempStart, tempEnd)
	start := pgtype.Date{tempStart, 2, 0}
	end := pgtype.Date{tempEnd, 2, 0}
	fmt.Println("开始日期2：", start)
	fmt.Println("结束日期2：", end)
	var err error
	var temp Model.Employee
	err = Global.Db.Where("id = ?", requestId).First(&temp).Error
	var employee []Model.Employee
	err = Global.Db.Where("entry_date >= ? and entry_date <= ? and department_name = ? and employee_type != ?", start, end, temp.Department_name, "总裁").Find(&employee).Error
	if err != nil {
		fmt.Println("查找指定日期区间的员工出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "查找指定日期区间的员工出错",
		})
		return
	}
	fmt.Println("查找结果：", employee)
	context.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "查找成功",
		"data" : employee,
	})
}