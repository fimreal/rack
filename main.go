package main

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/serve"
)

// @title    rack
// @version  0.3.0
func main() {
	ezap.Fatal(serve.Run())
}
