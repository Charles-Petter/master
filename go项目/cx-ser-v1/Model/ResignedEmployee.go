package Model

import (
	"github.com/jackc/pgtype"
)

/*离职员工表*/
type ResignedEmployee struct {
	Id                      string      `json:"id"`
	Name                    string      `json:"name"`
	Sex                     string      `json:"sex"`
	Resignation_date        pgtype.Date `json:"resignation_date"`
	Reasons_for_resignation string      `json:"reasons_for_resignation"`
	Department_number       int         `json:"department_number"`
	Department_name         string      `json:"department_name"`
	Post_number             int         `json:"post_number"`
	Post_name               string      `json:"post_name"`
}
