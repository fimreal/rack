package ipquery

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
	r.GET("/help/ipquery", help)

	// client ip
	r.GET("/ip", ClientIP)

	// reserved ip info service
	r.GET("/ipinfo", ClientIPInfo)
	r.GET("/ipinfo/:ip", GetIPInfo)

	// ip tools
	r.GET("/ip2dec/:ip", IP2Dec)
	r.GET("/dec2ip/:ip", Dec2IP)
	r.GET("/cidr2ip/:ip/:cidr", CIDR2IP)
	r.GET("/ip2cidr/:ipfrom/:ipto", IP2CIDR)
	r.GET("/pipv6/:ip", ParseIPv6)
	r.GET("/pip/:ip", IsPrivateIP)
	r.GET("/rip/:ip", IsReservedIP)
}
