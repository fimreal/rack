package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoutes(r *gin.RouterGroup) {
	if !viper.GetBool("common") {
		return
	}

	// client 相关
	r.GET("/ip", ClientIP)
	r.GET("/ipinfo", ClientIPInfo)

	// host 相关
	r.GET("/hostip", HostIP)
	r.GET("/hostname", Hostname)

	// 对外请求
	r.GET("/ipinfo/:ip", GetIPInfo)
	r.GET("/dns/:host", GetDNSRecord)
	r.GET("/whois/:domain", Whois)

	// 小函数
	r.GET("/code", SixNumber)
	r.GET("/genpass", GenRandomPassword)
	r.GET("/time", TimeStamp)
	r.GET("/time/:ts", TimeStampTrans)
}
