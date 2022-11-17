# Overview
1. Examples for [Jaeger](https://www.jaegertracing.io/) with [OpenTelemetry](https://opentelemetry.io/) (in go)


# Example [Jaeger Exporter](https://pkg.go.dev/go.opentelemetry.io/otel/exporters/jaeger)
```go
import (
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.uber.org/zap"
    ...
)

func NewJaegerExporter(cfg *conf.AppConf) (*jaeger.Exporter, error) {

	url := cfg.Tracing.JaegerExporter.Url

	exporter, err := jaeger.New(
		jaeger.WithCollectorEndpoint(
			jaeger.WithEndpoint(url)))

	if err != nil {
		zap.L().Error("failed to create Jaeger exporter",
			zap.Error(err),
			zap.String("serviceName", cfg.Tracing.ServiceName),
			zap.String("url", url),
		)

		return nil, err
	}

	return exporter, nil
}
```