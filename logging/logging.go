package logging

import (
    "context"
    "encoding/json"
    "log"
    "go.opentelemetry.io/otel/trace"
)

// LogData represents the structure of the log data model.
type LogData struct {
    Timestamp        string `json:"Timestamp"`
    TraceId          string `json:"TraceId"`
    SpanId           string `json:"SpanId"`
    SeverityText     string `json:"SeverityText"`
    Body             string `json:"Body"`
}

// LogEvent logs an event with trace and span IDs.
func LogEvent(ctx context.Context, severityText, message string) {
    span := trace.SpanFromContext(ctx)
    logData := LogData{
        Timestamp:    span.EndTime().String(),
        TraceId:      span.SpanContext().TraceID().String(),
        SpanId:       span.SpanContext().SpanID().String(),
        SeverityText: severityText,
        Body:         message,
    }

    logJSON, err := json.Marshal(logData)
    if err != nil {
        log.Fatalf("failed to marshal log data: %v", err)
    }

    log.Println(string(logJSON))
}
