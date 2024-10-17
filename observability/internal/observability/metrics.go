package observability

import (
	"log"

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
