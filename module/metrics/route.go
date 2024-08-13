package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("all_services") {
		return
	}

	r := g.Group(RoutePrefix)

	r.GET("/status", ShowStatus)
}
