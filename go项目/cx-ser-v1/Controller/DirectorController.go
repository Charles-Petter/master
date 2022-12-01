package Controller

import (
	"cx/Global"
	"cx/Model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//主管部门
func DirectorDepartment(context *gin.Context) {
	var requestDirector Model.Employee
	json.NewDecoder(context.Request.Body).Decode(&requestDirector)
	var director Model.Employee
	//db.Where(query interface{}, args ...interface{})
	//这里问号(?), 在执行的时候会被requestDirector.Id替代  First查询一条记录
	err := Global.Db.Where("id = ?", requestDirector.Id).First(&director).Error
	if err != nil {
		fmt.Println("查询主管的部门出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询主管的部门出错",
		})
		return
	}
	fmt.Println("主管：", director)
	context.JSON(http.StatusOK, director)
}

//主管信息
//func DirectorInitExamine(context *gin.Context) {
//	var requestDirector Model.Employee
//	json.NewDecoder(context.Request.Body).Decode(&requestDirector)
//	fmt.Println("请求到的主管：", requestDirector)
//	var application []Model.PostApplications
//	var err error
//	err = Global.Db.Where("new_department_name = ? and is_agree = ?", requestDirector.Department_name, 0).Find(&application).Error
//	if err != nil {
//		fmt.Println("寻找初始审核信息出错", err)
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "初始化审核信息出错",
//		})
//		return
//	}
//	fmt.Println("即将返回的申请数据：", application)
//	context.JSON(http.StatusOK, application)
//}


func EmployeeBasicByDirector(context *gin.Context) {
	request := make(map[string]interface{})
	context.ShouldBind(&request)
	requestId := request["id"].(string)
	fmt.Println("登录的主管ID：", requestId)
	//根据主管id查找部门，返回此部门的所有员工
	var employee []Model.Employee
	var err error
	var director Model.Employee
	err = Global.Db.Where("id = ?", requestId).First(&director).Error
	if err != nil {
		fmt.Println("查找主管ID出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找主管ID出错",
		})
		return
	}
	err = Global.Db.Where("department_name = ? and employee_type != ?", director.Department_name, "总裁").Find(&employee).Error
	if err != nil {
		fmt.Println("查找主管部门的员工出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找主管部门的员工出错",
		})
		return
	}
	context.JSON(http.StatusOK, employee)
}



func GetDepartmentName(context *gin.Context) {
	request := make(map[string]interface{})
	context.ShouldBind(&request)
	requestId := request["id"].(string)
	fmt.Println("登录的主管ID：", requestId)

	var err error
	var employee Model.Employee
	err = Global.Db.Where("id = ?", requestId).First(&employee).Error
	if err != nil {
		fmt.Println("查找主管信息出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找主管信息出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": employee.Department_name,
	})
}
//func TalentApplyPart(ctx *gin.Context) {
//	var requestDirector Model.Employee
//	json.NewDecoder(ctx.Request.Body).Decode(&requestDirector)
//	fmt.Println("请求到的总管：", requestDirector)
//
//	var application []Model.TalentPool
//	var err error
//	err = Global.Db.Where("is_agree = ? and department_name = ?", 0, requestDirector.Department_name).Find(&application).Error
//	if err != nil {
//		fmt.Println("寻找初始审核信息出错", err)
//		ctx.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "初始化审核信息出错",
//		})
//		return
//	}
//	fmt.Println("即将返回的申请数据：", application)
//	ctx.JSON(http.StatusOK, application)
//}
//主管组件查找
func EmployeeBasicSearchByDirector(context *gin.Context) {
	request := make(map[string]interface{})
	context.ShouldBind(&request)
	requestId := request["id"].(string)     //部门主管的id
	requestName := request["name"].(string) //要查找的姓名
	fmt.Println("id = ", requestId, ";name = ", requestName)
	var tempDepartment Model.Employee
	var err error
	err = Global.Db.Where("id = ?", requestId).First(&tempDepartment).Error
	if err != nil {
		fmt.Println("查找主管信息出错", err)
	}
	fmt.Println("主管信息：", tempDepartment)
	var employee []Model.Employee
	err = Global.Db.Where("name = ? and department_name = ? and employee_type != ?", requestName, tempDepartment.Department_name, "总裁").First(&employee).Error
	if err != nil {
		fmt.Println("未找到此员工")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "未找到此员工",
		})
		return
	}
	fmt.Println("查询结果：", employee)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": employee,
	})
}


