package trace

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("all_services") {
		return
	}

	// 加载 opentelemetry
	LoadOtel(g)
}
