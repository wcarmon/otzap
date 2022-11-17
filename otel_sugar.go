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

package otzap

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// AddDebugEvent simplifies span.debug(msg)
// AddDebugEvent adds an event to the span
// Any related attributes should be added to the span
func AddDebugEvent(span trace.Span, msg string, err ...error) {
	span.AddEvent(
		msg,
		trace.WithAttributes(attribute.String(defaultLevelKey, "debug")))

	RecordErrors(span, err...)
}

// AddInfoEvent simplifies span.info(msg)
// AddInfoEvent adds an event to the span
// Any related attributes should be added to the span
func AddInfoEvent(span trace.Span, msg string, err ...error) {
	span.AddEvent(
		msg,
		trace.WithAttributes(attribute.String(defaultLevelKey, "info")))

	RecordErrors(span, err...)
}

// AddWarnEvent simplifies span.warn(msg)
// AddWarnEvent adds an event to the span
// Any related attributes should be added to the span
func AddWarnEvent(span trace.Span, msg string, err ...error) {
	span.AddEvent(
		msg,
		trace.WithAttributes(attribute.String(defaultLevelKey, "warn")))

	RecordErrors(span, err...)
}

// AddErrorEvent simplifies span.error(msg)
// AddErrorEvent adds an event to the span
// Any related attributes should be added to the span
func AddErrorEvent(span trace.Span, msg string, err ...error) {
	span.AddEvent(
		msg,
		trace.WithAttributes(attribute.String(defaultLevelKey, "error")))

	RecordErrors(span, err...)
}

// AddFatalEvent simplifies span.fatal(msg)
// AddFatalEvent adds an event to the span
// Any related attributes should be added to the span
func AddFatalEvent(span trace.Span, msg string, err ...error) {
	span.AddEvent(
		msg,
		trace.WithAttributes(attribute.String(defaultLevelKey, "fatal")))

	RecordErrors(span, err...)
}

// AddDebugEventToSpan retrieves span from context and
// adds error to span when present
//
// AddDebugEventToSpan sets level to debug
func AddDebugEventToSpan(ctx context.Context, msg string, err ...error) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		// eg. noop span or missing
		return
	}

	AddDebugEvent(span, msg, err...)
}

// AddInfoEventToSpan retrieves span from context and
// adds error to span when present
//
// AddInfoEventToSpan sets level to info
func AddInfoEventToSpan(ctx context.Context, msg string, err ...error) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		// eg. noop span or missing
		return
	}

	AddInfoEvent(span, msg, err...)
}

// AddWarnEventToSpan retrieves span from context and
// adds error to span when present
//
// AddWarnEventToSpan sets level to warn
func AddWarnEventToSpan(ctx context.Context, msg string, err ...error) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		// eg. noop span or missing
		return
	}

	AddWarnEvent(span, msg, err...)
}

// AddErrorEventToSpan retrieves span from context and
// adds error to span when present
//
// AddErrorEventToSpan sets level to error
func AddErrorEventToSpan(ctx context.Context, msg string, err ...error) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		// eg. noop span or missing
		return
	}

	AddErrorEvent(span, msg, err...)
}

// AddFatalEventToSpan retrieves span from context and
// adds error to span when present
//
// AddFatalEventToSpan sets level to fatal
func AddFatalEventToSpan(ctx context.Context, msg string, err ...error) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		// eg. noop span or missing
		return
	}

	AddFatalEvent(span, msg, err...)
}

// RecordErrors is a low-level method, prefer the methods above
// RecordErrors simplifies recording errors correctly for Jaeger,
// Google Cloud Trace, AWS XRay, etc
func RecordErrors(span trace.Span, err ...error) {
	if err == nil || len(err) == 0 {
		return
	}

	span.SetStatus(codes.Error, "")

	// Jaeger needs this
	span.SetAttributes(attribute.Bool("error", true))

	for _, err := range err {
		span.RecordError(err)
	}
}
