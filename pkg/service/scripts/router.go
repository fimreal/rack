package scripts

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	serviceBasePath := "/i"
	i := r.Group(serviceBasePath)
	i.Use(func(ctx *gin.Context) {
		if !strings.HasSuffix(ctx.Request.RequestURI, ".html") {
			ctx.Header("Cache-Control", "max-age=86400")
			// 浏览器可读
			ctx.Header("Content-Type", "text/plain; charset=utf-8")
		}
	})

	staticFile, _ := fs.Sub(FSstatic, "static")
	i.StaticFS("/", http.FS(staticFile))

}
