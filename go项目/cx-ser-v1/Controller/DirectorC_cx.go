package Controller

import (
	"cx/Global"
	"cx/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//查找主管部门的员工
func EmpByDirector_cx(c_cx *gin.Context) {
	req_cx := make(map[string]interface{})
	c_cx.ShouldBind(&req_cx)
	reqId_cx := req_cx["id"].(string)
	fmt.Println("登录的主管ID：", reqId_cx)
	//根据主管id查找部门，返回此部门的所有员工
	var emp_cx []Model.Employee_cx
	var err_cx error
	var director_cx Model.Employee_cx
	err_cx = Global.Db.Where("id = ?", reqId_cx).First(&director_cx).Error
	if err_cx != nil {
		fmt.Println("查找主管ID出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找主管ID出错",
		})
		return
	}
	err_cx = Global.Db.Where("department_name = ? and employee_type != ?", director_cx.Department_name, "").Find(&emp_cx).Error
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




