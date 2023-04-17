package uproxy

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// go:embed assets
var assetsFiles embed.FS

func LoadRoute(r *gin.RouterGroup) {
	r.StaticFS("/assets/", http.FS(assetsFiles))
	r.Any("/", handler)
}
