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
func DirectorInitExamine(context *gin.Context) {
	var requestDirector Model.Employee
	json.NewDecoder(context.Request.Body).Decode(&requestDirector)
	fmt.Println("请求到的主管：", requestDirector)
	var application []Model.PostApplications
	var err error
	err = Global.Db.Where("new_department_name = ? and is_agree = ?", requestDirector.Department_name, 0).Find(&application).Error
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

//同意申请
func DirectorAgreeApply(context *gin.Context) {
	var requestApplication Model.PostApplications
	json.NewDecoder(context.Request.Body).Decode(&requestApplication)
	fmt.Println("收到的同意申请：", requestApplication)
	var err error
	//把申请表的是否同意置为1
	//var application Model.PostApplications
	//设置Where条件，Model参数绑定一个空的模型变量
	//等价于: UPDATE `post_applications`   //把申请表的是否同意置为1
	err = Global.Db.Table("post_applications").Where("id = ?", requestApplication.Id).Update("is_agree", 1).Error
	if err != nil {
		fmt.Println("修改申请表出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改申请表出错",
		})
		return
	}
	//修改对应员工的部门信息
	var employee Model.Employee
	err = Global.Db.Where("id = ?", requestApplication.Id).First(&employee).Error
	if err != nil {
		fmt.Println("查找员工出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找员工出错",
		})
		return
	}
	fmt.Println("即将修改的员工信息：", employee)
	//原始                               新的
	original_department_number := employee.Department_number
	original_post_number := employee.Post_number
	original_department_name := requestApplication.Original_department_name
	original_post_name := requestApplication.Original_post_name

	// 查找正确的部门和岗位编号
	var department Model.Department
	var post Model.PostTable
	err = Global.Db.Where("department_name = ?", requestApplication.New_department_name).First(&department).Error
	if err != nil {
		fmt.Println("查找部门编号出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找部门编号出错",
		})
		return
	}
	err = Global.Db.Where("post_name = ?", requestApplication.New_post_name).First(&post).Error
	if err != nil {
		fmt.Println("查找岗位编号出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找岗位编号出错",
		})
		return
	}
	employee.Department_number = department.Department_number
	employee.Post_number = post.Post_number
	employee.Department_name = requestApplication.New_department_name
	employee.Post_name = requestApplication.New_post_name
	err = Global.Db.Where("id = ?", requestApplication.Id).Updates(&employee).Error
	if err != nil {
		fmt.Println("修改员工出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改员工出错",
		})
		return
	}
	fmt.Println("修改之后的员工信息：", employee)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "已同意" + requestApplication.Name + "的申请",
		"data": employee,
	})
	//把岗位变动写入数据库
	if employee.Department_name != original_department_name { //如果更改的数据与原数据不相同
		//跨部门调动
		var tempCount int64
		err = Global.Db.Table("department_transfer_employee_tables").Count(&tempCount).Error
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
		var departmentTransfer Model.DepartmentTransferEmployeeTable
		departmentTransfer.Transfer_id = count
		departmentTransfer.Id = employee.Id
		departmentTransfer.Name = employee.Name
		departmentTransfer.Sex = employee.Sex
		departmentTransfer.Original_department_number = original_department_number
		departmentTransfer.Original_department_name = original_department_name
		departmentTransfer.New_department_number = department.Department_number
		departmentTransfer.New_department_name = requestApplication.New_department_name
		departmentTransfer.Transfer_date = pgtype.Date{time.Now(), 2, 0}
		departmentTransfer.Reasons_for_transfer = requestApplication.Reason_for_application

		fmt.Println("即将写入的部门调动信息：", departmentTransfer)
		Global.Db.Create(&departmentTransfer)

		var tempCountPost int64
		err = Global.Db.Table("post_transfer_employee_tables").Count(&tempCountPost).Error //Count直接返回查询匹配的行数
		if err != nil {
			fmt.Println("计算记录数量出错！", err)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "计算记录数量出错",
			})
			return
		}
		fmt.Println("记录数量：", tempCountPost)
		strInt64Post := strconv.FormatInt(tempCountPost, 10)
		countPost, _ := strconv.Atoi(strInt64Post)
		var postTransfer Model.PostTransferEmployeeTable
		postTransfer.Transfer_id = countPost
		postTransfer.Id = employee.Id
		postTransfer.Name = employee.Name
		postTransfer.Sex = employee.Sex
		postTransfer.Original_post_number = original_post_number
		postTransfer.Original_post_name = original_post_name
		postTransfer.New_post_number = post.Post_number
		postTransfer.New_post_name = requestApplication.New_post_name
		postTransfer.Transfer_date = departmentTransfer.Transfer_date
		postTransfer.Reasons_for_transfer = requestApplication.Reason_for_application

		fmt.Println("即将写入的岗位调动信息：", postTransfer)
		Global.Db.Create(&postTransfer)
	} else {
		//部门内部调动
		var tempCount int64
		err = Global.Db.Table("post_transfer_employee_tables").Count(&tempCount).Error
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
		var postTransfer Model.PostTransferEmployeeTable
		postTransfer.Transfer_id = count
		postTransfer.Id = employee.Id
		postTransfer.Name = employee.Name
		postTransfer.Sex = employee.Sex
		postTransfer.Original_post_number = original_post_number
		postTransfer.Original_post_name = original_post_name
		postTransfer.New_post_number = post.Post_number
		postTransfer.New_post_name = requestApplication.New_post_name
		postTransfer.Transfer_date = pgtype.Date{time.Now(), 2, 0}
		postTransfer.Reasons_for_transfer = requestApplication.Reason_for_application

		fmt.Println("即将写入的岗位调动信息：", postTransfer)
		Global.Db.Create(&postTransfer)
	}

}

func DirectorRefuseApply(context *gin.Context) {
	var requestApplication Model.PostApplications
	json.NewDecoder(context.Request.Body).Decode(&requestApplication)
	fmt.Println("收到的拒绝申请：", requestApplication)
	var err error
	err = Global.Db.Table("post_applications").Where("id = ?", requestApplication.Id).Update("is_agree", -1).Error
	if err != nil {
		fmt.Println("修改申请表出错")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改申请表出错",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "已拒绝" + requestApplication.Name + "的申请",
		"data": requestApplication,
	})
}

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

func EmployeeQuitByDirector(context *gin.Context) {
	request := make(map[string]interface{})
	context.ShouldBind(&request)
	requestId := request["id"].(string)
	fmt.Println("登录的主管ID：", requestId)

	var director Model.Employee
	var err error
	err = Global.Db.Where("id = ?", requestId).First(&director).Error
	if err != nil {
		fmt.Println("查找主管ID出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找主管ID出错",
		})
		return
	}

	var quit []Model.ResignedEmployee
	err = Global.Db.Where("department_name = ?", director.Department_name).Find(&quit).Error
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

func EmployeeBasicByDirectorBySearchByDate(context *gin.Context) {
	requestDate := make(map[string]interface{})
	context.ShouldBind(&requestDate)
	requestId := requestDate["id"].(string)
	fmt.Println("登录的主管ID：", requestId)
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

	var director Model.Employee
	var err error
	err = Global.Db.Where("id = ?", requestId).First(&director).Error
	if err != nil {
		fmt.Println("查找主管ID出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找主管ID出错",
		})
		return
	}

	var quit []Model.ResignedEmployee
	err = Global.Db.Where("department_name = ? and resignation_date >= ? and resignation_date <= ?", director.Department_name, start, end).Find(&quit).Error
	if err != nil {
		fmt.Println("查找离职员工表出错", err)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找离职员工表出错",
		})
		return
	}
	fmt.Println("即将返回的离职员工表：", quit)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": quit,
	})

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
func TalentApplyPart(ctx *gin.Context) {
	var requestDirector Model.Employee
	json.NewDecoder(ctx.Request.Body).Decode(&requestDirector)
	fmt.Println("请求到的总管：", requestDirector)

	var application []Model.TalentPool
	var err error
	err = Global.Db.Where("is_agree = ? and department_name = ?", 0, requestDirector.Department_name).Find(&application).Error
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

//根据入职日期查询按钮
func EmployeeBasicSearchDateByDirector(context *gin.Context) {
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
