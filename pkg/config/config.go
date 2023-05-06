package config

import (
	"strings"

	"github.com/fimreal/rack/pkg/service/ip2location"
	"github.com/spf13/viper"
)

func LoadConfigs() {
	// 加载配置文件。太麻烦了，先不加
	// LoadConfigFile()

	// shell 不允许带'.'的环境变量，识别环境变量时去除'.'
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, ``))
	viper.AutomaticEnv() // 加载环境变量

	// setFlag()            // 解析传入参数

	// 配置日志格式
	setLogger()

	if viper.GetBool("ip2location") {
		ip2location.GetDB()
	}
}
