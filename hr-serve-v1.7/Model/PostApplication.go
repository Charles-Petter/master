package Model

/*岗位申请表*/
type PostApplications struct {
	Application_id           int    `json:"application_id"`           //申请编号
	Id                       string `json:"id"`                       //员工编号
	Name                     string `json:"name"`                     //姓名
	Original_department_name string `json:"original_department_name"` //原部门编号
	New_department_name      string `json:"new_department_name"`      //新部门编号
	Original_post_name       string `json:"original_post_name"`       //原岗位编号
	New_post_name            string `json:"new_post_name"`            //新岗位编号
	Is_agree                 int    `json:"is_agree"`                 //领导是否同意
	Reason_for_application   string `json:"reason_for_application"`   //申请原因
}