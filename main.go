package main

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/config"
	"github.com/fimreal/rack/pkg/serve"
)

func main() {
	config.LoadConfigs()
	ezap.Fatal(serve.Run(":8000"))
}
