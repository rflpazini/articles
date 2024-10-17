package observability

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitTracerProvider() *trace.TracerProvider {
	client := otlptracehttp.NewClient()
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		log.Fatalf("Error creating OTLP exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewSchemaless(
			attribute.String("service.name", "observability-api"),
			attribute.String("environment", "development"),
			attribute.String("version", "1.0.0"),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp
}

func ShutdownTracerProvider(tp *trace.TracerProvider) error {
	return tp.Shutdown(context.Background())
}