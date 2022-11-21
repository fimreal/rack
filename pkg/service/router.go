package service

import (
	"github.com/fimreal/rack/pkg/service/aliyun"
	"github.com/fimreal/rack/pkg/service/email"
	"github.com/fimreal/rack/pkg/service/ip2location"
	"github.com/gin-gonic/gin"
	"github.com/ory/viper"
)

func ServiceRoutes(r *gin.RouterGroup) {
	serviceBasePath := "/s"
	srv := r.Group(serviceBasePath)

	// 阿里云
	if viper.GetBool("aliyun") {
		srv.POST("/addsgrule", aliyun.Allow)
	}

	if viper.GetBool("mail") {
		srv.POST("/mailto", email.SendMail) // 兼容 http2mail，注意修改路径
	}

	if viper.GetBool("ip2location") {
		srv.GET("/ip/:ip", ip2location.IpQuery)
	}

}
