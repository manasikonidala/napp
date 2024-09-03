package emptyexporter

import (
    "context"
    "go.opentelemetry.io/collector/component"
    "go.opentelemetry.io/collector/exporter"
    "go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
    typeStr = "emptyexporter"
)

func NewFactory() exporter.Factory {
    return exporter.NewFactory(
        typeStr,
        createDefaultConfig,
        exporter.WithTraces(createTracesExporter, component.StabilityLevelDevelopment),
        exporter.WithMetrics(createMetricsExporter, component.StabilityLevelDevelopment),
        exporter.WithLogs(createLogsExporter, component.StabilityLevelDevelopment),
    )
}

func createTracesExporter(
    ctx context.Context,
    set exporter.CreateSettings,
    config component.Config) (exporter.Traces, error) {

    cfg := config.(*Config)
    s := NewEmptyexporter()
    return exporterhelper.NewTracesExporter(ctx, set, cfg, s.pushTraces)
}

func createMetricsExporter(
    ctx context.Context,
    set exporter.CreateSettings,
    config component.Config) (exporter.Metrics, error) {

    cfg := config.(*Config)
    s := NewEmptyexporter()
    return exporterhelper.NewMetricsExporter(ctx, set, cfg, s.pushMetrics)
}

func createLogsExporter(
    ctx context.Context,
    set exporter.CreateSettings,
    config component.Config) (exporter.Logs, error) {

    cfg := config.(*Config)
    s := NewEmptyexporter()
    return exporterhelper.NewLogsExporter(ctx, set, cfg, s.pushLogs)
}

type Config struct {
}

func createDefaultConfig() component.Config {
    return &Config{}
}
