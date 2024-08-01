package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings") // 设置配置文件名称
	viper.SetConfigType("yml")      // 设置配置文件后缀
	viper.AddConfigPath("./conf/")  // 设置配置文件路径
	err := viper.ReadInConfig()     // 将配置文件读取到内存中

	if err != nil {
		panic(fmt.Sprintf("Load Config Error: %s", err.Error()))
	}

	fmt.Println(viper.GetString("server.port"))
}
