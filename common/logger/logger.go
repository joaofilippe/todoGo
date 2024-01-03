package logger

import (
	"log/slog"
	"os"
	"runtime"
)

type Logger struct {
	Logger  *slog.Logger
	Options LogOptions
}

type LogOptions struct {
	Level       string
	Environment string
}

// NewLogger returns a new logger
func NewLogger(logOptions LogOptions) *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &Logger{
		Logger:  logger,
		Options: logOptions,
	}
}

func (l *Logger) ErrorF(err error) {
	_, _, line, _ := runtime.Caller(1)

	l.Logger.Error(err.Error(), "line=", line)
}

func (l *Logger) InfoF(message string) {
	_, file, line, _ := runtime.Caller(1)

	l.Logger.Info(message, "line", line, "file", file)
}
