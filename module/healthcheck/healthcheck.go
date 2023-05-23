package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary     健康检查
// @Description healthcheck
// @Produce     plain
// @Success     200 {string} string "ok"
// @Router      /health [get]
// @Router      /hc [get]
func healthcheck(r *gin.Engine) {
	r.Any("/hc", func(ctx *gin.Context) { ctx.String(http.StatusOK, "ok") })
	r.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })
}
