package main

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/serve"
)

func main() {
	ezap.Fatal(serve.Run())
}
