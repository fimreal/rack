package randomgenerator

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("allservices") {
		return
	}
	r := g.Group(RoutePrefix)

	// list route
	r.GET("/help/random", help)

	//
	r.GET("/code", SixNumber)
	r.GET("/genpass", GenRandomPassword)
}
