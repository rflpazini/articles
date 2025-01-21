package observability

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracer() *sdktrace.TracerProvider {
	client := otlptracehttp.NewClient()

	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		log.Fatalf("Error creating OTLP exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			"",
			attribute.String("service.name", "observability-handler"),
			attribute.String("environment", "development"),
			attribute.String("version", "1.0.0"),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp
}

func ShutdownTracerProvider(tp *sdktrace.TracerProvider) error {
	return tp.Shutdown(context.Background())
}
