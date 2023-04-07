package serve

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary     robots.txt 配置
// @Description deny spider bot
// @Produce     plain
// @Success     200 {string} string "User-agent: *\nDisallow: /"
// @Router      /robots.txt [get]
func disallowRobots(r *gin.Engine) {
	r.GET("/robots.txt", func(ctx *gin.Context) { ctx.String(http.StatusOK, "User-agent: *\nDisallow: /") })
}
