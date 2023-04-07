package serve

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/components/ngrok"
	"github.com/fimreal/rack/pkg/config"
	"github.com/fimreal/rack/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.LoadConfigs()
}

// Run() 启动 web api 服务，传入 address 可以为端口 :8000
func Run() {
	// new gin engine with recovery()
	r := gin.New()
	r.Use(gin.Recovery(), gin.LoggerWithConfig(
		gin.LoggerConfig{
			SkipPaths: []string{"/favicon.ico", "/health", "/metrics", "/hc", "/"},
		},
	))

	// 装载路由
	r = loadRoutes(r)
	// 启动
	ezap.Fatal(serve(r))
}

func serve(r *gin.Engine) error {
	// ngrok
	if viper.GetBool("ngrok") {
		tun, err := ngrok.New()
		if err != nil {
			ezap.Fatal(err.Error())
		}
		ezap.Infof("ngrok tunnel created: %s", tun.URL())
		return r.RunListener(tun)
	}

	// listening on local addr
	port := ":" + viper.GetString("port")
	ezap.Infof("Listrning on %s", port)
	return r.Run(port)
}

func loadRoutes(r *gin.Engine) *gin.Engine {
	// only fileserver
	if viper.GetBool("fileserver") {
		service.RunFileserver(r)
	}

	healthcheck(r)
	disallowRobots(r)
	service.AddRoutes(r)

	return r
}
