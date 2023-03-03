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
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	consoleColorBlack  = 30
	consoleColorBlue   = 34
	consoleColorGreen  = 32
	consoleColorRed    = 31
	consoleColorYellow = 33
)

// NewPrettyConsoleCore builds a Core which prints to stdout in a pretty format
// Inspired by zerolog's console writer
func NewPrettyConsoleCore(minLevel zapcore.Level) zapcore.Core {
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.ConsoleSeparator = " "
	cfg.EncodeLevel = ColorConsoleLevelEncoder
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("3:04:05PM")

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		zapcore.Lock(os.Stdout),
		minLevel)
}

// ColorConsoleLevelEncoder adds color and uses exactly 3-chars for level
func ColorConsoleLevelEncoder(
	l zapcore.Level,
	enc zapcore.PrimitiveArrayEncoder,
) {

	s := ""
	switch l {
	case zapcore.DebugLevel:
		s = fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(consoleColorBlue), "DBG")
	case zapcore.InfoLevel:
		s = fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(consoleColorGreen), "INF")
	case zapcore.WarnLevel:
		s = fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(consoleColorYellow), "WRN")
	case zapcore.ErrorLevel:
		s = fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(consoleColorRed), "ERR")
	case zapcore.FatalLevel:
		s = fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(consoleColorRed), "FTL")
	default:
		s = fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(consoleColorBlack), "LOG")
	}

	enc.AppendString(s)
}
