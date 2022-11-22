package Model

import "github.com/jackc/pgtype"

/*岗位调动员工表*/
type PostTransferEmployeeTable struct {
	Transfer_id          int         `json:"transfer_id"`          //岗位调动编号
	Id                   string      `json:"id"`                   //员工编号
	Name                 string      `json:"name"`                 //姓名
	Sex                  string      `json:"sex"`                  //性别
	Original_post_number int         `json:"original_post_number"` //原岗位编号
	Original_post_name   string      `json:"original_post_name"`   //原岗位名称
	New_post_number      int         `json:"new_post_number"`      //新岗位编号
	New_post_name        string      `json:"new_post_name"`        //新岗位名称
	Transfer_date        pgtype.Date `json:"transfer_date"`        //调动日期
	Reasons_for_transfer string      `json:"reasons_for_transfer"` //调动原因
}
