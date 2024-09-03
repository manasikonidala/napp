package tracing

import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/trace"
    "context"
    "log"
)

// InitializeTracer sets up the global tracer.
func InitializeTracer() trace.Tracer {
    exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
    if err != nil {
        log.Fatalf("failed to initialize stdouttrace exporter %v", err)
    }

    provider := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
    )

    otel.SetTracerProvider(provider)
    return otel.Tracer("otel-go-example")
}

// StartTrace starts a new trace span.
func StartTrace(ctx context.Context, tracer trace.Tracer, spanName string) (context.Context, trace.Span) {
    return tracer.Start(ctx, spanName)
}
