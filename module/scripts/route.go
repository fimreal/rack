package scripts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("allservices") {
		return
	}
	r := g.Group(RoutePrefix)

	// list route
	r.GET("/help/"+ID, help)

	// 如果在 embed 中找到文件，直接返回结果
	// 如果没有找到，则执行 reqGHPage(c)，请求 github 最新页面
	g.Use(inFile)
	g.StaticFS("/", http.FS(StaticFiles))
}
