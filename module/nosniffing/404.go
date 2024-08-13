package nosniffing

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	NOROUTEMSG = "404 page not found\n"
	ROBOTSTXT  = "User-agent: *\nDisallow: /"
)

// @Summary     robots.txt 配置
// @Description deny spider bot
// @Produce     plain
// @Success     200 {string} string "User-agent: *\nDisallow: /"
// @Router      /robots.txt [get]
func disallowRobots(c *gin.Context) {
	c.String(http.StatusOK, ROBOTSTXT)
}

func helloWorld(c *gin.Context) {
	c.String(http.StatusNotFound, NOROUTEMSG)
}
