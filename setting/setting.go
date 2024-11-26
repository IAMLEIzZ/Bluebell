package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init()(err error){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Printf("viper ReadInConfig failed, err:%v\n", err)
		return err
	}
	// 支持配置热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改...")
	})

	return err
}