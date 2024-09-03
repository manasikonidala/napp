package emptyexporter

import (
    "context"
    "encoding/json"
    "fmt"
    "go.opentelemetry.io/collector/pdata/plog"
    "go.opentelemetry.io/collector/pdata/pmetric"
    "go.opentelemetry.io/collector/pdata/ptrace"
)

type emptyexporter struct {
}

func NewEmptyexporter() *emptyexporter {
    return &emptyexporter{}
}

type LogData struct {
    Timestamp        string                 `json:"Timestamp"`
    ObservedTimestamp string                `json:"ObservedTimestamp"`
    TraceId          string                 `json:"TraceId"`
    SpanId           string                 `json:"SpanId"`
    SeverityText     string                 `json:"SeverityText"`
    SeverityNumber   int                    `json:"SeverityNumber"`
    Body             string                 `json:"Body"`
    Resource         map[string]string      `json:"Resource"`
    Attributes       map[string]interface{} `json:"Attributes"`
}

func (s *emptyexporter) pushLogs(_ context.Context, ld plog.Logs) error {
    logs := ld.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords()

    for i := 0; i < logs.Len(); i++ {
        logRecord := logs.At(i)

        logData := LogData{
            Timestamp:        logRecord.Timestamp().AsTime().String(),
            ObservedTimestamp: logRecord.ObservedTimestamp().AsTime().String(),
            TraceId:          logRecord.TraceID().HexString(),
            SpanId:           logRecord.SpanID().HexString(),
            SeverityText:     logRecord.SeverityText(),
            SeverityNumber:   int(logRecord.SeverityNumber()),
            Body:             logRecord.Body().AsString(),
            Resource:         convertResourceMap(logRecord.Resource().Attributes()),
            Attributes:       convertAttributesMap(logRecord.Attributes()),
        }

        logJSON, err := json.Marshal(logData)
        if err != nil {
            return fmt.Errorf("failed to marshal log data: %w", err)
        }

        fmt.Println(string(logJSON))  // Replace with actual log handling
    }
    return nil
}

func (s *emptyexporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
    // Handle metrics here
    return nil
}

func (s *emptyexporter) pushTraces(_ context.Context, td ptrace.Traces) error {
    // Handle traces here
    return nil
}

func convertResourceMap(resourceAttrs map[string]interface{}) map[string]string {
    resourceMap := make(map[string]string)
    for key, val := range resourceAttrs {
        if strVal, ok := val.(string); ok {
            resourceMap[key] = strVal
        }
    }
    return resourceMap
}

func convertAttributesMap(attrs map[string]interface{}) map[string]interface{} {
    attributesMap := make(map[string]interface{})
    for key, val := range attrs {
        attributesMap[key] = val
    }
    return attributesMap
}
