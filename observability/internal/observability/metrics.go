package observability

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func InitMeterProvider() {
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("Prometheus exporter: %v", err)
	}

	meterProvider := metric.NewMeterProvider(metric.WithReader(exporter))
	otel.SetMeterProvider(meterProvider)
}

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}

func MetricsHandler() echo.HandlerFunc {
	return echo.WrapHandler(promhttp.Handler())
}
