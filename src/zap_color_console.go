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
