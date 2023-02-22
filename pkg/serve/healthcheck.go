package serve

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 健康检查
// @Produce plain text
// @Param
// healthcheck
func Healthcheck(r *gin.RouterGroup) {
	r.GET("/hc", func(ctx *gin.Context) { ctx.String(http.StatusOK, "ok") })
	r.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })
}
