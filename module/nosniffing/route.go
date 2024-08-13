package nosniffing

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("all_services") {
		return
	}

	g.Use(gin.LoggerWithConfig(
		gin.LoggerConfig{
			SkipPaths: []string{"/favicon.ico", "/health", "/metrics"},
		},
	))
	g.GET("/robots.txt", disallowRobots)
	g.NoRoute(helloWorld)
}
