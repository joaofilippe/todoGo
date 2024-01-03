package logger

import (
	"log/slog"
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
func NewLogger(l LogOptions) *slog.Logger {
	return slog.Default()
}
