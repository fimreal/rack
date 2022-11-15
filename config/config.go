package config

import (
	"github.com/ory/viper"
	"github.com/spf13/pflag"
)

var Debug = false

func LoadConfigs() {
	// ⬇️ 注意配置加载顺序，看起来不可变的优先级最低，但不代表文件修改不会自动重载 ⬇️
	// 加载配置文件
	LoadConfigFile()
	// 加载环境变量
	BindEnvFor()
	// 解析传入参数
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// debug all thing
	Debug = true
}
