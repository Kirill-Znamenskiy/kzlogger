package kzlogger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/Kirill-Znamenskiy/kzlogger/lga"
	"github.com/Kirill-Znamenskiy/kzlogger/lgl"
)

type (
	Handler = slog.Handler
)

var (
	NewTextHandler = slog.NewTextHandler
	NewJSONHandler = slog.NewJSONHandler
)

var _ LoggerInterface = (*Logger)(nil)

// var _ LoggerCloneWithInterface = (*Logger)(nil)
var _ LoggerCloneWithLoggerInterface = (*Logger)(nil)

type Logger struct {
	*slog.Logger
	LevelVar      *slog.LevelVar
	InCtxAttrsKey InCtxKeyType
}

func New(slogh Handler, opts ...WithOption) (ret *Logger, err error) {
	lvl := &slog.LevelVar{}
	if slogh == nil {
		slogh = slog.NewJSONHandler(os.Stdout, nil)
	}
	slogh = &LeveledHandler{
		Handler: slogh,
		Level:   lvl,
	}
	slogger := slog.New(slogh)
	ret = &Logger{
		Logger:        slogger,
		LevelVar:      lvl,
		InCtxAttrsKey: DefaultInCtxAttrsKey,
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		ret, err = opt(ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}
func MustNew(slogh slog.Handler, opts ...WithOption) *Logger {
	ret, err := New(slogh, opts...)
	if err != nil {
		panic(err)
	}
	return ret
}

func (l *Logger) SetLevel(lvl lgl.Level) {
	l.LevelVar.Set(lvl)
}
func (l *Logger) ParseAndSetLevel(level string) (err error) {
	var lvl lgl.Level
	err = lvl.UnmarshalText([]byte(level))
	if err != nil {
		return err
	}
	l.LevelVar.Set(lvl)
	return nil
}

func (l *Logger) AddAttrs(attrs ...lga.Attr) {
	l.Logger = slog.New(l.Logger.Handler().WithAttrs(attrs))
}

func (l *Logger) CloneWithName(name string) *Logger {
	return l.CloneWithAttrs(lga.String("name", name))
}
func (l *Logger) CloneWithAttrs(attrs ...lga.Attr) *Logger {
	lvl := &slog.LevelVar{}
	lvl.Set(l.LevelVar.Level())
	ret := &Logger{
		Logger:        slog.New(l.Logger.Handler().WithAttrs(attrs)),
		LevelVar:      lvl,
		InCtxAttrsKey: l.InCtxAttrsKey,
	}
	return ret
}

func (l *Logger) PutIntoCtx(ctx Ctx) Ctx {
	return PutIntoCtx(ctx, l)
}
func (l *Logger) SaveInCtx(ctx Ctx) Ctx {
	return SaveInCtx(ctx, l)
}

func (l *Logger) Log(ctx Ctx, lvl lgl.Level, msg any, attrs ...lga.Attr) {
	finalAttrs := attrs
	if ctx == nil {
		ctx = context.Background()
	} else {
		frCtxAttrs := ExtractAttrsFromCtxKey(ctx, l.InCtxAttrsKey)
		finalAttrs = append(frCtxAttrs, attrs...)
	}

	finalMsg := ""
	switch typedMsg := msg.(type) {
	case string:
		finalMsg = typedMsg
	case error:
		err := typedMsg
		finalMsg = "err: " + err.Error()
		finalAttrs = append(finalAttrs, lga.Err(err))
	case fmt.Stringer:
		finalMsg = typedMsg.String()
	default:
		finalMsg = fmt.Sprint(msg)
	}
	l.Logger.LogAttrs(ctx, lvl, finalMsg, finalAttrs...)
}
func (l *Logger) Debug(ctx Ctx, msg any, attrs ...lga.Attr) {
	l.Log(ctx, lgl.LevelDebug, msg, attrs...)
}
func (l *Logger) Info(ctx Ctx, msg any, attrs ...lga.Attr) {
	l.Log(ctx, lgl.LevelInfo, msg, attrs...)
}
func (l *Logger) Warn(ctx Ctx, msg any, attrs ...lga.Attr) {
	l.Log(ctx, lgl.LevelWarn, msg, attrs...)
}
func (l *Logger) Error(ctx Ctx, msg any, attrs ...lga.Attr) {
	l.Log(ctx, lgl.LevelError, msg, attrs...)
}
func (l *Logger) Critical(ctx Ctx, msg any, attrs ...lga.Attr) {
	l.Log(ctx, lgl.LevelCritical, msg, attrs...)
}
