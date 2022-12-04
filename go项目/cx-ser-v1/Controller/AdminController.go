package Controller

import (
	"cx/Global"
	"cx/Model"
	"cx/Response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
//登录信息
func Info_cx(c_cx *gin.Context) {
	admin, _ := c_cx.Get("user")
	c_cx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": admin,
		},
	})
}
// 身份验证
func VerifyLogin_cx(c_cx *gin.Context) {
	fmt.Println("进入身份验证")
	var requestAdmin = Model.Employee_cx{}
	json.NewDecoder(c_cx.Request.Body).Decode(&requestAdmin)
	fmt.Println("获取的数据:", requestAdmin)
	id := requestAdmin.Id
	//password := requestAdmin.Password
	//数据验证
	fmt.Println("开始数据验证")

	var admin Model.Employee_cx
	var err error

	err = Global.Db.Where("id = ?", id).First(&admin).Error
	if err != nil {
		fmt.Println("出错了!", err)
	}
	fmt.Println("数据库查到的信息：", admin)
	//验证员工类型和账户密码
	if admin.Password != requestAdmin.Password || admin.Employee_type != requestAdmin.Employee_type {
		fmt.Println("执行了if")
		c_cx.JSON(200, gin.H{
			"code": 200,
			"msg":  "账号或密码或身份错误",
		})
		return
	}
	//发放TOKEN
	token, err := Global.ReaeaseTokec(admin)
	if err != nil {
		c_cx.JSON(200, gin.H{
			"code": 200,
			"msg":  "token发放失败",
		})
		return
	}
	Global.Id = requestAdmin.Id
	Response.Success(c_cx, gin.H{"token": token, "data": admin}, "登录成功")
}
//查找所有员工
//func EmployeeAll_cx(c_cx *gin.Context) {
//	var employee []Model.Employee_cx
//	//根据id查找
//	err := Global.Db.Order("id").Find(&employee).Error
//	if err != nil {
//		fmt.Println("查找所有员工出错！", err)
//		return
//	}
//	fmt.Println("数据库查到的所有员工：", employee)
//	c_cx.JSON(http.StatusOK, employee)
//}
//员工工资计算
func EmployeeSalary_cx(c_cx *gin.Context)  {
	var requestEmployee = Model.Salaytable_cx{}
	json.NewDecoder(c_cx.Request.Body).Decode(&requestEmployee)
	fmt.Println("即将修改的工资条：", requestEmployee)
	err := Global.Db.Where("id = ?", requestEmployee.Id).Updates(&requestEmployee).Error
	if err != nil {
		fmt.Println("修改出错！")
		c_cx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "修改出错",
		})
		return
	}
	c_cx.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "修改成功",
	})
}
//员工数据修改
func EmployeeBasicUpdate_cx(c_cx *gin.Context)  {
	var requestEmployee = Model.Employee_cx{}
	json.NewDecoder(c_cx.Request.Body).Decode(&requestEmployee)
	fmt.Println("即将修改的用户：", requestEmployee)
	err := Global.Db.Where("id = ?", requestEmployee.Id).Updates(&requestEmployee).Error
	if err != nil {
		fmt.Println("修改出错！")
		c_cx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "修改出错",
		})
		return
	}
	c_cx.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "修改成功",
	})
}
//添加员工信息
func EmployeeBasicAdd_cx(c_cx *gin.Context) {
	var requestEmployee = Model.Employee_cx{}
	var err error
	//var Height, Department_number, Post_number int
	json.NewDecoder(c_cx.Request.Body).Decode(&requestEmployee)
	fmt.Println("即将添加的员工信息：", requestEmployee)
	/*查找重复的员工ID*/
	err = Global.Db.Where("id = ?", requestEmployee.Id).First(&requestEmployee).Error
	if err == nil {
		fmt.Println("找到重复的id")
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "编号已存在",
		})
		return
	}
	/*查找部门编号是否存在*/
	err = Global.Db.Where("department_number = ?", requestEmployee.Department_number).First(&Model.Collects_cx{}).Error
	if err != nil {
		fmt.Println("找不到此部门编号")
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "部门编号不存在",
		})
		return
	}
	/*检查部门编号和部门名称是否对应*/
	var tempDepartment Model.Collects_cx
	Global.Db.Where("department_name = ?", requestEmployee.Department_name).First(&tempDepartment) //根据请求来的部门名称检查是否与部门编号所对应
	if requestEmployee.Department_number != tempDepartment.Department_number {
		fmt.Println("部门编号与名称不对应")
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "部门编号与名称不对应",
		})
		return
	}
	//将值插入数据库
	Global.Db.Create(&requestEmployee)
	//JSON格式
	c_cx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

//搜索部门（搜索框）
func DepartmentSearch_cx(c_cx *gin.Context) {
	var requestEmployeeMent Model.Collects_cx
	//http中获取到json数据解码
	json.NewDecoder(c_cx.Request.Body).Decode(&requestEmployeeMent)
	fmt.Println("获取的搜索部门：", requestEmployeeMent.Department_name)
		var tempEmployee []Model.Collects_cx
	//根据部门名称搜索
	err := Global.Db.Where("Department_name = ?", requestEmployeeMent.Department_name).First(&tempEmployee).Error
	if err != nil {
		fmt.Println("未找到此部门")
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "未找到此部门",
		})
		return
	}
	fmt.Println("查询结果：", tempEmployee)
	c_cx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": tempEmployee,
	})
}
//编辑个人信息
//func EditMessage_cx(ctx *gin.Context) {
//	var request Model.Employee_cx
//	//获取参数
//	ctx.ShouldBind(&request)
//	fmt.Println(request)
//	var err error
//	err = Global.Db.Where("id = ?", request.Id).Updates(&request).Error
//	if err != nil {
//		fmt.Println("修改个人信息出错", err)
//		ctx.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "修改个人信息出错",
//		})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "修改成功",
//	})
//}
//初始化生成部门id
//func initEmployeeId(count int) string {
//	var result string
//	result = "2021"
//	//将count转为字符串
//	temp := strconv.Itoa(count)
//	length := len(temp)
//	//计算count字符串长度，生成前缀0
//	zero := "0"
//	for i := 0; i < 4-length; i++ {
//		result += zero
//		//fmt.Println(i)
//	}
//	result += temp
//	fmt.Println("生成的编号：", result)
//	return result
//}
//查询工资信息
func SearchEmpSalary_cx(c_cx *gin.Context) {
	var empsalary []Model.Salaytable_cx
	var err error
	//返回的是工资表数组
	err = Global.Db.Find(&empsalary).Error
	if err != nil {
		fmt.Println("查询工资出错", err)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询工资表出错",
		})
		return
	}
	fmt.Println("所有工资信息：", empsalary)
	c_cx.JSON(http.StatusOK, empsalary)
}



//查询部门信息
func DepartmentMessage_cx(c_cx *gin.Context) {
	var department []Model.Collects_cx
	var err error
	//返回的是部门表数组
	err = Global.Db.Find(&department).Error
	if err != nil {
		fmt.Println("查询部门出错", err)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询部门出错",
		})
		return
	}
	fmt.Println("所有部门信息：", department)
	c_cx.JSON(http.StatusOK, department)
}
//初始化部门信息
func DepartmentInit(c_cx *gin.Context) {

	var count int64
	var temp []Model.Collects_cx
	var err error
	err = Global.Db.Find(&temp).Count(&count).Error
	if err != nil {
		fmt.Println("初始化部门出错", err)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "初始化部门出错",
		})
		return
	}
	var department []string
	//类型转换 string 转int
	strInt64 := strconv.FormatInt(count, 10)
	tempCount, _ := strconv.Atoi(strInt64)
	//循环数组 遍历出部门数组所有值
	for i := 0; i < tempCount; i++ {
		department = append(department, temp[i].Department_name)
	}
	fmt.Println("找到的部门数组：", department)
	c_cx.JSON(http.StatusOK, department)
}


//修改部门
//func DepartmentUpdate(c_cx *gin.Context) {
//var requestDepartment Model.Collects_cx
//	json.NewDecoder(c_cx.Request.Body).Decode(&requestDepartment)
//	fmt.Println("即将修改的部门：", requestDepartment)
//	var err error
//	err = Global.Db.Where("department_number = ?", requestDepartment.Department_number).Updates(&requestDepartment).Error
//	if err != nil {
//		fmt.Println("修改部门出错")
//		c_cx.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "修改部门出错",
//		})
//		return
//	}
//	c_cx.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "修改成功",
//	})
//}

//初始化员工工资表
func EmpSalaryInit_cx(c_cx *gin.Context) {

	var count int64
	var t_cx []Model.Salaytable_cx
	var err error
	err = Global.Db.Find(&t_cx).Count(&count).Error
	if err != nil {
		fmt.Println("初始化员工工资表出错", err)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "初始化员工工资表出错",
		})
		return
	}
	var Empsalary []string
	strInt64 := strconv.FormatInt(count, 10)
	tempCount, _ := strconv.Atoi(strInt64)
	for i := 0; i < tempCount; i++ {
		//de := temp[i].Department_name
		Empsalary = append(Empsalary, t_cx[i].Name)
		//department[i] = temp[i].Department_name
	}
	fmt.Println("找到的部门数组：", Empsalary)
	c_cx.JSON(http.StatusOK, Empsalary)
}

//查找部门数量
//func SearchByDepartmentNumber_cx(c_cx *gin.Context) {
//	type t_cx struct {
//		Text  string `json:"text"`
//		Value int    `json:"value"`
//	}
//
//	var tempDepartment []Model.Collects_cx
//	var err error
//	var tempCount int64
//	err = Global.Db.Table("collects_cxes").Find(&tempDepartment).Count(&tempCount).Error
//	if err != nil {
//		fmt.Println("查找部门数量出错", err)
//	}
//	strInt64 := strconv.FormatInt(tempCount, 10)
//	count, _ := strconv.Atoi(strInt64)
//	fmt.Println("部门数量：", count)
//	var departmentNumber []t_cx
//	for i := 0; i < count; i++ {
//		var t t_cx
//		t.Text = strconv.Itoa(tempDepartment[i].Department_number) + ":" + tempDepartment[i].Department_name
//		t.Value = tempDepartment[i].Department_number
//		departmentNumber = append(departmentNumber, t)
//	}
//	fmt.Println("数组：", departmentNumber)
//	c_cx.JSON(http.StatusOK, departmentNumber)
//}



