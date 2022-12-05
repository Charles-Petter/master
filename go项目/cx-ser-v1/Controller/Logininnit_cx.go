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
	loginUser_cx, _ := c_cx.Get("user")
	c_cx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": loginUser_cx,
		},
	})
}
// 身份验证
func VerifyLogin_cx(c_cx *gin.Context) {
	fmt.Println("身份验证")
	//实例化
	var requestUser_cx = Model.Employee_cx{}
    //获取数据
	json.NewDecoder(c_cx.Request.Body).Decode(&requestUser_cx)
	fmt.Println("获取到的数据:", requestUser_cx)
	//数据验证
	id := requestUser_cx.Id
	fmt.Println("数据验证")

	var loginUser_cx Model.Employee_cx
	var err_cx error
    //在数据库中 获取id
	err_cx = Global.Db.Where("id = ?", id).First(&loginUser_cx).Error
	if err_cx != nil {
		fmt.Println("出错了!", err_cx)
	}
	fmt.Println("数据库查到的信息：", loginUser_cx)
	//验证员工类型和账户密码
	if loginUser_cx.Password != requestUser_cx.Password || loginUser_cx.Employee_type != requestUser_cx.Employee_type {
		fmt.Println("执行了if")
		c_cx.JSON(200, gin.H{
			"code": 200,
			"msg":  "账号或密码或身份错误",
		})
		return
	}
	//发放TOKEN
	token, err_cx := Global.ReaeaseTokec(loginUser_cx)
	if err_cx != nil {
		c_cx.JSON(200, gin.H{
			"code": 200,
			"msg":  "token发放失败",
		})
		return
	}
	Global.Id = requestUser_cx.Id
	Response.Success(c_cx, gin.H{"token": token, "data": loginUser_cx}, "登录成功")
}
//员工工资上传
func EmployeeSalary_cx(c_cx *gin.Context)  {
	var requestEmployee = Model.Salaytable_cx{}
	json.NewDecoder(c_cx.Request.Body).Decode(&requestEmployee)
	fmt.Println("即将修改的工资条：", requestEmployee)
	err := Global.Db.Where("id = ?", requestEmployee.Id).Updates(&requestEmployee).Error
	if err != nil {
		fmt.Println("上传出错！")
		c_cx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "上传出错",
		})
		return
	}
	c_cx.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "上传成功",
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
	var reqEmp_cx = Model.Employee_cx{}
	var err_cx error
	json.NewDecoder(c_cx.Request.Body).Decode(&reqEmp_cx)
	fmt.Println("即将添加的员工信息：", reqEmp_cx)
	/*查找重复的员工ID*/
	err_cx = Global.Db.Where("id = ?", reqEmp_cx.Id).First(&reqEmp_cx).Error
	if err_cx == nil {
		fmt.Println("找到重复的id")
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "编号已存在",
		})
		return
	}
	//查找部门编号是否存在
	err_cx = Global.Db.Where("department_number = ?", reqEmp_cx.Department_number).First(&Model.Collects_cx{}).Error
	if err_cx != nil {
		fmt.Println("找不到此部门编号")
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "部门编号不存在",
		})
		return
	}
	//检查部门编号和部门名称是否对应
	var tempCollect_cx Model.Collects_cx
	Global.Db.Where("department_name = ?", reqEmp_cx.Department_name).First(&tempCollect_cx) //根据请求来的部门名称检查是否与部门编号所对应
	if reqEmp_cx.Department_number != tempCollect_cx.Department_number {
		fmt.Println("部门编号与名称不对应")
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "部门编号与名称不对应",
		})
		return
	}
	//将值插入数据库
	Global.Db.Create(&reqEmp_cx)
	//JSON格式
	c_cx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

//搜索部门（搜索框）
func DepartmentSearch_cx(c_cx *gin.Context) {
	var requestCollect_cx Model.Employee_cx
	//http中获取到json数据解码
	json.NewDecoder(c_cx.Request.Body).Decode(&requestCollect_cx)
	fmt.Println("获取的搜索部门：", requestCollect_cx.Department_name)
		var tempEmployee []Model.Employee_cx
	//根据部门名称搜索
	err_cx := Global.Db.Where("Department_name = ?", requestCollect_cx.Department_name).Find(&tempEmployee).Error
	if err_cx != nil {
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

//查询工资信息
func SearchEmpSalary_cx(c_cx *gin.Context) {
	var empsalary_cx []Model.Salaytable_cx
	var err_cx error
	//返回的是工资表数组
	err_cx = Global.Db.Find(&empsalary_cx).Error
	if err_cx != nil {
		fmt.Println("查询工资出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询工资表出错",
		})
		return
	}
	fmt.Println("所有工资信息：", empsalary_cx)
	c_cx.JSON(http.StatusOK, empsalary_cx)
}



//查询部门信息
func DepartmentMessage_cx(c_cx *gin.Context) {
	var collect_cx []Model.Collects_cx
	var err_cx error
	//返回的是部门表数组
	err_cx = Global.Db.Find(&collect_cx).Error
	if err_cx != nil {
		fmt.Println("查询部门出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询部门出错",
		})
		return
	}
	fmt.Println("所有部门信息：", collect_cx)
	c_cx.JSON(http.StatusOK, collect_cx)
}
//初始化部门信息
func DepartmentInit_cx(c_cx *gin.Context) {
	var count_cx int64
	//实例化
	var temp_cx []Model.Collects_cx
	var err_cx error
	err_cx = Global.Db.Find(&temp_cx).Count(&count_cx).Error
	if err_cx != nil {
		fmt.Println("初始化部门出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "初始化部门出错",
		})
		return
	}
	var collect_cx []string
	//类型转换 string 转int
	strInt64 := strconv.FormatInt(count_cx, 10)
	tempCount, _ := strconv.Atoi(strInt64)
	//循环数组 遍历出部门数组所有值
	for i := 0; i < tempCount; i++ {
		collect_cx = append(collect_cx, temp_cx[i].Department_name)
	}
	fmt.Println("找到的部门数组：", collect_cx)
	c_cx.JSON(http.StatusOK, collect_cx)
}


//初始化员工工资表
func EmpSalaryInit_cx(c_cx *gin.Context) {
	var count_cx int64
	var t_cx []Model.Salaytable_cx
	var err_cx error
	err_cx = Global.Db.Find(&t_cx).Count(&count_cx).Error
	if err_cx != nil {
		fmt.Println("初始化员工工资表出错", err_cx)
		c_cx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "初始化员工工资表出错",
		})
		return
	}
	var Empsalary_cx []string
	strInt64 := strconv.FormatInt(count_cx, 10)
	tempCount, _ := strconv.Atoi(strInt64)
	//遍历数组输出所有部门
	for i := 0; i < tempCount; i++ {
		Empsalary_cx = append(Empsalary_cx, t_cx[i].Name)
	}
	fmt.Println("找到的部门数组：", Empsalary_cx)
	c_cx.JSON(http.StatusOK, Empsalary_cx)
}



