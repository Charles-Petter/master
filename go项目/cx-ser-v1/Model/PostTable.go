package Model

/*岗位表*/
type PostTable struct {
	Post_number        int    `json:"post_number"`        //岗位编号
	Post_name          string `json:"post_name"`          //岗位名称
	Department_number  int    `json:"department_number"`  //部门编号
	Post_type          string `json:"post_type"`          //岗位类型
	Post_establishment string `json:"post_establishment"` //岗位编制
}
