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
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

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
