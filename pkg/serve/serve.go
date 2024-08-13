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
	config.ShowInfo()
	if !viper.GetBool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	// new gin engine with recovery()
	g := gin.New()

	// gin 中间件配置
	allowedOrigins = viper.GetStringSlice("cors_allowed_origins")
	g.Use(gin.Recovery(), Cors(), gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/favicon.ico", "/health", "/metrics"}}))

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
	ezap.Infof("Listening on %s", port)
	return r.Run(port)
}
