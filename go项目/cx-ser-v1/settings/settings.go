package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

////viper 管理文件
//
//func Init() (err error) {
//	//viper.SetConfigFile("./config.yaml") // 指定配置文件路径
//	viper.SetConfigName("config") //指定配置文件的名称（不需要带后缀）
//	viper.SetConfigType("yaml")   //指定配置文件的类型
//	viper.AddConfigPath(".")
//	err = viper.ReadInConfig() // 读取配置信息
//	if err != nil {
//		// 读取配置信息失败
//		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
//		return
//	}
//
//	// 监控配置文件变化
//	viper.WatchConfig()
//	viper.OnConfigChange(func(in fsnotify.Event) {
//		fmt.Println("配置文件修改了...")
//	})
//	return
//}

//var Conf = new(AppConfig)
//
//type AppConfig struct {
//	Mode         string `mapstructure:"mode"`
//	Port         int    `mapstructure:"port"`
//	Name         string `mapstructure:"name"`
//	Version      string `mapstructure:"version"`
//	StartTime    string `mapstructure:"start_time"`
//	MachineID    int    `mapstructure:"machine_id"`
//	*LogConfig   `mapstructure:"log"`
//	*MySQLConfig `mapstructure:"mysql"`
//	*RedisConfig `mapstructure:"redis"`
//}
//
//type MySQLConfig struct {
//	Host         string `mapstructure:"host"`
//	User         string `mapstructure:"user"`
//	Password     string `mapstructure:"password"`
//	DB           string `mapstructure:"dbname"`
//	Port         int    `mapstructure:"port"`
//	MaxOpenConns int    `mapstructure:"max_open_conns"`
//	MaxIdleConns int    `mapstructure:"max_idle_conns"`
//}
//
//type RedisConfig struct {
//	Host         string `mapstructure:"host"`
//	Password     string `mapstructure:"password"`
//	Port         int    `mapstructure:"port"`
//	DB           int    `mapstructure:"db"`
//	PoolSize     int    `mapstructure:"pool_size"`
//	MinIdleConns int    `mapstructure:"min_idle_conns"`
//}
//
//type LogConfig struct {
//	Level      string `mapstructure:"level"`
//	Filename   string `mapstructure:"filename"`
//	MaxSize    int    `mapstructure:"max_size"`
//	MaxAge     int    `mapstructure:"max_age"`
//	MaxBackups int    `mapstructure:"max_backups"`
//}


func InitConfig() (v *viper.Viper){
	workDir, _ := os.Getwd()
	v = viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yml")
	v.AddConfigPath(workDir + "/Config")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("读取配置文件出错：%s\n", err)
	}
	//监控配置文件变化
		v.WatchConfig()
		v.OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("配置文件修改了...")
		})
	return v
}
