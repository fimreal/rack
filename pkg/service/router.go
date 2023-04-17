package service

import (
	"github.com/fimreal/rack/pkg/components/swagger"
	"github.com/fimreal/rack/pkg/service/aliyun"
	"github.com/fimreal/rack/pkg/service/chatgpt"
	"github.com/fimreal/rack/pkg/service/common"
	"github.com/fimreal/rack/pkg/service/dockerhub"
	"github.com/fimreal/rack/pkg/service/email"
	"github.com/fimreal/rack/pkg/service/fileserver"
	"github.com/fimreal/rack/pkg/service/ip2location"
	"github.com/fimreal/rack/pkg/service/scripts"
	"github.com/fimreal/rack/pkg/service/uproxy"
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
	if viper.GetBool("common") || pass {
		common.AddRoutes(r)
	}
	if viper.GetBool("scripts") || pass {
		scripts.AddRoutes(r)
	}
	if viper.GetBool("swagger") || pass {
		swagger.AddRoutes(r)
	}
	if viper.GetBool("docker") || pass {
		r.GET("/docker.io/:namespace/:repository/*result", dockerhub.ListTags)
	}

	if viper.GetBool("aliyun") || pass {
		r.POST("/addsgrule", aliyun.Allow)
	}

	if viper.GetBool("mail") || pass {
		r.POST("/mailto", email.SendMail) // 兼容 http2mail，注意修改路径
	}

	if viper.GetBool("ip2location") || pass {
		rg := r.Group("/ip2location")
		rg.GET("/:ip", ip2location.IpQuery)
		rg.POST("/:ip", ip2location.IpQuery)
	}

	if viper.GetBool("chatgpt") || pass {
		rg := r.Group("/chatgpt")
		rg.GET("/", chatgpt.Ask)
		rg.POST("/", chatgpt.Ask)
	}

	if viper.GetBool("uproxy") || pass {
		rg := r.Group("/uproxy")
		uproxy.LoadRoute(rg)
	}
}
