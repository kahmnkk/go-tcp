package logger

import (
	"io"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

func NewLogger(level string, w io.Writer) zerolog.Logger {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		f := filepath.Base(file)
		return f + ":" + strconv.Itoa(line)
	}

	cw := zerolog.ConsoleWriter{
		Out:        w,
		TimeFormat: time.RFC3339Nano,
	}
	return zerolog.New(cw).With().Timestamp().Caller().Logger().Level(convertToLevel(level))
}

func convertToLevel(level string) zerolog.Level {
	switch level {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "painc":
		return zerolog.PanicLevel
	}
	return zerolog.InfoLevel
}
