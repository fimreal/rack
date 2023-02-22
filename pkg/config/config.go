package config

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/service/ip2location"
	"github.com/spf13/viper"
)

func LoadConfigs() {
	ezap.SetLogTime("")
	// LoadConfigFile()  // 加载配置文件
	viper.AutomaticEnv() // 加载环境变量
	setFlag()            // 解析传入参数

	if viper.GetBool("ip2location") {
		ip2location.GetDB()
	}
}
