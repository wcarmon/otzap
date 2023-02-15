package otzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// BuildNormalZapCores returns a slice of zapcore.Core suitable for
// both local and GCloud based logging
// This is just an example, tweak to meet your needs
func BuildNormalZapCores(minLevel zapcore.Level) ([]zapcore.Core, error) {
	cores := make([]zapcore.Core, 0, 4)

	if IsInGoogleCloud() {
		// -- Google Cloud
		cores = append(cores, NewGoogleCloudCore(minLevel))

	} else {
		// -- Local
		cores = append(cores, NewPrettyConsoleCore(minLevel))
		cores = append(cores, NewRollingFileCore(minLevel))
	}

	otelCore := OTelZapCore{}
	err := otelCore.Validate()
	if err != nil {
		zap.L().Error("failed to validate OTelZapCore",
			zap.Error(err))

		return nil, err
	}

	cores = append(cores, otelCore)

	return cores, nil
}

// NewRollingFileCore shows how to configure a rolling file logger
// This is just an example, tweak to meet your needs
func NewRollingFileCore(
	minLevel zapcore.Level,
) zapcore.Core {

	cfg := zap.NewProductionEncoderConfig()
	cfg.MessageKey = "message"
	cfg.TimeKey = "time"

	rollingWriter := &lumberjack.Logger{
		Compress:   true,
		Filename:   "app.zap.log",
		MaxAge:     2, // in days
		MaxBackups: 2,
		MaxSize:    200, // in MB
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.AddSync(rollingWriter),
		minLevel)
}
