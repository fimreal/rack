package website

import (
	"embed"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Embed static files
var (
	//go:embed static
	staticFS embed.FS
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("all_services") {
		return
	}

	// Serve static files
	g.Use(static.Serve(RoutePrefix, static.EmbedFolder(staticFS, "static")))

	ezap.Infof("Serving static files at '%s'", RoutePrefix)
}
