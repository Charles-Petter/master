package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)
//viper 管理文件
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
