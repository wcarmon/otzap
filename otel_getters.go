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
