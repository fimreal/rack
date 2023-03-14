package serve

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/common"
	"github.com/fimreal/rack/pkg/config"
	"github.com/fimreal/rack/pkg/scripts"
	"github.com/fimreal/rack/pkg/service"
	"github.com/fimreal/rack/pkg/swagger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Run() 启动 web api 服务，传入 address 可以为端口 :8000
func Run() error {
	config.LoadConfigs()
	port := ":" + viper.GetString("port")
	if viper.GetBool("debug") {
		ezap.SetLevel("debug")
	} else {
		gin.SetMode(gin.ReleaseMode) // Switch to "release" mode in production.
	}

	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery(), gin.LoggerWithConfig(
		gin.LoggerConfig{
			SkipPaths: []string{"/health", "/metrics", "/hc", "/"},
		},
	))
	r.NoRoute(HelloWorld) // 404 => hello world

	apiroot := r.Group("/") // 装载路由
	// 健康检查
	Healthcheck(apiroot)
	// 基础服务
	common.AddRoutes(apiroot)
	// 特殊服务
	scripts.AddRoutes(apiroot)
	service.AddRoutes(apiroot)
	// swagger
	swagger.AddRoutes(apiroot)

	ezap.Infof("Listrning to %s", port)
	return r.Run(port)
}
