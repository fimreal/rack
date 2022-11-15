package config

import (
	"os"

	"github.com/fimreal/goutils/ezap"
	"github.com/fsnotify/fsnotify"
	"github.com/ory/viper"
)

func LoadConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath(viper.GetString(os.Getenv("HOME") + "/.rack/"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 如果是因为找不到文件，则忽略该错误
			ezap.Warn(err)
		} else {
			ezap.Fatalf("Loading config file failed: %v\n", err)
		}
	}
	// 监听文件修改，热加载配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		ezap.Warnf("Config file changed: %s, %s", in.Name, in.Op)
	})
}
