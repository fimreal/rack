package config

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	AppName    = "rack"
	Version    = ""
	BuildTime  = ""
	ModVersion []string
)

func AddModVer(mv string) {
	ModVersion = append(ModVersion, mv)
}

func GetModVer() []string {
	return ModVersion
}

func LoadConfigs() {
	// shell 不允许带'.'的环境变量，识别环境变量时修改为_
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	// FIXME: viper v1.17.0（支持） 之后会找不到环境变量，需要一个个手动绑定
	viper.AutomaticEnv() // 加载环境变量

	// 配置日志格式
	setLogger()
}
