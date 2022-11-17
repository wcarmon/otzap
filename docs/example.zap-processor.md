# Overview
1. Example [`otzap.ZapSpanProcessor`](https://github.com/wcarmon/otzap/blob/main/otel_span_processor.go#L39)


# Example Zap [Processor](https://opentelemetry.io/docs/collector/configuration/#processors) for [OpenTelemetry](https://opentelemetry.io/)
```go
func NewZapSpanEventProcessor() (*otzap.ZapSpanProcessor, error) {
	p := &otzap.ZapSpanProcessor{
		TimestampKey: "time",
	}

	if err := p.Validate(); err != nil {
		zap.L().Error("failed to validate ZapSpanProcessor",
			zap.Error(err))

		return nil, err
	}

	return p, nil
}
```
