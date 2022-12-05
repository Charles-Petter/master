package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"cx/Global"

	"cx/Model"
	"net/http"
)

func EmployeeBasic_cx(c_cx *gin.Context)  {
	request_cx := make(map[string]interface{})
	c_cx.ShouldBind(&request_cx)
	//请求员工ID
	requestId_cx := request_cx["id"].(string)
	fmt.Println("登录的员工ID：", requestId_cx)
    //员工表数组
	var employees_cx []Model.Employee_cx
	var err_cx error
	var employee Model.Employee_cx
	//根据员工id查找部门，返回此部门的所有同事
	err_cx = Global.Db.Where("id = ?", requestId_cx).First(&employee).Error
	if err_cx != nil {
		fmt.Println("查找员工ID出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "查找员工ID出错",
		})
		return
	}
	err_cx = Global.Db.Where("department_name = ? and employee_type != ?", employee.Department_name, "").Find(&employees_cx).Error
	if err_cx != nil {
		fmt.Println("查找员工部门的同事出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "查找员工部门的同事出错",
		})
		return
	}
	c_cx.JSON(http.StatusOK, employees_cx)
}
//查询员工
func EmployeeSearchEmp_cx(c_cx *gin.Context)  {
	request_cx := make(map[string]interface{})
	c_cx.ShouldBind(&request_cx)
	requestId_cx := request_cx["id"].(string)//登录员工的id
	requestName_cx := request_cx["name"].(string)//要查找的姓名
	fmt.Println("id = ", requestId_cx, ";name = ", requestName_cx)
	var tempDepartment Model.Employee_cx
	var err_cx error
	err_cx = Global.Db.Where("id = ?", requestId_cx).First(&tempDepartment).Error
	if err_cx != nil {
		fmt.Println("查找员工信息出错", err_cx)
	}
	fmt.Println("员工信息：", tempDepartment)
	var employee []Model.Employee_cx
	err_cx = Global.Db.Where("name = ? and department_name = ? and employee_type != ?", requestName_cx, tempDepartment.Department_name).First(&employee).Error
	if err_cx != nil {
		fmt.Println("未找到此员工")
		c_cx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "未找到此员工",
		})
		return
	}
	fmt.Println("查询结果：", employee)
	c_cx.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "查询成功",
		"data" : employee,
	})
}
