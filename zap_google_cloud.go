package otzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// IsInGoogleCloud returns true iff the application is running in Google Cloud
func IsInGoogleCloud() bool {
	return os.Getenv("HOME") == "/app"
}

// NewGoogleCloudCore builds a zapcore.Core that writes to stdout
// in Google Cloud Logging format
//
// See https://pkg.go.dev/go.uber.org/zap/zapcore#Core
// See https://cloud.google.com/logging
func NewGoogleCloudCore(minLevel zapcore.Level) zapcore.Core {

	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeLevel = GoogleCloudLevelEncoder

	// See https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)
	cfg.LevelKey = "severity"
	cfg.MessageKey = "message"
	cfg.TimeKey = "timestamp"

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.Lock(os.Stdout),
		minLevel)
}

// GoogleCloudLevelEncoder encodes a zapcore.Level to a Google Cloud Logging severity
//
// See https://pkg.go.dev/go.uber.org/zap/zapcore#LevelEncoder
// See https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
func GoogleCloudLevelEncoder(
	lvl zapcore.Level,
	enc zapcore.PrimitiveArrayEncoder,
) {
	s := "DEFAULT"
	switch lvl {
	case zapcore.DebugLevel:
		s = "DEBUG"
	case zapcore.InfoLevel:
		s = "INFO"
	case zapcore.WarnLevel:
		s = "WARNING"
	case zapcore.ErrorLevel:
		s = "ERROR"
	case zapcore.FatalLevel:
		s = "ALERT"
	}

	enc.AppendString(s)
}
