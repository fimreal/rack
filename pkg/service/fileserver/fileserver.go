package fileserver

import (
	"path/filepath"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func LoadRoute(r *gin.Engine) {
	httpRoot := viper.GetString("workdir")
	if absPath, err := filepath.Abs(httpRoot); err == nil {
		httpRoot = absPath
	}
	ezap.Info("启动文件服务器, 网站根目录: " + httpRoot)
	// r.Static("/download", httpRoot)
	// r.Static("/file", httpRoot)
	r.Any("/dav", webdavHandler(httpRoot))
	r.PUT("/upload", upload)
}
