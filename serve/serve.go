package serve

import (
	"log"

	"github.com/fimreal/go-tools/cmd/common"
	"github.com/fimreal/go-tools/cmd/ncmd"
	"github.com/fimreal/go-tools/cmd/sssub"
	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	// Switch to "release" mode in production.
	gin.SetMode(gin.ReleaseMode)

	// 创建一个 gin实例,返回一个 *engine 路由引擎
	r := gin.Default()

	// 返回健康信息
	r.GET("/healthz", common.Healthcheck)
	r.POST("/healthz", common.Healthcheck)

	// 一些内置命令
	r.GET("/ifip", ncmd.IfIP)
	r.GET("/ip", ncmd.GetIP)
	r.GET("/ipinfo", ncmd.GetIPInfo)
	r.GET("/hostname", common.Hostname)

	// 特殊请求
	r.POST("/shell", common.ShellExec)
	r.GET("/ss", sssub.SssubToUrl)

	// 启动 gin 服务
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err.Error())
	}

}
