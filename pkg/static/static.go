package static

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup) {

	serviceBasePath := "/h"
	i := r.Group(serviceBasePath)
	// i.Use(func(ctx *gin.Context) {
	// 	if strings.HasPrefix(ctx.Request.RequestURI, "/h/") {
	// 		ctx.Header("Cache-Control", "max-age=86400")
	// 		ctx.Header("Content-Type", "text/html; charset=utf-8")
	// 	}
	// })

	html, _ := fs.Sub(FShtml, "html")
	i.StaticFS("/", http.FS(html))

}
