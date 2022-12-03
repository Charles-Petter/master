package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)
//viper 管理文件
func InitConfig() (v_cx *viper.Viper){
	workDir_cx, _ := os.Getwd()
	v_cx = viper.New()
	v_cx.SetConfigName("application")
	v_cx.SetConfigType("yml")
	v_cx.AddConfigPath(workDir_cx + "/Config")
	err_cx := v_cx.ReadInConfig()
	if err_cx != nil {
		fmt.Printf("读取配置文件出错：%s\n", err_cx)
	}
	//监控配置文件变化
		v_cx.WatchConfig()
		v_cx.OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("配置文件修改了...")
		})
	return v_cx
}
