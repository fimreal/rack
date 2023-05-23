package config

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/spf13/viper"
)

func setLogger() {
	ezap.SetLogTime("")

	if viper.GetBool("debug") {
		ezap.SetLevel("debug")
	}
	ezap.Debug("logging debug  ✔")
}
