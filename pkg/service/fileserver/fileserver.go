package fileserver

import (
	"net/http"
	"strings"

	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func upload(c *gin.Context) {
	saveDir := c.PostForm("saveDir")
	filename := c.PostForm("filename")

	httpRoot := viper.GetString("workdir")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: "+err.Error())
		return
	}

	// 从 param 中配置存储文件信息
	if filename == "" {
		filename = fileHeader.Filename
	}
	if saveDir == "" {
		saveDir = httpRoot + "/"
	} else if strings.Contains(saveDir, "..") {
		// 不应该包含上级目录, 创建文件位置不可控
		ezap.Error("存储文件目录格式不合法: ", saveDir)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"result": "the saveDir " + saveDir + " is invalid",
		})
	} else {
		saveDir = httpRoot + "/" + saveDir + "/"
	}
	dst := saveDir + filename

	// 检查目录是否存在
	err = utils.MakeDir(saveDir)
	if err != nil {
		ezap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"result": err.Error(),
		})
	}

	ezap.Infof("开始接收文件, 存储到 %s", dst)

	err = c.SaveUploadedFile(fileHeader, dst)
	if err != nil {
		ezap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"result": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": filename + " uploaded.",
	})
}

func LoadRoute(r *gin.Engine) {
	httpRoot := viper.GetString("workdir")
	ezap.Info("启动文件服务器, 网站根目录: " + httpRoot)
	r.Static("/", httpRoot)
	r.POST("/upload", upload)
	r.PUT("/upload", upload)
}
