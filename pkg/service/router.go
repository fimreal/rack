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
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoutes(r *gin.Engine) {
	if viper.GetBool("common") {
		common.AddRoutes(r)
	}
	if viper.GetBool("scripts") {
		scripts.AddRoutes(r)
	}
	if viper.GetBool("swagger") {
		swagger.AddRoutes(r)
	}
	if viper.GetBool("docker") {
		r.GET("/docker.io/:namespace/:repository/*result", dockerhub.ListTags)
	}

	serviceBasePath := "/s"
	srv := r.Group(serviceBasePath)

	if viper.GetBool("aliyun") {
		srv.POST("/addsgrule", aliyun.Allow)
	}

	if viper.GetBool("mail") {
		srv.POST("/mailto", email.SendMail) // 兼容 http2mail，注意修改路径
	}

	if viper.GetBool("ip2location") {
		srv.GET("/ip/:ip", ip2location.IpQuery)
		srv.POST("/ip/:ip", ip2location.IpQuery)
	}

	if viper.GetBool("chatgpt") {
		srv.GET("/chatgpt", chatgpt.Ask)
		srv.POST("/chatgpt", chatgpt.Ask)
	}

}

func RunFileserver(r *gin.Engine) {
	fileserver.LoadRoute(r)
}
