package Model

import "github.com/jackc/pgtype"

/*人才库表*/
type TalentPool struct {
	Talent_id      int      `json:"talent_id"`
	Name           string      `json:"name"`
	Sex            string      `json:"sex"`
	Birthday       pgtype.Date `json:"birthday"`
	Id_card        string      `json:"id_card"`
	Political      string      `json:"political"`
	Nation         string      `json:"nation"`
	Native_place   string      `json:"native_place"`
	Phone          string      `json:"phone"`
	Email          string      `json:"email"`
	Height         int         `json:"height"`
	Blood_type     string      `json:"blood_type"`
	Marital_status string      `json:"marital_status"`
	Birthplace           string `json:"birthplace"`
	Registered_residence string `json:"registered_residence"`
	Department_name      string `json:"department_name"`
	Post_name            string `json:"post_name"`
	Employment_form      string `json:"employment_form"`
	Personnel_source     string `json:"personnel_source"`
	Highest_education    string `json:"highest_education"`
	Graduation_school    string `json:"graduation_school"`
	Major_studied        string `json:"major_studied"`
	Graduation_date      pgtype.Date `json:"graduation_date"`
	Is_agree             int    `json:"is_agree"`
}
