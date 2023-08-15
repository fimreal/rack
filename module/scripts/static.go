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

// 如果在 embed 中找到文件，直接返回结果
// 如果没有找到，则执行 reqGHPage(c)，请求 github 最新页面
func inFile(c *gin.Context) {
	filename := strings.TrimPrefix(c.Request.RequestURI, "/i/")
	if filename != "" {
		_, err := fs.Stat(StaticFiles, filename)
		if err != nil {
			if os.IsNotExist(err) {
				// do request github pages
				reqGHPage(c)
				c.Abort()
				return
			}
			ezap.Error(err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		// 如果不是 HTML 请求，则添加文本 MIME 类型和缓存头
		// if !strings.HasSuffix(c.Request.RequestURI, ".html") {
		if !strings.HasSuffix(filename, ".html") {
			c.Header("Cache-Control", "max-age=86400")
			c.Header("Content-Type", "text/plain; charset=utf-8")
		}
	}
}

// 找不到 embed 文件时，请求 https://s.epurs.com
func reqGHPage(c *gin.Context) {
	remote, _ := url.Parse(ScriptsHub)
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
