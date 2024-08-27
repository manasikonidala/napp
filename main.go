package main

import (
    "context"
    "log"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
    "google.golang.org/grpc"
)

func main() {
    ctx := context.Background()

    // Set up the OTLP exporter
    exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(),
        otlptracegrpc.WithEndpoint("localhost:4317"))
    if err != nil {
        log.Fatalf("failed to create exporter: %v", err)
    }

    // Set up the Tracer provider
    tp := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
        trace.WithResource(resource.Default()),
    )
    defer func() { _ = tp.Shutdown(ctx) }()

    otel.SetTracerProvider(tp)

    // Your application logic here
    tracer := otel.Tracer("example-tracer")
    _, span := tracer.Start(ctx, "example-operation")
    log.Println("Hello, OpenTelemetry!")
    span.End()
}
