package Mapper

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"

	"cx/Model"

	"gorm.io/gorm"
)
//连接数据库
func InitDB_cx(v_cx *viper.Viper) (db_cx *gorm.DB, erre_cx error) {
	var err_cx error
	args_cx := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		v_cx.GetString("mysql.user_cx"),
		v_cx.GetString("mysql.password_cx"),
		v_cx.GetString("mysql.host_cx"),
		v_cx.GetInt("mysql.port_cx"),
		v_cx.GetString("mysql.dbname_cx"),
	)
	db_cx, err_cx = gorm.Open(mysql.Open(args_cx), &gorm.Config{})
	if err_cx != nil {
		panic("连接数据库失败" + err_cx.Error())
		return db_cx, erre_cx
	}
	fmt.Println("数据库连接成功！")
	err_cx = db_cx.AutoMigrate(&Model.Employee_cx{})
	if err_cx != nil {
		return nil, err_cx
	}
	return db_cx, erre_cx
}


