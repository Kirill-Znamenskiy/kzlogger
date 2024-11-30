package kzlogger

import (
	"context"
	"errors"

	"github.com/Kirill-Znamenskiy/kzlogger/lga"
)

type InCtxKeyType string

var ErrContextIsNil = errors.New("context is nil")
var ErrNotFoundAnyLoggerInContext = errors.New("not found any logger in context")
var ErrNotFoundValidLoggerInContext = errors.New("not found valid logger in context")

func PutIntoCtxKey(ctx Ctx, key InCtxKeyType, l *Logger) Ctx {
	return context.WithValue(ctx, key, l)
}

var SaveInCtxKey = PutIntoCtxKey

func ExtractFromCtxKey(ctx Ctx, key InCtxKeyType) (ret *Logger, err error) {
	if ctx == nil {
		return nil, ErrContextIsNil
	}
	tmp := ctx.Value(key)
	if tmp == nil {
		return nil, ErrNotFoundAnyLoggerInContext
	}
	ret, ok := tmp.(*Logger)
	if !ok || ret == nil {
		return nil, ErrNotFoundValidLoggerInContext
	}
	return ret, nil
}

func FromKey(ctx Ctx, key InCtxKeyType) *Logger {
	ret, err := ExtractFromCtxKey(ctx, key)
	if err != nil {
		return nil
	}
	return ret
}

func SetAttrsInCtxKey(ctx Ctx, key InCtxKeyType, attrs ...lga.Attr) Ctx {
	return context.WithValue(ctx, key, attrs)
}
func ExtractAttrsFromCtxKey(ctx Ctx, key InCtxKeyType) []lga.Attr {
	if ctx == nil {
		return nil
	}
	tmp := ctx.Value(key)
	if tmp == nil {
		return nil
	}
	ret, ok := tmp.([]lga.Attr)
	if !ok || ret == nil {
		return nil
	}
	return ret
}
func AddAttrsIntoCtxKey(ctx Ctx, key InCtxKeyType, attrs ...lga.Attr) Ctx {
	aldAttrs := ExtractAttrsFromCtx(ctx)
	aldAttrs = append(aldAttrs, attrs...)
	return context.WithValue(ctx, key, aldAttrs)
}
