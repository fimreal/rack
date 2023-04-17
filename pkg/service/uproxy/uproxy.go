package uproxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	// ASSET_URL := "https://hunshcn.github.io/gh-proxy"
	url := c.Param("url")
	ParseReq(c.Request)
	// if strings.HasPrefix(url,"https://")
	c.String(200, url)
}

func ParseReq(req *http.Request) {

}
