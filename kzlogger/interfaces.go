package kzlogger

import (
	"github.com/Kirill-Znamenskiy/kzlogger/lga"
	"github.com/Kirill-Znamenskiy/kzlogger/lgl"
)

type LI = LoggerInterface
type LoggerInterface interface {
	SetLevel(lvl lgl.Level)
	ParseAndSetLevel(level string) error

	AddAttrs(attrs ...lga.Attr)

	Log(Ctx, lgl.Level, any, ...lga.Attr)
	Debug(Ctx, any, ...lga.Attr)
	Info(Ctx, any, ...lga.Attr)
	Warn(Ctx, any, ...lga.Attr)
	Error(Ctx, any, ...lga.Attr)
	Critical(Ctx, any, ...lga.Attr)
}

type LoggerCloneWithInterface interface {
	CloneWithName(name string) LoggerInterface
	CloneWithAttrs(attrs ...lga.Attr) LoggerInterface
}
type LoggerCloneWithLoggerInterface interface {
	CloneWithName(name string) *Logger
	CloneWithAttrs(attrs ...lga.Attr) *Logger
}
