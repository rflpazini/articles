package observability

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	m "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

var (
	meterProvider      *metric.MeterProvider
	requestCounter     m.Int64Counter
	activeRequests     m.Int64UpDownCounter
	requestHistogram   m.Float64Histogram
	customRequestCount m.Int64Counter
)

func InitMeterProvider() {
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("Prometheus exporter: %v", err)
	}

	meterProvider = metric.NewMeterProvider(
		metric.WithReader(exporter),
		metric.WithResource(resource.Default()),
	)

	otel.SetMeterProvider(meterProvider)

	meter := meterProvider.Meter("observability-handler")

	requestCounter, err = meter.Int64Counter(
		"http_requests_total",
		m.WithDescription("Total number of HTTP requests processed"),
	)
	if err != nil {
		log.Fatalf("Failed to create request counter: %v", err)
	}

	activeRequests, err = meter.Int64UpDownCounter(
		"http_active_requests",
		m.WithDescription("Number of active HTTP requests being processed"),
	)
	if err != nil {
		log.Fatalf("Failed to create active request gauge: %v", err)
	}

	requestHistogram, err = meter.Float64Histogram(
		"http_request_duration_seconds",
		m.WithDescription("The distribution of request durations in seconds"),
	)
	if err != nil {
		log.Fatalf("Failed to create request duration histogram: %v", err)
	}

	customRequestCount, err = meter.Int64Counter(
		"custom_request_count",
		m.WithDescription("Custom request count for specific endpoints"),
	)
	if err != nil {
		log.Fatalf("Failed to create request counter: %v", err)
	}
}

func RecordRequestMetrics(ctx context.Context, duration time.Duration) {
	requestCounter.Add(ctx, 1)

	activeRequests.Add(ctx, 1)
	defer activeRequests.Add(ctx, -1)

	requestHistogram.Record(ctx, duration.Seconds())
}

func RecordCustomRequestMetrics(ctx context.Context) {
	customRequestCount.Add(ctx, 1)
}
