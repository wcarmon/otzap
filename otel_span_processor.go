// MIT License
//
// Copyright (c) 2023 Wilbur Carmon II
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package otzap

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// ZapSpanProcessor forwards OpenTelemetry::Span events to a Zap logger
// See https://pkg.go.dev/go.opentelemetry.io/otel/sdk/trace#SpanProcessor
type ZapSpanProcessor struct {
	// debug | info | warn | error | fatal
	DefaultLevel string
	Logger       *zap.Logger

	EventSourceKey string

	// Useful for preventing infinite loops
	EventSourceValue string

	LevelKey     string
	SpanIdKey    string
	TimestampKey string

	// See https://cloud.google.com/resource-manager/docs/creating-managing-projects#before_you_begin
	GoogleCloudProjectId string
}

func (zp ZapSpanProcessor) OnStart(context.Context, tracesdk.ReadWriteSpan) {
	// nothing to do
}

func (zp ZapSpanProcessor) OnEnd(span tracesdk.ReadOnlySpan) {
	if len(span.Events()) == 0 {
		return
	}

	// NOTE: span events iterate from least to most recent
	for _, evt := range span.Events() {
		zp.LogEvent(evt, span.Attributes(), span.SpanContext())
	}
}

func (zp ZapSpanProcessor) ForceFlush(context.Context) error {
	return nil
}

func (zp ZapSpanProcessor) Shutdown(context.Context) error {
	return nil
}

// LogEvent delegates to configured Writers
func (zp ZapSpanProcessor) LogEvent(
	currentEvt tracesdk.Event,
	spanAttributes []attribute.KeyValue,
	spanCtx trace.SpanContext) {

	hexSpanId := spanCtx.SpanID().String()
	hexTraceId := spanCtx.TraceID().String()

	fields := make([]zap.Field, 0, 32)
	fields = append(fields, zap.String(zp.GetSpanIdKey(), hexSpanId))

	// Google cloud uses this
	// See https://cloud.google.com/trace/docs/trace-log-integration#associating
	if zp.GoogleCloudProjectId != "" {
		fields = append(fields,
			zap.String("trace",
				fmt.Sprintf("projects/%s/traces/%s",
					zp.GoogleCloudProjectId,
					hexTraceId)))
		// TODO: shouldn't I use the spanId?
	}

	// TODO: This just adds duplicate json field (which is undesirable)
	// TODO: need to use time field name and format for cloud provider (eg. google)
	fields = append(fields, zap.Time(zp.GetTimestampKey(), currentEvt.Time.UTC()))

	// -- Copy span attributes (lower priority)
	for _, attr := range spanAttributes {
		key := string(attr.Key)
		fields = append(fields, zap.Any(key, attr.Value.AsInterface()))
	}

	logger := zp.GetLogger()
	var logLevel zapcore.Level

	// -- Copy event attributes (higher priority)
	for _, attr := range currentEvt.Attributes {
		key := string(attr.Key)
		if key == zp.GetZapLevelKey() {
			logLevel := zp.GetZapLevel(attr.Value.AsString())

			if !logger.Core().Enabled(logLevel) {
				// logger will ignore this level
				return
			}

			continue
		}

		// -- Prevent infinite loop
		v := attr.Value.AsString()
		sourceIsMe := v == zp.GetEventSourceValue()
		sourceIsMyDual := v == defaultZapSourceValue

		if key == zp.GetEventSourceKey() && (sourceIsMe || sourceIsMyDual) {
			return
		}

		fields = append(fields, zap.Any(key, attr.Value.AsInterface()))
	}

	fields = append(fields, zap.String(zp.GetEventSourceKey(), zp.GetEventSourceValue()))

	ce := logger.Check(logLevel, currentEvt.Name)
	if ce == nil {
		return
	}

	// TODO: allow veto based on available vars (func on z)

	ce.Write(fields...)
}

func (zp ZapSpanProcessor) GetZapLevel(raw string) zapcore.Level {
	clean := strings.ToLower(strings.TrimSpace(raw))

	if clean == "" {
		clean = zp.DefaultLevel
	}

	switch clean {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel

	default:
		return zapcore.DebugLevel
	}
}

func (zp ZapSpanProcessor) Validate() error {
	if zp.GetLogger() == nil {
		return errors.New("logger required")
	}

	if strings.TrimSpace(zp.GetEventSourceKey()) == "" {
		return errors.New("eventSourceKey required")
	}

	if strings.TrimSpace(zp.GetEventSourceValue()) == "" {
		return errors.New("eventSourceValue required")
	}

	if strings.TrimSpace(zp.GetZapLevelKey()) == "" {
		return errors.New("zapLevelKey required")
	}

	if strings.TrimSpace(zp.GetSpanIdKey()) == "" {
		return errors.New("spanIdKey required")
	}

	if strings.TrimSpace(zp.GetTimestampKey()) == "" {
		return errors.New("timestampKey required")
	}

	return nil
}
