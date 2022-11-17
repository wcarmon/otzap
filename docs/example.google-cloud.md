# Overview
1. Examples for [Google Cloud Trace](https://cloud.google.com/trace) with [OpenTelemetry](https://opentelemetry.io/) (in go)


# Example [Google Cloud Exporter](https://github.com/GoogleCloudPlatform/opentelemetry-operations-go)
```go
import (
    texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
    "go.uber.org/zap"
    ...
)

func NewGoogleCloudExporter(defaultProjectId string) (*texporter.Exporter, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	if strings.TrimSpace(projectID) == "" && otzap.IsInGoogleCloud() {
		projectID = defaultProjectId

		zap.L().Warn("GOOGLE_CLOUD_PROJECT is missing, defaulting",
			zap.String("projectID", projectID))
	}

	return texporter.New(texporter.WithProjectID(projectID))
}
```
