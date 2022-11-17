// MIT License
//
// Copyright (c) 2022 Wilbur Carmon II
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

// Package otzap connect opentelemetry and zap
package otzap

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap/zapcore"
	"strings"
)

// OtelZapCore attaches zap log entries to current span
// Span can be attached directly to zap Log Entry or nested in an attached context.Context
// OTelZapCore ignores Log Entries which lack a Span
//
// OtelZapCore implements zapcore.Core
// See https://pkg.go.dev/go.uber.org/zap/zapcore#Core
// if log event contains a span, OTelZapCore forwards log events to span
type OTelZapCore struct {

	// Matches zap.Field.Key
	// default: "ctx"
	// Context may contain a Span
	//
	// See https://pkg.go.dev/go.uber.org/zap#Any
	// See https://pkg.go.dev/go.opentelemetry.io/otel/trace#SpanFromContext
	ContextAttrKey string

	// Matches zap.Field.Key
	// default: "span"
	// See https://pkg.go.dev/go.uber.org/zap#Any
	SpanAttrKey string

	// default: "logEventSource"
	EventSourceKey string

	// default: "zap"
	EventSourceValue string

	// default: "level"
	LevelKey string

	//TODO: use these
	extraFields []zapcore.Field
}

func (oc OTelZapCore) Check(
	ent zapcore.Entry,
	ce *zapcore.CheckedEntry,
) *zapcore.CheckedEntry {
	return ce.AddCore(ent, oc)
}

func (oc OTelZapCore) Enabled(zapcore.Level) bool {
	// OTelZapCore never writes to disk/io/network,
	// so OpenTelemetry is responsible for filtering
	return true
}

func (oc OTelZapCore) Sync() error {
	// nothing to do here
	return nil
}

// TODO: add tests for error handling
// TODO: add tests for mapping each type
// TODO: add tests for my span processor source
// TODO: add tests for self source
func (oc OTelZapCore) AddEventToSpan(
	span trace.Span,
	entry zapcore.Entry,
	fields []zapcore.Field,
) error {
	opts := make([]trace.EventOption, 0, 2+len(fields))
	opts = append(opts, trace.WithAttributes(attribute.String(oc.GetEventSourceKey(), oc.GetEventSourceValue())))
	opts = append(opts, trace.WithAttributes(attribute.String(oc.GetLevelKey(), entry.Level.String())))

	errorsToRecord := make([]error, 0)

	// -- copy attributes from zap log Entry to span event
	for _, f := range fields {
		if f.Key == oc.GetSpanAttrKey() || f.Key == oc.GetContextAttrKey() {
			continue
		}

		// -- Prevent infinite loop
		sourceIsMe := f.String == oc.GetEventSourceValue()
		sourceIsMyDual := f.String == defaultOtelSourceValue
		if f.Key == oc.GetEventSourceKey() && (sourceIsMe || sourceIsMyDual) {
			return nil
		}

		// -- Copy attributes from zap Event to Span
		switch f.Type {
		case zapcore.ErrorType:
			errorsToRecord = append(errorsToRecord, f.Interface.(error))

		case zapcore.DurationType:
			opts = append(opts, trace.WithAttributes(attribute.Int64(f.Key, f.Integer)))

		case zapcore.Int64Type:
			opts = append(opts, trace.WithAttributes(attribute.Int64(f.Key, f.Integer)))

		case zapcore.StringType:
			opts = append(opts, trace.WithAttributes(attribute.String(f.Key, f.String)))

		case zapcore.BoolType:
			opts = append(opts, trace.WithAttributes(attribute.Bool(f.Key, f.Interface.(bool))))

		// case zapcore.Float64Type:
		// 	opts = append(opts, trace.WithAttributes(attribute.Float64(f.Key, f.Float)))
		// case zapcore.ErrorType:
		// opts = append(opts, trace.WithAttributes(attribute.Er(f.Key, f.String)))

		default:
			opts = append(opts, trace.WithAttributes(attribute.String(
				f.Key, fmt.Sprintf("%v", f.Interface))))
		}

		// TODO: allow ignoring other keys via func on OtelZapCore
	}

	// TODO: allow veto based on everything available (func on oc)
	span.AddEvent(entry.Message, opts...)

	for _, err := range errorsToRecord {
		span.RecordError(err, opts...)
	}

	return nil
}

// AddEventToSpanInContext retrieves span from contex and adds the zap Entry
func (oc OTelZapCore) AddEventToSpanInContext(
	ctx context.Context,
	entry zapcore.Entry,
	fields []zapcore.Field,
) error {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		// eg. noop span or missing
		return nil
	}

	return oc.AddEventToSpan(span, entry, fields)
}

// TODO: add test that this doesn't affect original
func (oc OTelZapCore) With(fields []zapcore.Field) zapcore.Core {
	oc.extraFields = make([]zapcore.Field, 0, 2*len(fields))
	oc.extraFields = append(oc.extraFields, fields...)
	return oc
}

func (oc OTelZapCore) Write(
	entry zapcore.Entry,
	fields []zapcore.Field,
) error {
	if len(fields) == 0 {
		// No span, no context
		return nil
	}

	for _, f := range fields {
		if f.Key == oc.GetSpanAttrKey() {
			// -- found Span attr

			span, ok := f.Interface.(tracesdk.ReadWriteSpan)
			if !ok {
				return errors.New("invalid span type")
			}

			return oc.AddEventToSpan(span, entry, fields)
		}

		if f.Key == oc.ContextAttrKey {
			// -- found context.Context attr
			ctx, ok := f.Interface.(context.Context)
			if !ok {
				return errors.New("invalid context type")
			}

			return oc.AddEventToSpanInContext(ctx, entry, fields)
		}
	}

	return nil
}

func (oc OTelZapCore) Validate() error {

	if strings.TrimSpace(oc.GetContextAttrKey()) == "" {
		return errors.New("contextAttrKey required")
	}

	if strings.TrimSpace(oc.GetEventSourceKey()) == "" {
		return errors.New("eventSourceKey required")
	}

	if strings.TrimSpace(oc.GetEventSourceValue()) == "" {
		return errors.New("eventSourceValue required")
	}

	if strings.TrimSpace(oc.GetLevelKey()) == "" {
		return errors.New("levelKey required")
	}

	if strings.TrimSpace(oc.GetSpanAttrKey()) == "" {
		return errors.New("spanAttrKey required")
	}

	return nil
}
