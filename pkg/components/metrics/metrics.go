package metrics

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func ServeMetrics() {
	if !viper.GetBool("metrics") && !viper.GetBool("all_services") {
		return
	}

	m := gin.New()
	m.GET("/status", func(c *gin.Context) {
		uptime := time.Since(startTime)
		c.JSON(200, gin.H{
			"app_name":       config.AppName,
			"version":        config.Version,
			"build_time":     config.BuildTime,
			"build_modules":  config.GetModVer(),
			"gin_version":    gin.Version,
			"gin_mode":       gin.Mode(),
			"go_arch":        runtime.GOARCH,
			"go_version":     runtime.Version(),
			"uptime_seconds": fmt.Sprintf("%.f", uptime.Seconds()),
			"uptime":         uptime.Round(time.Second).String(),
		})
	})
	m.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// gin metrics
	useGinMetrics(m)

	port := ":" + viper.GetString("metrics_port")
	ezap.Info("metrics server started on " + port)
	ezap.Error(m.Run(port))

}
