package serve

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/module"
	"github.com/fimreal/rack/pkg/components/ngrok"
	"github.com/fimreal/rack/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	config.LoadConfigs()
	if !viper.GetBool("debug") {
		gin.SetMode(gin.ReleaseMode) // Default mode is debug, please switch to "release" mode in production.
	}
	// new gin engine with recovery()
	g := gin.New()
	g.Use(gin.Recovery(), gin.LoggerWithConfig(
		gin.LoggerConfig{
			SkipPaths: []string{"/favicon.ico", "/health", "/metrics", "/hc", "/"},
		},
	))
	// service.AddRoutes(g)
	module.GinLoad(g)
	ezap.Fatal(serve(g))
}

// ngrok or local
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
