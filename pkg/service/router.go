package service

import (
	"github.com/fimreal/rack/pkg/components/swagger"
	"github.com/fimreal/rack/pkg/service/fileserver"
	"github.com/fimreal/rack/pkg/service/ip2location"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoutes(r *gin.Engine) {
	pass := viper.GetBool("allservices")
	// only fileserver
	if viper.GetBool("fileserver") || pass {
		fileserver.LoadRoute(r)
		// return
	}
	if viper.GetBool("swagger") || pass {
		swagger.AddRoutes(r)
	}
	if viper.GetBool("ip2location") || pass {
		rg := r.Group("/ip2location")
		rg.GET("/:ip", ip2location.IpQuery)
		rg.POST("/:ip", ip2location.IpQuery)
	}

}
