package scripts

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoutes(r *gin.RouterGroup) {
	if !viper.GetBool("scripts") {
		return
	}

	serviceBasePath := "/i"
	i := r.Group(serviceBasePath)
	i.Use(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.RequestURI, "/i/") {
			ctx.Header("Cache-Control", "max-age=86400")
			// 浏览器可读
			ctx.Header("Content-Type", "text/plain; charset=utf-8")
		}
	})

	shellScripts, _ := fs.Sub(FSsh, "shell")
	i.StaticFS("/", http.FS(shellScripts))

}
