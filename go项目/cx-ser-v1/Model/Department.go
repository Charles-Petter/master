package Model

import "github.com/jackc/pgtype"

/*部门表*/
type Department struct {
	Department_number   int    `json:"department_number"`   //部门编号
	Department_name     string `json:"department_name"`     //部门名称
	Department_type     string `json:"department_type"`     //部门类型
	Department_head     string    `json:"department_head"`     //部门主管员工编号
	Phone               string `json:"phone"`               //部门电话
	Fax                 string `json:"fax"`                 //传真
	Describe            string `json:"describe"`            //描述
	Superior_department string `json:"superior_department"` //上级部门
	Incorporation_date  pgtype.Date `json:"incorporation_date"`  //成立日期
}
