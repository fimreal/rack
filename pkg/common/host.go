package common

import (
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

func Hostname(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		c.String(500, hostname)
		ezap.Error(err)
		return
	}
	c.String(http.StatusOK, hostname)
}

func HostIP(c *gin.Context) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		c.String(500, "")
		ezap.Error(err)
		return
	}
	var ips []string
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	c.String(http.StatusOK, strings.Join(ips, " "))
}
