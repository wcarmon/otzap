package otzap

import (
	"go.opentelemetry.io/otel/attribute"
	"os"
	"runtime"
	"strings"
)

// CollectResourceAttributes reads Environment Variables
// and returns corresponding OpenTelemetry Attributes
func CollectResourceAttributes() []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0, 64)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		if MustNotLogEnvVar(pair[0]) {
			continue
		}

		attrs = append(attrs, attribute.String(pair[0], pair[1]))
	}

	attrs = append(attrs,
		attribute.Int("cpuCount", runtime.NumCPU()),
		attribute.String("goVersion", runtime.Version()),
	)

	return attrs
}

// MustNotLogEnvVar is a predicate
// MustNotLogEnvVar returns true for sensitive env vars
// MustNotLogEnvVar relies on 2 global vars:
// - BlockedEnvVars
// - BlockedEnvVarSubstrings
func MustNotLogEnvVar(varName string) bool {

	for _, value := range BlockedEnvVars {
		if strings.EqualFold(varName, value) {
			return true
		}
	}

	for _, needle := range BlockedEnvVarSubstrings {
		if strings.Contains(strings.ToLower(varName), strings.ToLower(needle)) {
			return true
		}
	}

	return false
}
