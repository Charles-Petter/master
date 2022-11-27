package Controller

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
	"cx/Global"
	"cx/Model"

	"cx/Response"
	"net/http"
	"path"
	"strconv"
	"time"
)

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

func EmployeeBasicUpdate(context *gin.Context) {
	var requestEmployee = Model.Employee{}
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)

	if requestEmployee.Is_quit == "是" {
		//写入离职表
		var quit Model.ResignedEmployee
		var err error
		quit.Id = requestEmployee.Id
		quit.Name = requestEmployee.Name
		quit.Sex = requestEmployee.Sex
		quit.Resignation_date = pgtype.Date{time.Now(), 2, 0}
		//quit.Reasons_for_resignation = "总裁手动离职"
		quit.Department_number = requestEmployee.Department_number
		quit.Department_name = requestEmployee.Department_name
		quit.Post_number = requestEmployee.Post_number
		quit.Post_name = requestEmployee.Post_name
		err = Global.Db.Create(&quit).Error
		if err != nil {
			fmt.Println("创建离职员工信息记录出错", err)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "创建离职员工信息出错",
			})
			return
		}
	} else {
		//从离职表删除
		var quit Model.ResignedEmployee
		var err error
		quit.Id = requestEmployee.Id
		err = Global.Db.Table("resigned_employees").Where("id = ?", quit.Id).Delete(&quit).Error
		if err != nil {
			fmt.Println("删除离职员工信息出错", err)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "删除离职员工信息出错",
			})
			return
		}
	}

	fmt.Println("即将修改的用户：", requestEmployee)
	err := Global.Db.Where("id = ?", requestEmployee.Id).Updates(&requestEmployee).Error
	if err != nil {
		fmt.Println("修改出错！")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}
//在basic组件中删除员工
func EmployeeBasicDelete(context *gin.Context) {
	var requestEmployee = Model.Employee{}
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
	fmt.Println("即将删除的员工：", requestEmployee)
	err := Global.Db.Where("id = ?", requestEmployee.Id).Delete(&requestEmployee).Error
	if err != nil {
		fmt.Println("删除出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func EmployeeBasicMultiDelete(context *gin.Context) {
	var request []string
	count := 0
	context.ShouldBindJSON(&request)
	fmt.Println(request)
	for key := range request {
		fmt.Println(key)
		var err error
		var employee Model.Employee
		err = Global.Db.Where("id = ?", request[key]).First(&employee).Error
		if err != nil {
			fmt.Println("查找员工出错", err)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "查找员工出错",
			})
			return
		}
		err = Global.Db.Table("employees").Where("id = ?", request[key]).Delete(&employee).Error
		if err != nil {
			fmt.Println("删除员工出错", err)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "删除员工出错",
			})
			return
		}
		count++
	}
	fmt.Println(count)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
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
	err = Global.Db.Where("post_number = ?", requestEmployee.Post_number).First(&Model.PostTable{}).Error
	if err != nil {
		fmt.Println("找不到此岗位编号")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "岗位编号不存在",
		})
		return
	}
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
	var tempPost Model.PostTable // 将Model层PostTable表定义为一个临时变量tempPost
	Global.Db.Where("post_name = ?", requestEmployee.Post_name).First(&tempPost)
	if requestEmployee.Post_number != tempPost.Post_number {
		fmt.Println("岗位编号与名称不对应")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "岗位编号与名称不对应",
		})
		return
	}
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

//查询指定日期间的数据
func EmployeeBasicSearchDate(context *gin.Context) {
	requestDate := make(map[string]interface{})
	context.ShouldBind(&requestDate)
	fmt.Println(requestDate)
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
	var employee []Model.Employee
	err = Global.Db.Where("entry_date >= ? and entry_date <= ?", start, end).Find(&employee).Error
	if err != nil {
		fmt.Println("查找指定日期区间的员工出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找指定日期区间的员工出错",
		})
		return
	}
	fmt.Println("查找结果：", employee)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查找成功",
		"data": employee,
	})
}

func EmployeeBasicSearchAdvance(context *gin.Context) {
	requestSearch := make(map[string]interface{})
	context.ShouldBind(&requestSearch)
	requestPolitical := requestSearch["political"].(string)
	requestNation := requestSearch["nation"].(string)
	requestDepartment := requestSearch["department_name"].(string)
	requestPost := requestSearch["post_name"].(string)
	requestSchool := requestSearch["graduation_school"].(string)
	requestMajor := requestSearch["major_studied"].(string)
	fmt.Println("高级搜索：", requestPolitical, requestNation, requestDepartment, requestPost, requestSchool, requestMajor)

	var tempEmployee []Model.Employee
	err := Global.Db.Where("political = ? and nation = ? and department_name = ? and post_name = ? and graduation_school = ? and major_studied = ?", requestPolitical, requestNation, requestDepartment, requestPost, requestSchool, requestMajor).First(&tempEmployee).Error
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

func EmployeeBasicImport(context *gin.Context) {
	fmt.Println("进入文件上传")
	files, _ := context.FormFile("file")
	dst := path.Join("../", files.Filename)
	a := context.SaveUploadedFile(files, dst)
	fmt.Println("进入文件上传2")
	if a != nil {
		fmt.Println("a = ", a)
		context.JSON(http.StatusOK, gin.H{
			"code": 203,
			"msg":  a,
		})
		return
	}
	fmt.Println("进入文件上传3")
	xlsx, err := excelize.OpenFile(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := xlsx.GetRows("Sheet" + "1")
	fmt.Println("进入文件上传4")
	for key, row := range rows {
		if key > 0 {
			fmt.Println("第五列", row[5])
			Birthday, _ := time.Parse("2006-01-02", row[5])
			fmt.Println("生日：", Birthday)
			Entry_date, _ := time.Parse("2006-01-02", row[21])
			Graduation_date, _ := time.Parse("2006-01-02", row[27])
			Height, _ := strconv.Atoi(row[12])
			Department_number, _ := strconv.Atoi(row[17])
			Post_number, _ := strconv.Atoi(row[19])
			employee := Model.Employee{
				Id: row[0], Password: row[1], Employee_type: row[2],
				Name: row[3], Sex: row[4], Birthday: pgtype.Date{Birthday, 2, 0},
				Id_card: row[6], Political: row[7], Nation: row[8],
				Native_place: row[9], Phone: row[10], Email: row[11],
				Height: Height, Blood_type: row[13], Marital_status: row[14],
				Birthplace: row[15], Registered_residence: row[16], Department_number: Department_number,
				Department_name: row[18], Post_number: Post_number, Post_name: row[20],
				Entry_date: pgtype.Date{Entry_date, 2, 0}, Employment_form: row[22], Personnel_source: row[23],
				Highest_education: row[24], Graduation_school: row[25], Major_studied: row[26],
				Graduation_date: pgtype.Date{Graduation_date, 2, 0}, Is_quit: row[28],
			}
			fmt.Println("employee = ", employee)
			err := Global.Db.Create(&employee).Error
			if err != nil {
				fmt.Println("发现错误！", err)
				context.JSON(http.StatusOK, gin.H{
					"code": 202,
					"msg":  "员工:" + employee.Name + "已经存在",
				})
				return
			}
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
	})
}

func EmployeeBasicExport(context *gin.Context) {
	fmt.Println("")
}

func DepartmentTransfer(context *gin.Context) {
	var err error
	var department []Model.DepartmentTransferEmployeeTable
	err = Global.Db.Find(&department).Error
	if err != nil {
		fmt.Println("查询所有部门调动信息出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询所有部门调动信息出错",
		})
		return
	}
	context.JSON(http.StatusOK, department)
}

func DepartmentTransferSearch(context *gin.Context) {
	var requestEmployee Model.DepartmentTransferEmployeeTable
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
	fmt.Println("获取的搜索名字：", requestEmployee.Name)
	var tempEmployee []Model.DepartmentTransferEmployeeTable
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

func PostTransfer(context *gin.Context) {
	var err error
	var post []Model.PostTransferEmployeeTable
	err = Global.Db.Find(&post).Error
	if err != nil {
		fmt.Println("查询所有岗位调动信息出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询所有岗位调动信息出错",
		})
		return
	}
	context.JSON(http.StatusOK, post)
}

func PostTransferSearch(context *gin.Context) {
	var requestEmployee Model.PostTransferEmployeeTable
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
	fmt.Println("获取的搜索名字：", requestEmployee.Name)
	var tempEmployee []Model.PostTransferEmployeeTable
	err := Global.Db.Where("name = ?", requestEmployee.Name).Find(&tempEmployee).Error
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

func EmployeeQuit(context *gin.Context) {
	var quit []Model.ResignedEmployee
	var err error
	err = Global.Db.Find(&quit).Error
	if err != nil {
		fmt.Println("查找离职员工表出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找离职员工表出错",
		})
		return
	}
	fmt.Println("即将返回的离职员工表：", quit)
	context.JSON(http.StatusOK, quit)
}

func EmployeeQuitSearchByDate(context *gin.Context) {
	requestDate := make(map[string]interface{})
	context.ShouldBind(&requestDate)
	fmt.Println(requestDate)
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
	var resigned []Model.ResignedEmployee
	err = Global.Db.Where("resignation_date >= ? and resignation_date <= ?", start, end).Find(&resigned).Error
	if err != nil {
		fmt.Println("日期筛选出错", err)
		return
	}
	fmt.Println(resigned)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": resigned,
	})
}

func ExpelEmployee(context *gin.Context) {
	var requestEmployee Model.Employee
	json.NewDecoder(context.Request.Body).Decode(&requestEmployee)
	fmt.Println("即将开除的员工信息：", requestEmployee)
	var err error
	//将基本信息的是否离职字段设置为 是
	err = Global.Db.Table("employees").Where("id = ?", requestEmployee.Id).Update("is_quit", "是").Error
	if err != nil {
		fmt.Println("开除员工出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "开除员工出错",
		})
		return
	}
	//写入员工离职表
	var quit Model.ResignedEmployee
	quit.Id = requestEmployee.Id
	quit.Name = requestEmployee.Name
	quit.Sex = requestEmployee.Sex
	quit.Resignation_date = pgtype.Date{time.Now(), 2, 0}
	quit.Reasons_for_resignation = "总裁手动开除"
	quit.Department_number = requestEmployee.Department_number
	quit.Department_name = requestEmployee.Department_name
	quit.Post_number = requestEmployee.Post_number
	quit.Post_name = requestEmployee.Post_name
	err = Global.Db.Create(&quit).Error
	if err != nil {
		fmt.Println("创建离职员工信息记录出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "创建离职员工信息出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "开除成功",
	})
}

func AdminInitExamine(context *gin.Context) {
	var requestDirector Model.Employee
	json.NewDecoder(context.Request.Body).Decode(&requestDirector)
	fmt.Println("请求到的总裁：", requestDirector)

	var application []Model.PostApplications
	var err error
	err = Global.Db.Where("is_agree = ?", 0).Find(&application).Error
	if err != nil {
		fmt.Println("寻找初始审核信息出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "初始化审核信息出错",
		})
		return
	}
	fmt.Println("即将返回的申请数据：", application)
	context.JSON(http.StatusOK, application)
}

func TalentPoolAdd(context *gin.Context) {
	var requestTalent Model.TalentPool
	json.NewDecoder(context.Request.Body).Decode(&requestTalent)
	fmt.Println("人才：", requestTalent)
	var err error

	var tempCount int64
	err = Global.Db.Table("talent_pools").Count(&tempCount).Error
	if err != nil {
		fmt.Println("计算记录数量出错！", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "计算记录数量出错",
		})
		return
	}
	fmt.Println("记录数量：", tempCount)
	strInt64 := strconv.FormatInt(tempCount, 10)
	count, _ := strconv.Atoi(strInt64)
	requestTalent.Talent_id = count

	err = Global.Db.Create(&requestTalent).Error
	if err != nil {
		fmt.Println("创建人才失败")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "创建人才失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "已成功加入我司人才部，请等待审核",
	})
}

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

/*人才申请信息*/
func TalentApply(ctx *gin.Context) {
	var requestDirector Model.Employee
	json.NewDecoder(ctx.Request.Body).Decode(&requestDirector)
	fmt.Println("请求到的总裁：", requestDirector)

	var application []Model.TalentPool
	var err error
	err = Global.Db.Where("is_agree = ?", 0).Find(&application).Error
	if err != nil {
		fmt.Println("寻找初始审核信息出错", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "初始化审核信息出错",
		})
		return
	}
	fmt.Println("即将返回的申请数据：", application)
	ctx.JSON(http.StatusOK, application)
}

func TalentAgreeApply(ctx *gin.Context) {
	var requestApplication Model.TalentPool
	json.NewDecoder(ctx.Request.Body).Decode(&requestApplication)
	fmt.Println("收到的同意申请：", requestApplication)
	var err error
	//把人才库申请表的是否同意置为1
	//var application Model.PostApplications
	err = Global.Db.Table("talent_pools").Where("talent_id = ?", requestApplication.Talent_id).Update("is_agree", 1).Error
	if err != nil {
		fmt.Println("修改申请表出错")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改申请表出错",
		})
		return
	}
	var employee Model.Employee
	Global.Db.Last(&employee)
	// 查找正确的部门和岗位编号
	var department Model.Department
	var post Model.PostTable
	err = Global.Db.Where("department_name = ?", requestApplication.Department_name).Find(&department).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改申请表出错",
		})
		return
	}
	err = Global.Db.Where("post_name = ?", requestApplication.Post_name).Find(&post).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改申请表出错",
		})
		return
	}
	temp, _ := strconv.Atoi(employee.Id)
	temp = temp + 1
	Id := strconv.Itoa(temp)
	Post_name := requestApplication.Post_name
	Department_name := requestApplication.Department_name
	Birthday := requestApplication.Birthday
	Password := "123456"
	Name := requestApplication.Name
	Birthplace := requestApplication.Birthplace
	Email := requestApplication.Email
	Employee_type := "员工"
	Blood_type := requestApplication.Blood_type
	Sex := requestApplication.Sex
	Registered_residence := requestApplication.Registered_residence
	Political := requestApplication.Political
	Phone := requestApplication.Phone
	Personnel_source := requestApplication.Personnel_source
	Native_place := requestApplication.Native_place
	Nation := requestApplication.Nation
	Marital_status := requestApplication.Marital_status
	Major_studied := requestApplication.Major_studied
	Is_quit := "否"
	Id_card := requestApplication.Id_card
	Highest_education := requestApplication.Highest_education
	Height := requestApplication.Height
	Graduation_school := requestApplication.Graduation_school
	Graduation_date := requestApplication.Graduation_date
	Entry_date := pgtype.Date{time.Now(), 2, 0}
	Employment_form := requestApplication.Employment_form
	Department_number := department.Department_number
	Post_number := post.Post_number
	newEmployee := Model.Employee{
		Id:                   Id,
		Password:             Password,
		Employee_type:        Employee_type,
		Name:                 Name,
		Sex:                  Sex,
		Birthday:             Birthday,
		Id_card:              Id_card,
		Political:            Political,
		Nation:               Nation,
		Native_place:         Native_place,
		Phone:                Phone,
		Email:                Email,
		Height:               Height,
		Blood_type:           Blood_type,
		Marital_status:       Marital_status,
		Birthplace:           Birthplace,
		Registered_residence: Registered_residence,
		Department_number:    Department_number,
		Department_name:      Department_name,
		Post_number:          Post_number,
		Post_name:            Post_name,
		Entry_date:           Entry_date,
		Employment_form:      Employment_form,
		Personnel_source:     Personnel_source,
		Highest_education:    Highest_education,
		Graduation_school:    Graduation_school,
		Major_studied:        Major_studied,
		Graduation_date:      Graduation_date,
		Is_quit:              Is_quit,
	}
	fmt.Println("人才库员工信息", newEmployee)
	Global.Db.Create(&newEmployee)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "已同意" + requestApplication.Name + "的申请",
		"data": employee,
	})
}

func TalentRefuseApply(ctx *gin.Context) {
	var requestApplication Model.TalentPool
	json.NewDecoder(ctx.Request.Body).Decode(&requestApplication)
	fmt.Println("收到的拒绝申请：", requestApplication)
	var err error
	err = Global.Db.Table("talent_pools").Where("talent_id  = ?", requestApplication.Talent_id).Update("is_agree", -1).Error
	if err != nil {
		fmt.Println("修改申请表出错")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改申请表出错",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "已拒绝" + requestApplication.Name + "的申请",
		"data": requestApplication,
	})
}

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

func DepartmentPost(context *gin.Context) {
	var requestDepartment Model.Department
	json.NewDecoder(context.Request.Body).Decode(&requestDepartment)
	fmt.Println("请求的部门：", requestDepartment)
	var post []Model.PostTable
	var err error
	err = Global.Db.Where("department_number = ?", requestDepartment.Department_number).Find(&post).Error
	if err != nil {
		fmt.Println("查找岗位出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找岗位出错",
		})
		return
	}
	fmt.Println("找到的岗位：", post)
	context.JSON(http.StatusOK, post)
}

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

func DepartmentDelete(context *gin.Context) {
	var requestDepartment = Model.Department{}
	json.NewDecoder(context.Request.Body).Decode(&requestDepartment)
	fmt.Println("即将删除的部门：", requestDepartment)
	var err error
	var count int64
	err = Global.Db.Table("post_tables").Where("department_number = ?", requestDepartment.Department_number).Count(&count).Error
	if err != nil {
		fmt.Println("计算部门下属岗位出错")
		return
	}
	fmt.Println("此部门的岗位数量：", count)
	if count > 0 {
		fmt.Println("部门下有岗位，禁止删除")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "此部门下岗位数量不为0，禁止删除",
		})
		return
	}
	err = Global.Db.Where("department_number = ?", requestDepartment.Department_number).Delete(&requestDepartment).Error
	if err != nil {
		fmt.Println("删除出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func DepartmentAdd(context *gin.Context) {
	var requestDepartment Model.Department
	json.NewDecoder(context.Request.Body).Decode(&requestDepartment)
	fmt.Println("即将新增的部门：", requestDepartment)
	var err error
	//查询部门名称是否重复
	err = Global.Db.Table("departments").Where("department_name = ?", requestDepartment.Department_name).Error
	if err != nil {
		fmt.Println("部门名称重复！")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "部门名称已存在",
		})
		return
	}
	//根据现有部门数量生成部门编号
	var tempDepartment Model.Department
	Global.Db.Last(&tempDepartment)
	var count int
	count = tempDepartment.Department_number + 1
	fmt.Println("生成的部门编号：", count)
	var newDepartment Model.Department
	newDepartment.Department_number = count
	newDepartment.Department_name = requestDepartment.Department_name
	newDepartment.Department_type = requestDepartment.Department_type
	//默认部门主管是空
	newDepartment.Department_head = "暂无"
	newDepartment.Phone = requestDepartment.Phone
	newDepartment.Fax = requestDepartment.Fax
	newDepartment.Describe = requestDepartment.Describe
	newDepartment.Superior_department = requestDepartment.Superior_department
	newDepartment.Incorporation_date = requestDepartment.Incorporation_date
	fmt.Println("即将新增的部门：", newDepartment)
	err = Global.Db.Create(&newDepartment).Error
	if err != nil {
		fmt.Println("新增部门出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "新增部门出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "新增成功",
	})

}

func PostInit(context *gin.Context) {

	// 找所有的部门
	//var post_names [][]string
	var tempDepartmentCount int64 //部门数量
	var tempDepartment []Model.Department
	var err error
	err = Global.Db.Find(&tempDepartment).Count(&tempDepartmentCount).Error
	if err != nil {
		fmt.Println("初始化部门出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "初始化部门出错",
		})
		return
	}
	var department []string
	strInt64 := strconv.FormatInt(tempDepartmentCount, 10)
	DepartmentCount, _ := strconv.Atoi(strInt64) //部门数量
	for i := 0; i < DepartmentCount; i++ {
		//de := temp[i].Department_name
		department = append(department, tempDepartment[i].Department_name)
		//department[i] = temp[i].Department_name
	}
	fmt.Println("找到的部门数组：", department)

	//post_names := map[string][]string{}
	var post_names map[string][]string = make(map[string][]string)

	for i := 0; i < DepartmentCount; i++ {
		//根据部门名称查找岗位
		var post []Model.PostTable
		var temp Model.Department
		err = Global.Db.Where("department_name = ?", department[i]).First(&temp).Error
		if err != nil {
			fmt.Println("错误1")
		}
		var postCount int64
		err = Global.Db.Where("department_number = ?", temp.Department_number).Find(&post).Count(&postCount).Error
		if err != nil {
			fmt.Println("错误2")
		}
		var allPost []string
		strInt64 = strconv.FormatInt(postCount, 10)
		postCountTrue, _ := strconv.Atoi(strInt64)
		for i := 0; i < postCountTrue; i++ {
			allPost = append(allPost, post[i].Post_name)
		}

		post_names[department[i]] = allPost
	}
	fmt.Println("即将返回的重要数据：", post_names)
	context.JSON(http.StatusOK, post_names)
	// 根据所有的部门，查找每一个的所有岗位

	// 把部门和岗位赋值给特殊数组
}

func PostAdd(context *gin.Context) {
	var requestPost Model.PostTable
	json.NewDecoder(context.Request.Body).Decode(&requestPost)
	fmt.Println("请求添加的岗位：", requestPost)
	var err error
	var tempPost Model.PostTable
	err = Global.Db.Where("department_number = ?", requestPost.Department_number).Last(&tempPost).Error
	//此部门下的岗位数量是否为0
	if err != nil { //是
		fmt.Println("此部门下还没有岗位", err)
		//统计现有部门数量来生成岗位编号
		var tempDepartmentCount int64
		err = Global.Db.Table("departments").Count(&tempDepartmentCount).Error
		if err != nil {
			fmt.Println("统计部门数量出错")
			return
		}
		strInt64 := strconv.FormatInt(tempDepartmentCount, 10)
		departmentCount, _ := strconv.Atoi(strInt64)
		fmt.Println("部门数量：", departmentCount)
		newPostNumber := departmentCount * 10
		fmt.Println("空部门下生成的岗位编号：", newPostNumber)
		var newPost Model.PostTable
		newPost.Post_number = newPostNumber
		newPost.Post_name = requestPost.Post_name
		newPost.Department_number = requestPost.Department_number
		newPost.Post_type = requestPost.Post_type
		newPost.Post_establishment = requestPost.Post_establishment
		err = Global.Db.Create(&newPost).Error
		if err != nil {
			fmt.Println("新建岗位出错", err)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "新建岗位出错",
			})
			return
		}
	} else {
		fmt.Println("此部门下已有岗位")
		//统计此部门下岗位数量来生成岗位编号
		var tempPostCount int64
		err = Global.Db.Table("post_tables").Where("department_number = ?", requestPost.Department_number).Count(&tempPostCount).Error
		if err != nil {
			fmt.Println("统计岗位数量出错", err)
			return
		}
		strInt64 := strconv.FormatInt(tempPostCount, 10)
		postCount, _ := strconv.Atoi(strInt64)
		fmt.Println("岗位数量：", postCount)
		if postCount == 10 {
			fmt.Println("此部门的岗位数量已达上限")
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "此部门的岗位数量已达上限",
			})
			return
		}
		var tempPost Model.PostTable
		err = Global.Db.Where("department_number = ?", requestPost.Department_number).First(&tempPost).Order("post_number").Error
		newPostNumber := tempPost.Post_number + postCount
		fmt.Println("生成的岗位编号：", newPostNumber)
		var newPost Model.PostTable
		newPost.Post_number = newPostNumber
		newPost.Post_name = requestPost.Post_name
		newPost.Department_number = requestPost.Department_number
		newPost.Post_type = requestPost.Post_type
		newPost.Post_establishment = requestPost.Post_establishment
		err = Global.Db.Create(&newPost).Error
		if err != nil {
			fmt.Println("新建岗位出错", err)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "新建岗位出错",
			})
			return
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "新增成功",
	})
}

func PostUpdate(context *gin.Context) {
	var requestPost Model.PostTable
	json.NewDecoder(context.Request.Body).Decode(&requestPost)
	fmt.Println("请求编辑的岗位：", requestPost)
	var err error
	err = Global.Db.Where("post_number = ?", requestPost.Post_number).Updates(&requestPost).Error
	if err != nil {
		fmt.Println("编辑出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "编辑岗位出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

func PostDelete(context *gin.Context) {
	var requestPost = Model.PostTable{}
	json.NewDecoder(context.Request.Body).Decode(&requestPost)
	fmt.Println("即将删除的岗位：", requestPost)
	var err error
	var count int64
	err = Global.Db.Table("employees").Where("post_number = ?", requestPost.Post_number).Count(&count).Error
	if err != nil {
		fmt.Println("查询岗位下属员工数量出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询岗位下属员工数量出错",
		})
		return
	}
	fmt.Println("此岗位的员工数量：", count)
	if count > 0 {
		fmt.Println("此岗位下员工数量不为0，禁止删除")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "此岗位下员工数量不为0，禁止删除",
		})
		return
	}
	err = Global.Db.Where("post_number = ?", requestPost.Post_number).Delete(&requestPost).Error
	if err != nil {
		fmt.Println("删除出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

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

func SearchByPostNumber(context *gin.Context) {
	type temp struct {
		Text  string `json:"text"`
		Value int    `json:"value"`
	}

	var tempPost []Model.PostTable
	var err error
	var tempCount int64
	err = Global.Db.Table("post_tables").Find(&tempPost).Count(&tempCount).Error
	if err != nil {
		fmt.Println("查找岗位数量出错", err)
	}
	strInt64 := strconv.FormatInt(tempCount, 10)
	count, _ := strconv.Atoi(strInt64)
	fmt.Println("岗位数量：", count)
	var postNumber []temp
	for i := 0; i < count; i++ {
		//text := tempDepartment[i].Department_number
		var t temp
		t.Text = strconv.Itoa(tempPost[i].Post_number) + ":" + tempPost[i].Post_name
		t.Value = tempPost[i].Post_number
		postNumber = append(postNumber, t)
	}
	fmt.Println("数组：", postNumber)
	context.JSON(http.StatusOK, postNumber)
}

func DepartmentTransferSearchByDate(context *gin.Context) {
	requestDate := make(map[string]interface{})
	context.ShouldBind(&requestDate)
	fmt.Println(requestDate)
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
	var department []Model.DepartmentTransferEmployeeTable
	err = Global.Db.Where("transfer_date >= ? and transfer_date <= ?", start, end).Find(&department).Error
	if err != nil {
		fmt.Println("日期筛选出错", err)
		return
	}
	fmt.Println(department)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": department,
	})

}

func PostTransferSearchByDate(context *gin.Context) {
	requestDate := make(map[string]interface{})
	context.ShouldBind(&requestDate)
	fmt.Println(requestDate)
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
	var post []Model.PostTransferEmployeeTable
	err = Global.Db.Where("transfer_date >= ? and transfer_date <= ?", start, end).Find(&post).Error
	if err != nil {
		fmt.Println("日期筛选出错", err)
		return
	}
	fmt.Println(post)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": post,
	})
}
