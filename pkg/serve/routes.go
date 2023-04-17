package serve

import (
	"net/http"

	"github.com/fimreal/rack/pkg/service"
	"github.com/gin-gonic/gin"
)

func loadRoutes(r *gin.Engine) {
	service.AddRoutes(r)
	healthcheck(r)
	disallowRobots(r)
	r.NoRoute(helloWorld) // 404 => hello world
}

// @Summary     robots.txt 配置
// @Description deny spider bot
// @Produce     plain
// @Success     200 {string} string "User-agent: *\nDisallow: /"
// @Router      /robots.txt [get]
func disallowRobots(r *gin.Engine) {
	r.GET("/robots.txt", func(ctx *gin.Context) { ctx.String(http.StatusOK, "User-agent: *\nDisallow: /") })
}

// @Summary     健康检查
// @Description healthcheck
// @Produce     plain
// @Success     200 {string} string "ok"
// @Router      /health [get]
// @Router      /hc [get]
func healthcheck(r *gin.Engine) {
	r.HEAD("/", func(ctx *gin.Context) { ctx.Status(200) })
	r.GET("/hc", func(ctx *gin.Context) { ctx.String(http.StatusOK, "ok") })
	r.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })
}

func helloWorld(c *gin.Context) {
	c.String(http.StatusNotFound, "Hey, world!\n")
}
