package service

import (
	"github.com/fimreal/rack/pkg/components/swagger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoutes(r *gin.Engine) {
	pass := viper.GetBool("allservices")

	if viper.GetBool("swagger") || pass {
		swagger.AddRoutes(r)
	}

}
