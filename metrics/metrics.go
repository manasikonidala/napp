package metrics

import (
    "context"
    "log"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/metric"
    "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
    sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

// InitializeMeter sets up the global meter.
func InitializeMeter() metric.Meter {
    exporter, err := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
    if err != nil {
        log.Fatalf("failed to initialize stdoutmetric exporter %v", err)
    }

    provider := sdkmetric.NewMeterProvider(
        sdkmetric.WithReader(exporter),
    )

    otel.SetMeterProvider(provider)
    return otel.Meter("otel-go-example")
}

// RecordMetric records a simple metric.
func RecordMetric(ctx context.Context, meter metric.Meter, name string, value int64) {
    counter := metric.Must(meter).NewInt64Counter(name)
    counter.Add(ctx, value)
}
