package Mapper

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"cx/Model"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
//连接数据库
func InitDB(v_cx *viper.Viper) (db *gorm.DB, erre_cx error) {
	var err_cx error
	args_cx := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		v_cx.GetString("mysql.user"),
		v_cx.GetString("mysql.password"),
		v_cx.GetString("mysql.host"),
		v_cx.GetInt("mysql.port"),
		v_cx.GetString("mysql.dbname"),
	)
	db, err_cx = gorm.Open(mysql.Open(args_cx), &gorm.Config{})
	if err_cx != nil {
		panic("连接数据库失败" + err_cx.Error())
		return db, erre_cx
	}
	fmt.Println("数据库连接成功！")
	err_cx = db.AutoMigrate(&Model.Employee_cx{})
	if err_cx != nil {
		return nil, err_cx
	}
	return db, erre_cx
}


