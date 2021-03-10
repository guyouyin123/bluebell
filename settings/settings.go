// yaml配置文件方式

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("config.yaml") // 指定配置文件，如果是json就写json
	//viper.SetConfigName("config")  //远程指定配置文件名称
	//viper.SetConfigType("yaml")   // 远程指定配置类型
	viper.AddConfigPath(".")   // 指定查找配置文件的路径
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		fmt.Println("读取配置信息失败:", err)
		return err
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置问价修改了")
	})
	return err
}
