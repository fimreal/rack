package cap

import (
	"net/http"

	"github.com/fimreal/rack/pkg/cap/common"
	"github.com/fimreal/rack/pkg/cap/ncmd"
	"github.com/fimreal/rack/pkg/cap/randomstring"
	"github.com/fimreal/rack/pkg/cap/sssub"
	"github.com/gin-gonic/gin"
	"github.com/ory/viper"
)

func CapRoutes(r *gin.RouterGroup) {
	if !viper.GetBool("tools") {
		return
	}
	capBasePath := "/t"
	cap := r.Group(capBasePath)

	// 一些内置命令
	cap.GET("/ifip", ncmd.IfIP)
	cap.GET("/ip", ncmd.GetIP)
	cap.GET("/ipinfo", ncmd.GetIPInfo)
	cap.GET("/hostname", common.Hostname)
	cap.POST("/random", randomstring.GenRandomString)

	// 特殊请求
	cap.POST("/shell", common.ShellExec)
	cap.GET("/sssub", sssub.SssubToUrl)
}

// healthcheck
func Healthcheck(r *gin.RouterGroup) {
	r.GET("/hc", func(ctx *gin.Context) { ctx.String(http.StatusOK, "ok") })
	r.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })
}
