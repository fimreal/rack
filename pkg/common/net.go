package common

import (
	"bytes"
	"io"
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

func Whois(c *gin.Context) {
	var (
		domain       = c.Param("domain")
		result       = bytes.NewBuffer(nil)
		WhoissServer = []string{"whois.internic.net:43", "whois.arin.net:43", "whois.godaddy.com:43", "whois.porkbun.com:43"}
		conn         net.Conn
		err          error
	)
	// 从列表中随机获取一个 tcp 连接
	for i, server := range WhoissServer {
		conn, err = net.Dial("tcp", server)
		if err == nil {
			ezap.Debug("request whois " + domain + " to " + server)
			break
		}
		ezap.Error(err.Error())
		if i != len(WhoissServer)-1 {
			continue
		}
		c.String(http.StatusInternalServerError, "client to whois server occurs error")
		return
	}
	defer conn.Close()

	// http://www.faqs.org/rfcs/rfc812.html
	_, err = conn.Write([]byte(domain + " \r\n"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		result.Write(buf[:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			ezap.Error(err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.String(http.StatusOK, result.String())
}
