package nosniffing

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("all_services") {
		return
	}

	g.Use(checkBannedIPs)

	// 封禁 ip 配置
	r := g.Group(RoutePrefix)
	if viper.GetBool(ID + "_banip") {
		ezap.Info("启用恶意检查及 IP 封禁功能")
		g.Use(maliciousRequestChecker)
		go cleanupBannedIPs()

		keyPath := "/" + viper.GetString(ID+"_key")
		if keyPath != "/" {
			keyPath = keyPath + "/"
		}
		// 查看当前被禁止的 IP 列表
		r.GET(keyPath+"denyip", listBannedIPs)
		// 增加禁止的 IP
		r.GET(keyPath+"denyip/:ip", addBannedIPs)
		// 解除封禁某个 IP
		r.GET(keyPath+"allowip/:ip", unbanIP)
	}

	g.GET("/robots.txt", disallowRobots)
	g.NoRoute(noRouteMsg)
}
