package common

import (
	"math/big"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-go/v9"
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
		req.URL.RawQuery = remote.RawQuery
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

// 默认返回 v4,否则检查是否为 v6
func IP2Dec(c *gin.Context) {
	ip := c.Param("ip")
	t := ip2location.OpenTools()

	if ipv4Dec, err := t.IPv4ToDecimal(ip); err == nil {
		c.String(http.StatusOK, ipv4Dec.String())
		return
	}
	if ipv6Dec, err := t.IPv6ToDecimal(ip); err == nil {
		c.String(http.StatusOK, ipv6Dec.String())
		return
	}

	c.String(http.StatusBadRequest, "Not a valid IPv4/IPv6 address.")
}

// 默认返回 v4,否则再按照 v6 返回
func Dec2IP(c *gin.Context) {
	ipDec := new(big.Int)
	ipDec, ok := ipDec.SetString(c.Param("ip"), 10)
	if !ok {
		c.String(http.StatusBadRequest, "Not a valid IPv4/IPv6 address.")
		return
	}
	t := ip2location.OpenTools()

	if ipv4, err := t.DecimalToIPv4(ipDec); err == nil {
		c.String(http.StatusOK, ipv4)
		return
	}
	if ipv6, err := t.DecimalToIPv6(ipDec); err == nil {
		c.String(http.StatusOK, ipv6)
		return
	}

	c.String(http.StatusBadRequest, "Not a valid IPv4/IPv6 decimal address.")
}

// 默认返回 v4,否则再按照 v6 返回
func CIDR2IP(c *gin.Context) {
	cidr := c.Param("ip") + "/" + c.Param("cidr")
	t := ip2location.OpenTools()
	if ipv4cidr, err := t.CIDRToIPv4(cidr); err == nil {
		c.String(http.StatusOK, strings.Join(ipv4cidr, "-"))
		return
	}
	if ipv6cidr, err := t.CIDRToIPv6(cidr); err == nil {
		c.String(http.StatusOK, strings.Join(ipv6cidr, "-"))
		return
	}
	c.String(http.StatusBadRequest, "Not a valid IPv4/IPv6 cidr. e.g. 192.168.0.0/16 ")
}

// 默认返回 v4,否则再按照 v6 返回
func IP2CIDR(c *gin.Context) {
	ipfrom := c.Param("ipfrom")
	ipto := c.Param("ipto")
	t := ip2location.OpenTools()
	if ipv4cidr, err := t.IPv4ToCIDR(ipfrom, ipto); err == nil {
		c.String(http.StatusOK, strings.Join(ipv4cidr, ","))
		return
	}
	if ipv6cidr, err := t.IPv6ToCIDR(ipfrom, ipto); err == nil {
		c.String(http.StatusOK, strings.Join(ipv6cidr, ","))
		return
	}
	c.String(http.StatusBadRequest, "Not a valid IPv4/IPv6 ip range. e.g. 0.0.0.0/255.255.255.255")
}

func ParseIPv6(c *gin.Context) {
	ip := c.Param("ip")
	t := ip2location.OpenTools()
	cip, err := t.CompressIPv6(ip)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	eip, _ := t.ExpandIPv6(ip)
	c.String(http.StatusOK, cip+" "+eip)
}

// 判断是否为私网 IP
func IsPrivateIP(c *gin.Context) {
	ip := net.ParseIP(c.Param("ip"))
	if ip != nil && ip.IsPrivate() {
		c.String(http.StatusOK, "true")
		return
	}
	c.String(http.StatusOK, "false")
}

// 判断是否为保留网络地址
// https://en.wikipedia.org/wiki/Reserved_IP_addresses
func IsReservedIP(c *gin.Context) {
	ip := net.ParseIP(c.Param("ip"))
	revNet := []string{
		"0.0.0.0/8",
		"10.0.0.0/8",
		"100.64.0.0/10",
		"127.0.0.0/8",
		"169.254.0.0/16",
		"172.16.0.0/12",
		"192.0.0.0/24",
		"192.0.2.0/24",
		"192.88.99.0/24",
		"192.168.0.0/16",
		"198.18.0.0/15",
		"198.51.100.0/24",
		"203.0.113.0/24",
		"224.0.0.0/4",
		"240.0.0.0/4",
		"255.255.255.255/32",
	}

	for _, cidr := range revNet {
		_, ipNet, _ := net.ParseCIDR(cidr)
		if ipNet.Contains(ip) {
			c.String(http.StatusOK, "true")
			return
		}
	}
	c.String(http.StatusOK, "false")
}
