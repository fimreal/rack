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
func healthcheck(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.OPTIONS("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}
