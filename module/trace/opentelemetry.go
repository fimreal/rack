// https://opentelemetry.io/blog/2022/go-web-app-instrumentation/
package trace

import (
	"context"

	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/config"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

// initTracer 初始化 OpenTelemetry
func initTracer() (*trace.TracerProvider, error) {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
	otel.SetTracerProvider(tp)
	return tp, nil
}

func LoadOtel(g *gin.Engine) {
	// 初始化 OpenTelemetry
	tp, err := initTracer()
	if err != nil {
		ezap.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			ezap.Fatal(err)
		}
	}()

	// 使用 OpenTelemetry 中间件
	g.Use(otelgin.Middleware(config.AppName))
}
