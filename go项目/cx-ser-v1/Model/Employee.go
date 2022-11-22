package Model

import "github.com/jackc/pgtype"

type Test struct {
	Id int `json:"id"`
}

/*员工表*/
type Employee struct {
	Id                   string `json:"id"`                   //员工编号
	Password             string `json:"password"`             //密码
	Employee_type        string    `json:"employee_type"`        //员工类型
	Name                 string `json:"name"`                 //姓名
	Sex                  string `json:"sex"`                  //性别
	Birthday             pgtype.Date `json:"birthday"`             //出生日期
	Id_card              string `json:"id_card"`              //身份证号
	Political            string `json:"political"`            //政治面貌
	Nation               string `json:"nation"`               //民族
	Native_place         string `json:"native_place"`         //籍贯
	Phone                string `json:"phone"`                //电话
	Email                string `json:"email"`                //电子邮箱
	Height               int    `json:"height"`               //身高
	Blood_type           string `json:"blood_type"`           //血型
	Marital_status       string `json:"marital_status"`       //婚姻状况
	Birthplace           string `json:"birthplace"`           //出生地
	Registered_residence string `json:"registered_residence"` //户口所在地
	Department_number    int    `json:"department_number"`    //部门编号
	Department_name      string `json:"department_name"`      //部门名称
	Post_number          int    `json:"post_number"`          //岗位编号
	Post_name            string `json:"post_name"`            //岗位名称
	Entry_date           pgtype.Date `json:"entry_date"`           //入职日期
	Employment_form      string `json:"employment_form"`      //用工形式
	Personnel_source     string `json:"personnel_source"`     //人员来源
	Highest_education    string `json:"highest_education"`    //最高学历
	Graduation_school    string `json:"graduation_school"`    //毕业院校
	Major_studied        string `json:"major_studied"`        //所学专业
	Graduation_date      pgtype.Date `json:"graduation_date"`      //毕业日期
	Is_quit              string `json:"is_quit"`              //是否离职
}
