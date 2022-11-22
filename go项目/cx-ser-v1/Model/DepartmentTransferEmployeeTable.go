package Model

import "github.com/jackc/pgtype"

/*部门调动员工表*/
type DepartmentTransferEmployeeTable struct {
	Transfer_id                int         `json:"transfer_id"`                //部门调动编号
	Id                         string      `json:"id"`                         //员工编号
	Name                       string      `json:"name"`                       //姓名
	Sex                        string      `json:"sex"`                        //性别
	Original_department_number int         `json:"original_department_number"` //原部门编号
	Original_department_name   string      `json:"original_department_name"`   //原部门名称
	New_department_number      int         `json:"new_department_number"`      //新部门编号
	New_department_name        string      `json:"new_department_name"`        //新部门名称
	Transfer_date              pgtype.Date `json:"transfer_date"`              //调动日期
	Reasons_for_transfer       string      `json:"reasons_for_transfer"`       //调动原因
}
