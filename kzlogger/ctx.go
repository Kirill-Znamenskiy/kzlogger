package kzlogger

import (
	"context"

	"github.com/Kirill-Znamenskiy/kzlogger/lga"
)

type Ctx = context.Context

var DefaultInCtxKey InCtxKeyType = "kzlogger"
var DefaultInCtxAttrsKey InCtxKeyType = "kzlogger-attrs"

func PutIntoCtx(ctx Ctx, l *Logger) Ctx {
	return PutIntoCtxKey(ctx, DefaultInCtxKey, l)
}

var SaveInCtx = PutIntoCtx

func ExtractFromCtx(ctx Ctx) (ret *Logger, err error) {
	return ExtractFromCtxKey(ctx, DefaultInCtxKey)
}

func From(ctx Ctx) *Logger {
	return FromKey(ctx, DefaultInCtxKey)
}

var Fr = From

func SetAttrsInCtx(ctx Ctx, attrs ...lga.Attr) Ctx {
	return SetAttrsInCtxKey(ctx, DefaultInCtxAttrsKey, attrs...)
}
func ExtractAttrsFromCtx(ctx Ctx) []lga.Attr {
	return ExtractAttrsFromCtxKey(ctx, DefaultInCtxAttrsKey)
}
func AddAttrsIntoCtx(ctx Ctx, attrs ...lga.Attr) Ctx {
	return AddAttrsIntoCtxKey(ctx, DefaultInCtxAttrsKey, attrs...)
}
