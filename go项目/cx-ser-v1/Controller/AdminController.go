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
func Info(context *gin.Context) {
	admin, _ := context.Get("user")
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": admin,
		},
	})
}
//进入AdminLogin  身份验证
func AdminLogin(context *gin.Context) {
	fmt.Println("进入AdminLogin")
	var requestAdmin = Model.Employee{}
	json.NewDecoder(context.Request.Body).Decode(&requestAdmin)
	fmt.Println("获取的数据:", requestAdmin)

	id := requestAdmin.Id
	//password := requestAdmin.Password
	//数据验证
	fmt.Println("开始数据验证")

	var admin Model.Employee
	var err error

	err = Global.Db.Where("id = ?", id).First(&admin).Error
	if err != nil {
		fmt.Println("出错了!", err)
	}
	fmt.Println("数据库查到的信息：", admin)
	if admin.Password != requestAdmin.Password || admin.Employee_type != requestAdmin.Employee_type {
		fmt.Println("执行了if")
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "账号或密码或身份错误",
		})
		return
	}
	if admin.Is_quit == "是" {
		fmt.Println("离职员工无法登录")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "离职员工无法登录",
		})
		return
	}
	token, err := Global.ReaeaseTokec(admin)
	if err != nil {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "token发放失败",
		})
		return
	}
	Global.Id = requestAdmin.Id
	Response.Success(context, gin.H{"token": token, "data": admin}, "登录成功")
}
//查找所有员工
func EmployeeBasic(context *gin.Context) {
	var employee []Model.Employee
	err := Global.Db.Order("id").Find(&employee).Error
	if err != nil {
		fmt.Println("查找所有员工出错！", err)
		return
	}
	fmt.Println("数据库查到的所有员工：", employee)
	context.JSON(http.StatusOK, employee)
}

//func SalaryBasic_cx(context *gin.Context) {
//	var employee []Model.Salaytable
//	err := Global.Db.Order("id").Find(&employee).Error
//	if err != nil {
//		fmt.Println("查找所有员工出错！", err)
//		return
//	}
//	fmt.Println("数据库查到的所有员工：", employee)
//	context.JSON(http.StatusOK, employee)
//}

////在basic组件中删除员工
//func EmployeeBasicDelete(context *gin.Context) {
//	var requestEmployee = Model.Employee{}
//	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
//	fmt.Println("即将删除的员工：", requestEmployee)
//	err := Global.Db.Where("id = ?", requestEmployee.Id).Delete(&requestEmployee).Error
//	if err != nil {
//		fmt.Println("删除出错")
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "删除出错",
//		})
//		return
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "删除成功",
//	})
//}

//func EmployeeBasicMultiDelete(context *gin.Context) {
//	var request []string
//	count := 0
//	context.ShouldBindJSON(&request)
//	fmt.Println(request)
//	for key := range request {
//		fmt.Println(key)
//		var err error
//		var employee Model.Employee
//		err = Global.Db.Where("id = ?", request[key]).First(&employee).Error
//		if err != nil {
//			fmt.Println("查找员工出错", err)
//			context.JSON(http.StatusOK, gin.H{
//				"code": 200,
//				"msg":  "查找员工出错",
//			})
//			return
//		}
//		err = Global.Db.Table("employees").Where("id = ?", request[key]).Delete(&employee).Error
//		if err != nil {
//			fmt.Println("删除员工出错", err)
//			context.JSON(http.StatusOK, gin.H{
//				"code": 200,
//				"msg":  "删除员工出错",
//			})
//			return
//		}
//		count++
//	}
//	fmt.Println(count)
//	context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "删除成功",
//	})
//
//}




func EmployeeBasicUpdate(context *gin.Context)  {
	var requestEmployee = Model.Employee{}
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
	fmt.Println("即将修改的用户：", requestEmployee)
	err := Global.Db.Where("id = ?", requestEmployee.Id).Updates(&requestEmployee).Error
	if err != nil {
		fmt.Println("修改出错！")
		context.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "修改出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "修改成功",
	})
}
//添加员工信息
func EmployeeBasicAdd(context *gin.Context) {
	var requestEmployee = Model.Employee{}
	var err error
	//var Height, Department_number, Post_number int
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
	fmt.Println("即将添加的员工信息：", requestEmployee)

	/*查找重复的员工ID*/
	err = Global.Db.Where("id = ?", requestEmployee.Id).First(&requestEmployee).Error
	if err == nil {
		fmt.Println("找到重复的id")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "编号已存在",
		})
		return
	}
	/*查找部门编号是否存在*/
	err = Global.Db.Where("department_number = ?", requestEmployee.Department_number).First(&Model.Department{}).Error
	if err != nil {
		fmt.Println("找不到此部门编号")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "部门编号不存在",
		})
		return
	}
	/*查找岗位编号是否存在*/
	//err = Global.Db.Where("post_number = ?", requestEmployee.Post_number).First(&Model.PostTable{}).Error
	//if err != nil {
	//	fmt.Println("找不到此岗位编号")
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"msg":  "岗位编号不存在",
	//	})
	//	return
	//}
	/*检查部门编号和部门名称是否对应*/
	var tempDepartment Model.Department
	Global.Db.Where("department_name = ?", requestEmployee.Department_name).First(&tempDepartment) //根据请求来的部门名称检查是否与部门编号所对应
	if requestEmployee.Department_number != tempDepartment.Department_number {
		fmt.Println("部门编号与名称不对应")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "部门编号与名称不对应",
		})
		return
	}
	/*检查岗位编号和名称是否对应*/
	//var tempPost Model.PostTable // 将Model层PostTable表定义为一个临时变量tempPost
	//Global.Db.Where("post_name = ?", requestEmployee.Post_name).First(&tempPost)
	//if requestEmployee.Post_number != tempPost.Post_number {
	//	fmt.Println("岗位编号与名称不对应")
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"msg":  "岗位编号与名称不对应",
	//	})
	//	return
	//}
	//将值插入数据库
	Global.Db.Create(&requestEmployee)
	//JSON格式
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}
//搜索名字（搜索框）
func EmployeeBasicSearch(context *gin.Context) {
	var requestEmployee Model.Employee
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
	fmt.Println("获取的搜索名字：", requestEmployee.Name)
	var tempEmployee []Model.Employee
	err := Global.Db.Where("name = ?", requestEmployee.Name).First(&tempEmployee).Error
	if err != nil {
		fmt.Println("未找到此员工")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "未找到此员工",
		})
		return
	}
	fmt.Println("查询结果：", tempEmployee)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": tempEmployee,
	})
}

//搜索部门（搜索框）
func DepartmentSearch(context *gin.Context) {
	var requestEmployeeMent Model.Employee
	json.NewDecoder(context.Request.Body).Decode(&requestEmployeeMent)
	fmt.Println("获取的搜索部门：", requestEmployeeMent.Department_name)
		var tempEmployee []Model.Employee
	err := Global.Db.Where("Department_name = ?", requestEmployeeMent.Department_name).First(&tempEmployee).Error
	if err != nil {
		fmt.Println("未找到此部门")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "未找到此部门",
		})
		return
	}
	fmt.Println("查询结果：", tempEmployee)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": tempEmployee,
	})
}









//查询指定日期间的数据
//func EmployeeBasicSearchDate(context *gin.Context) {
//	requestDate := make(map[string]interface{})
//	context.ShouldBind(&requestDate)
//	fmt.Println(requestDate)
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
//	var employee []Model.Employee
//	err = Global.Db.Where("entry_date >= ? and entry_date <= ?", start, end).Find(&employee).Error
//	if err != nil {
//		fmt.Println("查找指定日期区间的员工出错", err)
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "查找指定日期区间的员工出错",
//		})
//		return
//	}
//	fmt.Println("查找结果：", employee)
//	context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "查找成功",
//		"data": employee,
//	})
//}
//高级搜索
//func EmployeeBasicSearchAdvance(context *gin.Context) {
//	requestSearch := make(map[string]interface{})
//	context.ShouldBind(&requestSearch)
//	//requestPolitical := requestSearch["political"].(string)
//	//requestNation := requestSearch["nation"].(string)
//	//根据部门查询
//	requestDepartment := requestSearch["department_name"].(string)
//	//requestPost := requestSearch["post_name"].(string)
//	//requestSchool := requestSearch["graduation_school"].(string)
//	//requestMajor := requestSearch["major_studied"].(string)
//	//fmt.Println("高级搜索：", requestPolitical, requestNation, requestDepartment, requestPost, requestSchool, requestMajor)
//	fmt.Println("高级搜索：",requestDepartment)
//
//	var tempEmployee []Model.Employee
//	err := Global.Db.Where("political = ? and nation = ? and department_name = ? and post_name = ? and graduation_school = ? and major_studied = ?", requestDepartment ).First(&tempEmployee).Error
//	if err != nil {
//		fmt.Println("未找到此员工")
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "未找到此员工",
//		})
//		return
//	}
//	fmt.Println("查询结果：", tempEmployee)
//	context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "查询成功",
//		"data": tempEmployee,
//	})
//}
//文件上传
//func EmployeeBasicImport(context *gin.Context) {
//	fmt.Println("进入文件上传")
//	files, _ := context.FormFile("file")
//	dst := path.Join("../", files.Filename)
//	a := context.SaveUploadedFile(files, dst)
//	fmt.Println("进入文件上传2")
//	if a != nil {
//		fmt.Println("a = ", a)
//		context.JSON(http.StatusOK, gin.H{
//			"code": 203,
//			"msg":  a,
//		})
//		return
//	}
//	fmt.Println("进入文件上传3")
//	xlsx, err := excelize.OpenFile(dst)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	rows := xlsx.GetRows("Sheet" + "1")
//	fmt.Println("进入文件上传4")
//	for key, row := range rows {
//		if key > 0 {
//			fmt.Println("第五列", row[5])
//			Birthday, _ := time.Parse("2006-01-02", row[5])
//			fmt.Println("生日：", Birthday)
//			Entry_date, _ := time.Parse("2006-01-02", row[21])
//			Graduation_date, _ := time.Parse("2006-01-02", row[27])
//			Height, _ := strconv.Atoi(row[12])
//			Department_number, _ := strconv.Atoi(row[17])
//			Post_number, _ := strconv.Atoi(row[19])
//			employee := Model.Employee{
//				Id: row[0], Password: row[1], Employee_type: row[2],
//				Name: row[3], Sex: row[4], Birthday: pgtype.Date{Birthday, 2, 0},
//				Id_card: row[6], Political: row[7], Nation: row[8],
//				Native_place: row[9], Phone: row[10], Email: row[11],
//				Height: Height, Blood_type: row[13], Marital_status: row[14],
//				Birthplace: row[15], Registered_residence: row[16], Department_number: Department_number,
//				Department_name: row[18], Post_number: Post_number, Post_name: row[20],
//				Entry_date: pgtype.Date{Entry_date, 2, 0}, Employment_form: row[22], Personnel_source: row[23],
//				Highest_education: row[24], Graduation_school: row[25], Major_studied: row[26],
//				Graduation_date: pgtype.Date{Graduation_date, 2, 0}, Is_quit: row[28],
//			}
//			fmt.Println("employee = ", employee)
//			err := Global.Db.Create(&employee).Error
//			if err != nil {
//				fmt.Println("发现错误！", err)
//				context.JSON(http.StatusOK, gin.H{
//					"code": 202,
//					"msg":  "员工:" + employee.Name + "已经存在",
//				})
//				return
//			}
//		}
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "上传成功",
//	})
//}

func EmployeeBasicExport(context *gin.Context) {
	fmt.Println("")
}

//func DepartmentTransfer(context *gin.Context) {
//	var err error
//	var department []Model.DepartmentTransferEmployeeTable
//	err = Global.Db.Find(&department).Error
//	if err != nil {
//		fmt.Println("查询所有部门调动信息出错", err)
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "查询所有部门调动信息出错",
//		})
//		return
//	}
//	context.JSON(http.StatusOK, department)
//}

//func DepartmentTransferSearch(context *gin.Context) {
//	var requestEmployee Model.DepartmentTransferEmployeeTable
//	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
//	fmt.Println("获取的搜索名字：", requestEmployee.Name)
//	var tempEmployee []Model.DepartmentTransferEmployeeTable
//	err := Global.Db.Where("name = ?", requestEmployee.Name).First(&tempEmployee).Error
//	if err != nil {
//		fmt.Println("未找到此员工")
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "未找到此员工",
//		})
//		return
//	}
//	fmt.Println("查询结果：", tempEmployee)
//	context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "查询成功",
//		"data": tempEmployee,
//	})
//}



//func PostTransferSearch(context *gin.Context) {
//	var requestEmployee Model.PostTransferEmployeeTable
//	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
//	fmt.Println("获取的搜索名字：", requestEmployee.Name)
//	var tempEmployee []Model.PostTransferEmployeeTable
//	err := Global.Db.Where("name = ?", requestEmployee.Name).Find(&tempEmployee).Error
//	if err != nil {
//		fmt.Println("未找到此员工")
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "未找到此员工",
//		})
//		return
//	}
//	fmt.Println("查询结果：", tempEmployee)
//	context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "查询成功",
//		"data": tempEmployee,
//	})
//}



	//写入员工离职表

//func AdminInitExamine(context *gin.Context) {
//	var requestDirector Model.Employee
//	json.NewDecoder(context.Request.Body).Decode(&requestDirector)
//	fmt.Println("请求到的总裁：", requestDirector)
//
//	var application []Model.PostApplications
//	var err error
//	err = Global.Db.Where("is_agree = ?", 0).Find(&application).Error
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


//编辑个人信息
func Edit(ctx *gin.Context) {
	var request Model.Employee
	//获取参数
	ctx.ShouldBind(&request)
	fmt.Println(request)
	var err error
	err = Global.Db.Where("id = ?", request.Id).Updates(&request).Error
	if err != nil {
		fmt.Println("修改个人信息出错", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改个人信息出错",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}
//初始化生成部门id
func initEmployeeId(count int) string {
	var result string
	result = "2021"
	//将count转为字符串
	temp := strconv.Itoa(count)
	length := len(temp)
	//计算count字符串长度，生成前缀0
	zero := "0"
	for i := 0; i < 4-length; i++ {
		result += zero
		//fmt.Println(i)
	}
	result += temp
	fmt.Println("生成的编号：", result)
	return result
}
//查询工资信息
func SearchEmpSalary(context *gin.Context) {
	var empsalary []Model.Salaytable
	var err error
	err = Global.Db.Find(&empsalary).Error
	if err != nil {
		fmt.Println("查询工资出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询工资表出错",
		})
		return
	}
	fmt.Println("所有工资信息：", empsalary)
	context.JSON(http.StatusOK, empsalary)
}



//查询部门信息
func DepartmentBasic(context *gin.Context) {
	var department []Model.Department
	var err error
	err = Global.Db.Find(&department).Error
	if err != nil {
		fmt.Println("查询部门出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询部门出错",
		})
		return
	}
	fmt.Println("所有部门信息：", department)
	context.JSON(http.StatusOK, department)
}
//初始化部门信息
func DepartmentInit(context *gin.Context) {

	var count int64
	var temp []Model.Department
	var err error
	err = Global.Db.Find(&temp).Count(&count).Error
	if err != nil {
		fmt.Println("初始化部门出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "初始化部门出错",
		})
		return
	}
	var department []string
	strInt64 := strconv.FormatInt(count, 10)
	tempCount, _ := strconv.Atoi(strInt64)
	for i := 0; i < tempCount; i++ {
		//de := temp[i].Department_name
		department = append(department, temp[i].Department_name)
		//department[i] = temp[i].Department_name
	}
	fmt.Println("找到的部门数组：", department)
	context.JSON(http.StatusOK, department)
}


//修改部门
func DepartmentUpdate(context *gin.Context) {
	var requestDepartment Model.Department
	json.NewDecoder(context.Request.Body).Decode(&requestDepartment)
	fmt.Println("即将修改的部门：", requestDepartment)
	var err error
	err = Global.Db.Where("department_number = ?", requestDepartment.Department_number).Updates(&requestDepartment).Error
	if err != nil {
		fmt.Println("修改部门出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改部门出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

//初始化员工工资表
func EmpSalaryInit(context *gin.Context) {

	var count int64
	var temp []Model.Salaytable
	var err error
	err = Global.Db.Find(&temp).Count(&count).Error
	if err != nil {
		fmt.Println("初始化员工工资表出错", err)
		context.JSON(http.StatusOK, gin.H{
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
		Empsalary = append(Empsalary, temp[i].Name_cx)
		//department[i] = temp[i].Department_name
	}
	fmt.Println("找到的部门数组：", Empsalary)
	context.JSON(http.StatusOK, Empsalary)
}

//func DepartmentDelete(context *gin.Context) {
//	var requestDepartment = Model.Department{}
//	json.NewDecoder(context.Request.Body).Decode(&requestDepartment)
//	fmt.Println("即将删除的部门：", requestDepartment)
//	var err error
//	var count int64
//	err = Global.Db.Table("post_tables").Where("department_number = ?", requestDepartment.Department_number).Count(&count).Error
//	if err != nil {
//		fmt.Println("计算部门下属岗位出错")
//		return
//	}
//	fmt.Println("此部门的岗位数量：", count)
//	if count > 0 {
//		fmt.Println("部门下有岗位，禁止删除")
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "此部门下岗位数量不为0，禁止删除",
//		})
//		return
//	}
//	err = Global.Db.Where("department_number = ?", requestDepartment.Department_number).Delete(&requestDepartment).Error
//	if err != nil {
//		fmt.Println("删除出错")
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "删除出错",
//		})
//		return
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "删除成功",
//	})
//}
//
//func DepartmentAdd(context *gin.Context) {
//	var requestDepartment Model.Department
//	json.NewDecoder(context.Request.Body).Decode(&requestDepartment)
//	fmt.Println("即将新增的部门：", requestDepartment)
//	var err error
//	//查询部门名称是否重复
//	err = Global.Db.Table("departments").Where("department_name = ?", requestDepartment.Department_name).Error
//	if err != nil {
//		fmt.Println("部门名称重复！")
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "部门名称已存在",
//		})
//		return
//	}
//	//根据现有部门数量生成部门编号
//	var tempDepartment Model.Department
//	Global.Db.Last(&tempDepartment)
//	var count int
//	count = tempDepartment.Department_number + 1
//	fmt.Println("生成的部门编号：", count)
//	var newDepartment Model.Department
//	newDepartment.Department_number = count
//	newDepartment.Department_name = requestDepartment.Department_name
//	newDepartment.Department_type = requestDepartment.Department_type
//	//默认部门主管是空
//	newDepartment.Department_head = "暂无"
//	newDepartment.Phone = requestDepartment.Phone
//	newDepartment.Fax = requestDepartment.Fax
//	newDepartment.Describe = requestDepartment.Describe
//	newDepartment.Superior_department = requestDepartment.Superior_department
//	newDepartment.Incorporation_date = requestDepartment.Incorporation_date
//	fmt.Println("即将新增的部门：", newDepartment)
//	err = Global.Db.Create(&newDepartment).Error
//	if err != nil {
//		fmt.Println("新增部门出错", err)
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "新增部门出错",
//		})
//		return
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "新增成功",
//	})
//
//}

//func PostInit(context *gin.Context) {
//
//	// 找所有的部门
//	//var post_names [][]string
//	var tempDepartmentCount int64 //部门数量
//	var tempDepartment []Model.Department
//	var err error
//	err = Global.Db.Find(&tempDepartment).Count(&tempDepartmentCount).Error
//	if err != nil {
//		fmt.Println("初始化部门出错", err)
//		context.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "初始化部门出错",
//		})
//		return
//	}
//	var department []string
//	strInt64 := strconv.FormatInt(tempDepartmentCount, 10)
//	DepartmentCount, _ := strconv.Atoi(strInt64) //部门数量
//	for i := 0; i < DepartmentCount; i++ {
//		//de := temp[i].Department_name
//		department = append(department, tempDepartment[i].Department_name)
//		//department[i] = temp[i].Department_name
//	}
//	fmt.Println("找到的部门数组：", department)
//
//	//post_names := map[string][]string{}
//	var post_names map[string][]string = make(map[string][]string)
//
//	for i := 0; i < DepartmentCount; i++ {
//		//根据部门名称查找岗位
//		var post []Model.PostTable
//		var temp Model.Department
//		err = Global.Db.Where("department_name = ?", department[i]).First(&temp).Error
//		if err != nil {
//			fmt.Println("错误1")
//		}
//		var postCount int64
//		err = Global.Db.Where("department_number = ?", temp.Department_number).Find(&post).Count(&postCount).Error
//		if err != nil {
//			fmt.Println("错误2")
//		}
//		var allPost []string
//		strInt64 = strconv.FormatInt(postCount, 10)
//		postCountTrue, _ := strconv.Atoi(strInt64)
//		for i := 0; i < postCountTrue; i++ {
//			allPost = append(allPost, post[i].Post_name)
//		}
//
//		post_names[department[i]] = allPost
//	}
//	fmt.Println("即将返回的重要数据：", post_names)
//	context.JSON(http.StatusOK, post_names)
//	// 根据所有的部门，查找每一个的所有岗位
//
//	// 把部门和岗位赋值给特殊数组
//}






func SearchByDepartmentNumber(context *gin.Context) {
	type temp struct {
		Text  string `json:"text"`
		Value int    `json:"value"`
	}

	var tempDepartment []Model.Department
	var err error
	var tempCount int64
	err = Global.Db.Table("departments").Find(&tempDepartment).Count(&tempCount).Error
	if err != nil {
		fmt.Println("查找部门数量出错", err)
	}
	strInt64 := strconv.FormatInt(tempCount, 10)
	count, _ := strconv.Atoi(strInt64)
	fmt.Println("部门数量：", count)
	var departmentNumber []temp
	for i := 0; i < count; i++ {
		//text := tempDepartment[i].Department_number
		var t temp
		t.Text = strconv.Itoa(tempDepartment[i].Department_number) + ":" + tempDepartment[i].Department_name
		t.Value = tempDepartment[i].Department_number
		departmentNumber = append(departmentNumber, t)
	}
	fmt.Println("数组：", departmentNumber)
	context.JSON(http.StatusOK, departmentNumber)
}



