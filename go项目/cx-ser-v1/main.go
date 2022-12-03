package main

import (
	"cx/Global"
	"cx/Mapper"
	"cx/Router"
	"cx/logger"
	"cx/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//GO Web 开发通用的脚手架模板
func main() {
	//1.加载配置
	v := settings.InitConfig()
	fmt.Println("配置文件读取结果：", v.Get("server.port"))
	//2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger  init success... ")
	//3.初始化Mysql 连接
	fmt.Println("主函数")
	var err error
	Global.Db, err = Mapper.InitDB(v)

	re := gin.Default()
	if err != nil {
		fmt.Println("数据库连接失败！")
		return
	}

	//4.注册路由
	//re := routers.Setup()
	re = Router.CreateRoute(re)// 创建路由
	// 监听并在 :8080 上启动服务
	panic(re.Run(":" + v.GetString("server.port")))//若报端口被占用 重启电脑


	//5.启动服务（优雅关机）
	//re.GET("/", func(c *gin.Context) {
	//	time.Sleep(5 * time.Second)
	//	c.String(http.StatusOK, "Welcome Gin Server")
	//})
	//srv := &http.Server{
	//	Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
	//	Handler: re,
	//}
	//go func() {
	//	// 开启一个goroutine启动服务
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()
	//// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	//quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	//// kill 默认会发送 syscall.SIGTERM 信号
	//// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	//// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	//// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	//<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	//zap.L().Info("Shutdown Server ...")
	//// 创建一个5秒超时的context
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	//if err := srv.Shutdown(ctx); err != nil {
	//	zap.L().Fatal("Server Shutdown: ")
	//}
	//zap.L().Info("Server exiting")


}

