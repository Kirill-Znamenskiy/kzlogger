// Package lg = logger
package lg

import (
	"context"

	"github.com/Kirill-Znamenskiy/kzlogger/kzlogger"
	"github.com/Kirill-Znamenskiy/kzlogger/lga"
	"github.com/Kirill-Znamenskiy/kzlogger/lgl"
)

type (
	Ctx             = context.Context
	LI              = kzlogger.LoggerInterface
	LoggerInterface = kzlogger.LoggerInterface
	Logger          = kzlogger.Logger
	Handler         = kzlogger.Handler
)

var (
	NewLogger      = kzlogger.New
	MustNewLogger  = kzlogger.MustNew
	NewTextHandler = kzlogger.NewTextHandler
	NewJSONHandler = kzlogger.NewJSONHandler

	Fr             = kzlogger.From
	From           = kzlogger.From
	SaveInCtx      = kzlogger.SaveInCtx
	PutIntoCtx     = kzlogger.PutIntoCtx
	ExtractFromCtx = kzlogger.ExtractFromCtx
)

var DefaultLogger = kzlogger.MustNew(nil)
var DLG = DefaultLogger
var LGD = DefaultLogger

// Default returns the default kzlogger.Logger
func Default() *kzlogger.Logger {
	return DefaultLogger
}

var IsTryExtractWrkLoggerFromCtx = true

// Wrk returns the work kzlogger.Logger
func Wrk(ctx Ctx) *kzlogger.Logger {
	if IsTryExtractWrkLoggerFromCtx {
		l, err := kzlogger.ExtractFromCtx(ctx)
		if err == nil && l != nil {
			return l
		}
	}
	return DefaultLogger
}

func Log(ctx Ctx, lvl lgl.Level, msg any, attrs ...lga.Attr) {
	Wrk(ctx).Log(ctx, lvl, msg, attrs...)
}
func Debug(ctx Ctx, msg any, args ...lga.Attr) {
	Wrk(ctx).Debug(ctx, msg, args...)
}
func Info(ctx Ctx, msg any, args ...lga.Attr) {
	Wrk(ctx).Info(ctx, msg, args...)
}
func Warn(ctx Ctx, msg any, args ...lga.Attr) {
	Wrk(ctx).Warn(ctx, msg, args...)
}
func Error(ctx Ctx, msg any, args ...lga.Attr) {
	Wrk(ctx).Error(ctx, msg, args...)
}
func Critical(ctx Ctx, msg any, args ...lga.Attr) {
	Wrk(ctx).Critical(ctx, msg, args...)
}
