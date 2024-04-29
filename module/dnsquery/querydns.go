package dnsquery

import (
	"net"
	"net/http"
	"strings"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

func DNSQuery(c *gin.Context) {
	host := c.Param("host")
	answer, err := CloudFlareDoh(host)
	if err != nil {
		ezap.Error(err.Error())
		c.AbortWithStatusJSON(500, answer)
		return
	}
	c.JSON(200, answer)
}

// golang 解析 dns
func DNSQueryLocal(c *gin.Context) {
	host := c.Param("host")
	cname, _ := net.LookupCNAME(host)
	if cname != host+"." {
		c.String(http.StatusOK, cname)
		return
	}
	iprecords, err := net.LookupIP(host)
	if err != nil {
		ezap.Error(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
	var ips []string
	for _, ip := range iprecords {
		ips = append(ips, ip.String())
	}
	c.String(http.StatusOK, strings.Join(ips, " "))
}
