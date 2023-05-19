package scripts

import (
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	serviceBasePath := "/i"
	i := r.Group(serviceBasePath)

	// i.StaticFS("/", http.FS(staticFile))
	// i.GET("/*filepath", fileHandler())

	staticFile, _ := fs.Sub(FSstatic, "static")

	i.Use(func(ctx *gin.Context) {
		filename := strings.TrimPrefix(ctx.Request.RequestURI, "/i/")
		if filename != "" {
			_, err := fs.Stat(staticFile, filename)
			if err != nil {
				if os.IsNotExist(err) {
					// do request github pages
					reqGHPage(ctx)
					ctx.Done()
				}
				ezap.Error(err.Error())
				ctx.String(http.StatusInternalServerError, err.Error())
				ctx.Done()
			}
			// 如果不是 HTML 请求，则添加文本 MIME 类型和缓存头
			// if !strings.HasSuffix(ctx.Request.RequestURI, ".html") {
			if !strings.HasSuffix(filename, ".html") {
				ctx.Header("Cache-Control", "max-age=86400")
				ctx.Header("Content-Type", "text/plain; charset=utf-8")
			}
		}
	})

	i.StaticFS("/", http.FS(staticFile))
}

// 找不到 embed 文件时，请求 https://s.epurs.com
func reqGHPage(c *gin.Context) {
	remote, _ := url.Parse("http://s.epurs.com")
	filename := strings.TrimPrefix(c.Request.RequestURI, "/i/")
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.Header = c.Request.Header
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path + "/" + filename
		req.URL.RawQuery = ""
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
