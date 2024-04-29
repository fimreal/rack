package config

import (
	"strings"

	"github.com/spf13/viper"
)

func LoadConfigs() {
	// shell 不允许带'.'的环境变量，识别环境变量时去除'.'
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, ``))
	viper.AutomaticEnv() // 加载环境变量

	// 配置日志格式
	setLogger()

}
