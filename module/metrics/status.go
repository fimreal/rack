package metrics

import (
	"github.com/fimreal/rack/pkg/config"
	"github.com/gin-gonic/gin"
)

func ShowStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"app_name":   config.AppName,
		"version":    config.Version,
		"build_time": config.BuildTime,
		"modules":    config.GetModVer(),
	})
}
