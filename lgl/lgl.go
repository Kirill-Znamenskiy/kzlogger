// Package lgl = logger level
package lgl

import (
	"log/slog"
)

type Level = slog.Level

const (
	Debug    = slog.LevelDebug
	Info     = slog.LevelInfo
	Warn     = slog.LevelWarn
	Error    = slog.LevelError
	Critical = 11

	LevelDebug    = slog.LevelDebug
	LevelInfo     = slog.LevelInfo
	LevelWarn     = slog.LevelWarn
	LevelError    = slog.LevelError
	LevelCritical = 11
)
