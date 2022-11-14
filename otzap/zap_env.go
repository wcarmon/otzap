package otzap

import (
	"go.uber.org/zap"
	"os"
	"strings"
)

// PrintEnvVars prints local environment variables using zap API
func PrintEnvVars() {
	fields := make([]zap.Field, 0, len(os.Environ()))
	fields = append(fields, zap.Int("count", len(os.Environ())))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fields = append(fields, zap.String(pair[0], pair[1]))
	}

	zap.L().Info("env vars", fields...)
}
