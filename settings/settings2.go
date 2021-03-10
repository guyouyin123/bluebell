//结构体配置文件方式
package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// conf全局变量
var Conf = new(AppConfig)

type AppConfig struct {
	Name       string `mapstructrue:"name"`
	Mode       string `mapstructrue:"mode"`
	Version    string `mapstructrue:"version"`
	Port       string `mapstructrue:"port"`
	*LogConfig `mapstructrue:"log"`
	*Mysql     `mapstructrue:"log"`
	*Redis     `mapstructrue:"log"`
}

type LogConfig struct {
	Level       string `mapstructrue:"level"`
	Filename    string `mapstructrue:"filename"`
	Max_size    int    `mapstructrue:"max_size"`
	Max_age     int    `mapstructrue:"max_age"`
	Max_backups int    `mapstructrue:"max_backups"`
}

type Mysql struct {
	Host           string `mapstructrue:"host"`
	Port           int    `mapstructrue:"port"`
	User           string `mapstructrue:"user"`
	Password       string `mapstructrue:"password"`
	Dbname         string `mapstructrue:"dbname"`
	Max_open_conns int    `mapstructrue:"max_open_conns"`
	Max_idle_conns int    `mapstructrue:"max_idle_conns"`
}

type Redis struct {
	Host      string `mapstructrue:"host"`
	Port      int    `mapstructrue:"port"`
	Db        int    `mapstructrue:"db"`
	Password  string `mapstructrue:"password"`
	Pool_size int    `mapstructrue:"pool_size"`
}

func Init2() (err error) {
	viper.SetConfigFile("config") // 指定配置文件
	viper.SetConfigType("yaml")   // 指定配置文件
	viper.AddConfigPath(".")      // 指定查找配置文件的路径
	err = viper.ReadInConfig()    // 读取配置信息
	if err != nil {               // 读取配置信息失败
		fmt.Println("读取配置信息失败:", err)
		return err
	}

	//把读取到到配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置问价修改了")
	})
	return err
}
