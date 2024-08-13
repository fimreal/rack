package metrics

import (
	"github.com/fimreal/rack/pkg/config"
	"github.com/gin-gonic/gin"
)

func ShowStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"appname": config.AppName,
		"version": config.Version,
		"modules": config.GetModVer(),
	})
}
