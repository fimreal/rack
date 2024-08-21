package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func useGinMetrics(g *gin.Engine) {
	monitor := ginmetrics.GetMonitor()
	monitor.SetMetricPath("/gin-metrics")
	monitor.Use(g)
}
