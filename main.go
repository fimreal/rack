package main

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/config"
	"github.com/fimreal/rack/pkg/serve"
	"github.com/ory/viper"
)

func main() {
	config.LoadConfigs()
	port := ":" + viper.GetString("port")
	ezap.Infof("Listrning to %s", port)
	ezap.Fatal(serve.Run(port))
}
