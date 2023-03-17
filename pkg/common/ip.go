package common

import (
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-go/v9"
)

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
