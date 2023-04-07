package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	r.NoRoute(helloWorld) // 404 => hello world

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

	// ip tools
	r.GET("/ip2dec/:ip", IP2Dec)
	r.GET("/dec2ip/:ip", Dec2IP)
	r.GET("/cidr2ip/:ip/:cidr", CIDR2IP)
	r.GET("/ip2cidr/:ipfrom/:ipto", IP2CIDR)
	r.GET("/pipv6/:ip", ParseIPv6)
	r.GET("/pip/:ip", IsPrivateIP)
	r.GET("/rip/:ip", IsReservedIP)

	// 小函数
	r.GET("/code", SixNumber)
	r.GET("/genpass", GenRandomPassword)
	r.GET("/time", TimeStamp)
	r.GET("/time/:ts", TimeStampTrans)

	r.GET("/help", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, `/ip
/ipinfo
/hostip
/hostname
/ipinfo/:ip
/dns/:host
/whois/:domain
/ip2dec/:ip
/dec2ip/:ip
/cidr2ip/:ip/:cidr
/ip2cidr/:ipfrom/:ipto
/pipv6/:ip
/pip/:ip
/rip/:ip
/code
/genpass
/time
/time/:ts
`)
	})
}
