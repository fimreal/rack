package fileserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
)

func webdavHandler(httpRoot string) gin.HandlerFunc {
	handler := &webdav.Handler{
		FileSystem: webdav.Dir(httpRoot),
		LockSystem: webdav.NewMemLS(), // 在内存中创建锁系统，避免并发写入的冲突
	}
	return gin.WrapH(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
}
