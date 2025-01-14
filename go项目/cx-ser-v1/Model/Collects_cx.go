package Model

import "github.com/jackc/pgtype"

/*学院表*/
type Collects_cx struct {
	Department_number   int    `json:"department_number"`   //学院编号
	Department_name     string `json:"department_name"`     //学院名称
	Name                 string `json:"name"`
	Phone               string `json:"phone"`               //学院电话
	Description            string `json:"description"`            //学院描述
	Incorporation_date  pgtype.Date `json:"incorporation_date"`  //学院成立日期
}
