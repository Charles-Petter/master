package Model
import _ "github.com/jackc/pgtype"


/*薪水表*/
type Salaytable_cx struct {
	Id                  string `json:"id"`                   //员工编号
	Employee_type        string    `json:"employee_type"`        //员工类型
	Name                 string `json:"name"`                 //姓名
	Sex                 string `json:"sex"`                  //性别
	Initsalay            string `json:"initsalay"`             //税前工资
	Taxafter          string `json:"taxafter"`              //实际工资
}
