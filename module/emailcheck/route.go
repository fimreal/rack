package emailcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("allservices") {
		return
	}

	g.GET("/remail/:email", IsRealMail)
	g.POST("/remail/:email", IsRealMail)
}
