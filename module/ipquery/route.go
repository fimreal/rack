package ipquery

import "github.com/gin-gonic/gin"

func AddRoute(g *gin.Engine) {
	r := g.Group(RoutePrefix)
	r.GET("/myip", ClientIP)
}

func ClientIP(c *gin.Context) {
	c.String(200, c.ClientIP())
}
