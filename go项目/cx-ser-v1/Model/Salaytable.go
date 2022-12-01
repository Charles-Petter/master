package Model
import _ "github.com/jackc/pgtype"


/*薪水表*/
type Salaytable struct {
	Id                  string `json:"id"`                   //员工编号
	Employee_type_cx        string    `json:"employee_type_cx"`        //员工类型
	Name_cx                 string `json:"name_cx"`                 //姓名
	Sex_cx                  string `json:"sex_cx"`                  //性别
	Initsalay_cx            string `json:"Initsalay_cx"`             //税前工资
	SalayCount_cx           string `json:"SalayCount_cx"`              //税后工资
}
