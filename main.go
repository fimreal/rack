package main

import (
	"os"

	"github.com/fimreal/go-tools/serve"
)

func main() {
	// 启动端口配置为第一个参数
	var port string
	if len(os.Args[:]) > 1 {
		port = os.Args[1]
	} else {
		port = "3333"
	}
	serve.Serve(port)
}
