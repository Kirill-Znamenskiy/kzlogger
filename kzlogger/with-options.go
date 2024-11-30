package kzlogger

import (
	"log/slog"

	"github.com/Kirill-Znamenskiy/kzlogger/lgl"
)

type WithOption func(*Logger) (*Logger, error)

func WithLevel(lvl lgl.Level) WithOption {
	return func(l *Logger) (*Logger, error) {
		l.SetLevel(lvl)
		return l, nil
	}
}

func WithInCtxAttrsKey(key InCtxKeyType) WithOption {
	return func(l *Logger) (*Logger, error) {
		l.InCtxAttrsKey = key
		return l, nil
	}
}

func WithSLogger(slogger *slog.Logger) WithOption {
	return func(l *Logger) (*Logger, error) {
		l.Logger = slogger
		return l, nil
	}
}
