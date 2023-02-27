package common

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// return client ip
// @Summary     Return client ip
// @Description Return client ip
// @Produce     plain
// @Success     200 {string} 1.1.1.1
// @Router      /ip [get]
func ClientIP(c *gin.Context) {
	c.String(http.StatusOK, c.ClientIP())
}

// 需要访问互联网，使用 ip2location 更好
// @Summary     Describe client ip
// @Description Describe client ip，需要访问互联网，使用 ip2location 更好
// @Produce     json
// @Success     200 {object} ipinfo
// @Router      /ipinfo [get]
func ClientIPInfo(c *gin.Context) {
	remote, _ := url.Parse("http://ip-api.com/json/?lang=zh-CN")
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.Header = c.Request.Header
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path + "/" + c.ClientIP()
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
