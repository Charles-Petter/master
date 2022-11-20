package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"m/Global"
	"m/Mapper"
	"m/Router"
)

func main() {
	fmt.Println("主函数")
	v := Mapper.InitConfig()
	fmt.Println("配置文件读取结果：", v.Get("server.port"))

	var err error
	Global.Db, err = Mapper.InitDB(v)

	re := gin.Default()
	if err != nil {
		fmt.Println("数据库连接失败！")
		return
	}
	re = Router.CreateRoute(re)
	panic(re.Run(":" + v.GetString("server.port")))
}
