package dnsquery

import (
	"bytes"
	"io"
	"net"
	"net/http"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

// golang 建立 tcp 连接，指定发送数据内容
func Whois(c *gin.Context) {
	var (
		domain       = c.Param("domain")
		result       = bytes.NewBuffer(nil)
		WhoissServer = []string{"whois.internic.net:43", "whois.arin.net:43", "whois.godaddy.com:43", "whois.porkbun.com:43"}
		conn         net.Conn
		err          error
	)
	// 从列表中随机获取一个 tcp 连接(range 不按顺序)
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
