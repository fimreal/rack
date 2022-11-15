package cap

import (
	"net/http"

	"github.com/fimreal/rack/pkg/cap/common"
	"github.com/fimreal/rack/pkg/cap/ncmd"
	"github.com/fimreal/rack/pkg/cap/sssub"
	"github.com/gin-gonic/gin"
)

func AddCapRoutes(r *gin.RouterGroup) {

	capBasePath := "/tools"
	cap := r.Group(capBasePath)

	// 一些内置命令
	cap.GET("/ifip", ncmd.IfIP)
	cap.GET("/ip", ncmd.GetIP)
	cap.GET("/ipinfo", ncmd.GetIPInfo)
	cap.GET("/hostname", common.Hostname)

	// 特殊请求
	cap.POST("/shell", common.ShellExec)
	cap.GET("/sssub", sssub.SssubToUrl)
}

// healthcheck
func AddHealthcheck(r *gin.RouterGroup) {
	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(http.StatusOK, "ok") })
	r.POST("/healthz", func(ctx *gin.Context) { ctx.String(http.StatusOK, "ok") })
}
