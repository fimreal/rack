package dnsquery

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
	r.GET("/help/dnsquery", help)

	// whois
	r.GET("/whois/:domain", Whois)

	// dns query
	r.GET("/dns/:host", DNSQueryLocal)
	r.GET("/dnscf/:host", DNSQuery)
}
