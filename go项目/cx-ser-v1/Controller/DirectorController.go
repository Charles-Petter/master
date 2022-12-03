package Controller

import (
	"cx/Global"
	"cx/Model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//查询主管的部门
func DirectorDepartment(c_cx *gin.Context) {
	var reqDirector_cx Model.Employee_cx
	json.NewDecoder(c_cx.Request.Body).Decode(&reqDirector_cx)
	var director_cx Model.Employee_cx
	//根据ID或者请求到的Id查询
	err_cx := Global.Db.Where("id = ?", reqDirector_cx.Id).First(&director_cx).Error
	if err_cx != nil {
		fmt.Println("查询主管的部门出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询主管的部门出错",
		})
		return
	}
	fmt.Println("主管：", director_cx)
	//返回数据
	c_cx.JSON(http.StatusOK, director_cx)
}


//查找主管部门的员工
func EmployeeByDirector_cx(c_cx *gin.Context) {
	req_cx := make(map[string]interface{})
	c_cx.ShouldBind(&req_cx)
	requestId := req_cx["id"].(string)
	fmt.Println("登录的主管ID：", requestId)
	//根据主管id查找部门，返回此部门的所有员工
	var emp_cx []Model.Employee_cx
	var err_cx error
	var director Model.Employee_cx
	err_cx = Global.Db.Where("id = ?", requestId).First(&director).Error
	if err_cx != nil {
		fmt.Println("查找主管ID出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找主管ID出错",
		})
		return
	}
	err_cx = Global.Db.Where("department_name = ? and employee_type != ?", director.Department_name, "总裁").Find(&emp_cx).Error
	if err_cx != nil {
		fmt.Println("查找主管部门的员工出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找主管部门的员工出错",
		})
		return
	}
	c_cx.JSON(http.StatusOK, emp_cx)
}


//获取登录主管的信息
func GetDepartmentName_cx(c_cx *gin.Context) {
	request := make(map[string]interface{})
	c_cx.ShouldBind(&request)
	requestId := request["id"].(string)
	fmt.Println("登录的主管ID：", requestId)
	var err_cx error
	var emp_cx Model.Employee_cx
	err_cx = Global.Db.Where("id = ?", requestId).First(&emp_cx).Error
	if err_cx != nil {
		fmt.Println("查找主管信息出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找主管信息出错",
		})
		return
	}
	c_cx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": emp_cx.Department_name,
	})
}
//主管组件查找
//func EmployeeSearchByDirector_cx(c_cx *gin.Context) {
//	req_cx := make(map[string]interface{})
//	c_cx.ShouldBind(&req_cx)
//	requestId := req_cx["id"].(string)     //部门主管的id
//	requestName := req_cx["name"].(string) //要查找的姓名
//	fmt.Println("id = ", requestId, ";name = ", requestName)
//	var tempDepartment Model.Employee_cx
//	var err_cx error
//	err_cx = Global.Db.Where("id = ?", requestId).First(&tempDepartment).Error
//	if err_cx != nil {
//		fmt.Println("查找主管信息出错", err_cx)
//	}
//	fmt.Println("主管信息：", tempDepartment)
//	//var emp_cx []Model.Employee_cx
//	//err_cx = Global.Db.Where("name = ? and department_name = ? and employee_type != ?", requestName, tempDepartment.Department_name, "总裁").First(&emp_cx).Error
//	//if err_cx != nil {
//	//	fmt.Println("未找到此员工")
//	//	c_cx.JSON(http.StatusOK, gin.H{
//	//		"code": 200,
//	//		"msg":  "未找到此员工",
//	//	})
//	//	return
//	//}
//	//fmt.Println("查询结果：", emp_cx)
//	//c_cx.JSON(http.StatusOK, gin.H{
//	//	"code": 200,
//	//	"msg":  "查询成功",
//	//	"data": emp_cx,
//	//})
//}


