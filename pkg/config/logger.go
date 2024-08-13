package config

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/spf13/viper"
)

func setLogger() {
	if viper.GetBool("debug") {
		ezap.SetLevel("debug")
	}
	ezap.Debug("logging debug  ✔")

	ezap.SetProjectName("[" + AppName + "]")
}
