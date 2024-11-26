package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port int `mapstructure:"port"`
	*LogConfig `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"max_size"`
	MaxAge int `mapstructure:"max_age"`
	MaxBackups int `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host string `mapstructure:"host"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName string `mapstructure:"db_name"`
	Port int `mapstructure:"port"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port int `mapstructure:"port"`
	DB int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}

func Init()(err error){
	viper.SetConfigFile("./conf/config.yaml")
	// 或者
	// viper.SetConfigName() 搭配 viper.AddConfigPath()
	err = viper.ReadInConfig()

	if err != nil {
		fmt.Printf("viper ReadInConfig failed, err:%v\n", err)
		return err
	}
	// 把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper Unmarshal failed, err:%v\n", err)
	}
	// 支持配置热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper Unmarshal failed, err:%v\n", err)
		}
	})

	return err
}