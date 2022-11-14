package otzap

import (
	"go.uber.org/zap"
	"strings"
)

func (zp ZapSpanProcessor) GetEventSourceKey() string {
	clean := strings.TrimSpace(zp.EventSourceKey)
	if clean != "" {
		return clean
	}

	return defaultLogEventSourceKey
}

func (zp ZapSpanProcessor) GetEventSourceValue() string {
	clean := strings.TrimSpace(zp.EventSourceValue)
	if clean != "" {
		return clean
	}

	return defaultOtelSourceValue
}

func (zp ZapSpanProcessor) GetLogger() *zap.Logger {
	if zp.Logger != nil {
		return zp.Logger
	}

	return zap.L()
}

func (zp ZapSpanProcessor) GetSpanIdKey() string {
	clean := strings.TrimSpace(zp.SpanIdKey)
	if clean != "" {
		return clean
	}

	return defaultSpanIdKey
}

func (zp ZapSpanProcessor) GetTimestampKey() string {
	clean := strings.TrimSpace(zp.TimestampKey)
	if clean != "" {
		return clean
	}

	return defaultTimestampKey
}

func (zp ZapSpanProcessor) GetZapLevelKey() string {
	clean := strings.TrimSpace(zp.LevelKey)
	if clean != "" {
		return clean
	}

	return defaultLevelKey
}
