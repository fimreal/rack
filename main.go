package main

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/serve"
)

// @title    rack
// @version  0.1.1
func main() {
	ezap.Fatal(serve.Run())
}
