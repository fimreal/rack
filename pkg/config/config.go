package config

import (
	"os"
	"strings"

	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/service/ip2location"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func LoadConfigs() {
	ezap.SetLogTime("")
	// LoadConfigFile()  // 加载配置文件
	// shell 不允许带'.'的环境变量，配置绑定环境变量不含'.'
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, ``))
	viper.AutomaticEnv() // 加载环境变量
	setFlag()            // 解析传入参数

	if viper.GetBool("version") {
		ezap.Println("Rack", VERSION)
		os.Exit(0)
	}

	if viper.GetBool("debug") {
		ezap.SetLevel("debug")
	} else {
		gin.SetMode(gin.ReleaseMode) // Switch to "release" mode in production.
	}

	if viper.GetBool("ip2location") {
		ip2location.GetDB()
	}
}
