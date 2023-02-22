package common

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

// 需要访问互联网，使用 ip2location 更好
func GetIPInfo(c *gin.Context) {
	remote, _ := url.Parse("http://ip-api.com/json/?lang=zh-CN")
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.Header = c.Request.Header
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path + "/" + c.Param("ip")
		req.URL.RawQuery = remote.RawQuery
	}
	ezap.Debugf("%v+", proxy.Director)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func GetDNSRecord(c *gin.Context) {
	host := c.Param("host")
	cname, _ := net.LookupCNAME(host)
	if cname != host+"." {
		c.String(http.StatusOK, cname)
		return
	}
	iprecords, err := net.LookupIP(c.Param("host"))
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
