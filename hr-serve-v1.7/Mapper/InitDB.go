package Mapper

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"m/Model"
	"os"
)

func InitDB(v *viper.Viper) (db *gorm.DB, erre error) {
	var err error
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		v.GetString("mysql.user"),
		v.GetString("mysql.password"),
		v.GetString("mysql.host"),
		v.GetInt("mysql.port"),
		v.GetString("mysql.dbname"),
	)
	//db, err = gorm.Open(postgres.Open(args), &gorm.Config{})
	db, err = gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败" + err.Error())
		return db, erre
	}
	fmt.Println("数据库连接成功！")
	err = db.AutoMigrate(&Model.Employee{})
	if err != nil {
		return nil, err
	}
	return db, erre
}

func InitConfig() (v *viper.Viper) {
	workDir, _ := os.Getwd()
	v = viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yml")
	v.AddConfigPath(workDir + "/Config")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("读取配置文件出错：%s\n", err)
	}
	return v
}
