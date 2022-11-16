package serve

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/cap"
	"github.com/fimreal/rack/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/ory/viper"
)

// Run() 启动 web api 服务，传入 address 可以为端口 :8000
func Run(address string) error {
	if viper.GetBool("debug") {
		ezap.SetLevel("debug")
	} else {
		gin.SetMode(gin.ReleaseMode) // Switch to "release" mode in production.
	}

	r := gin.Default()
	apiroot := r.Group("/")

	cap.Healthcheck(apiroot)
	cap.CapRoutes(apiroot)
	service.ServiceRoutes(apiroot)

	return r.Run(address)
}
