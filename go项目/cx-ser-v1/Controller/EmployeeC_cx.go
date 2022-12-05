package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"cx/Global"

	"cx/Model"
	"net/http"
)
//员工信息
func EmployeeBasic_cx(c_cx *gin.Context)  {
	//通过 请求键值对
	request_cx := make(map[string]interface{})
	c_cx.ShouldBind(&request_cx)
	//请求员工ID
	requestId_cx := request_cx["id"].(string)
	fmt.Println("登录的员工ID：", requestId_cx)
    //员工表数组
	var employees_cx []Model.Employee_cx
	var err_cx error
	var employee_cx Model.Employee_cx
	//根据员工id查找部门，返回此部门的所有同事
	err_cx = Global.Db.Where("id = ?", requestId_cx).First(&employee_cx).Error
	if err_cx != nil {
		fmt.Println("查找员工ID出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "查找员工ID出错",
		})
		return
	}
	//员工姓名和员工类型类型查找相关部门同事
	err_cx = Global.Db.Where("department_name = ? and employee_type != ?", employee_cx.Department_name, "").Find(&employees_cx).Error
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
//查询员工信息
func EmployeeSearchEmp_cx(c_cx *gin.Context)  {
	request_cx := make(map[string]interface{})
	c_cx.ShouldBind(&request_cx)
	requestId_cx := request_cx["id"].(string)//登录员工的id
	requestName_cx := request_cx["name"].(string)//要查找的姓名
	fmt.Println("id = ", requestId_cx, ";name = ", requestName_cx)
	var tempCollect_cx Model.Employee_cx
	var err_cx error
	err_cx = Global.Db.Where("id = ?", requestId_cx).First(&tempCollect_cx).Error
	if err_cx != nil {
		fmt.Println("查找员工信息出错", err_cx)
	}
	fmt.Println("员工信息：", tempCollect_cx)
	var employee_cx []Model.Employee_cx
	err_cx = Global.Db.Where("name = ? and department_name = ? and employee_type != ?", requestName_cx, tempCollect_cx.Department_name).First(&employee_cx).Error
	if err_cx != nil {
		fmt.Println("未找到此员工")
		c_cx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "未找到此员工",
		})
		return
	}
	fmt.Println("查询结果：", employee_cx)
	c_cx.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "查询成功",
		"data" : employee_cx,
	})
}
