package ncmd

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/fimreal/go-tools/cmd"
	"github.com/gin-gonic/gin"
)

func IfIP(context *gin.Context) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	var ips []string
	res := cmd.CommandResult{}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	res = cmd.CommandResult{Command: "getIP", ReturnCode: 0, Result: ips}
	context.JSON(http.StatusOK, res)
}

func GetIP(context *gin.Context) {
	res := cmd.CommandResult{}
	resp, err := http.Get("http://epurs.com/ip")
	if err != nil {
		res = cmd.CommandResult{Command: "/serverip", ReturnCode: 1, Result: err}
		context.JSON(http.StatusOK, res)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	var content string
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		content += string(buf[:n])
	}
	res = cmd.CommandResult{Command: "/serverip", ReturnCode: 0, Result: strings.Replace(content, "\n", "", -1)}
	context.JSON(http.StatusOK, res)
}

func GetIPInfo(context *gin.Context) {
	res := cmd.CommandResult{}
	resp, err := http.Get("http://epurs.com/ipinfo")
	if err != nil {
		res = cmd.CommandResult{Command: "/serverip", ReturnCode: 1, Result: err}
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	var content string
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		content += string(buf[:n])
	}
	res = cmd.CommandResult{Command: "/serverip", ReturnCode: 0, Result: content}
	context.JSON(http.StatusOK, res)
}
