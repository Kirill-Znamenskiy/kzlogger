// Package lga = logger attribute
package lga

import (
	"fmt"
	"log/slog"

	"github.com/Kirill-Znamenskiy/kzlogger/callers"
)

type (
	Attr  = slog.Attr
	Value = slog.Value
)

func New(key string, val Value) Attr {
	return Attr{Key: key, Value: val}
}

var (
	Int      = slog.Int
	Int64    = slog.Int64
	Uint64   = slog.Uint64
	Bool     = slog.Bool
	Group    = slog.Group
	Str      = slog.String
	String   = slog.String
	Time     = slog.Time
	Duration = slog.Duration
	Any      = slog.Any
	Struct   = slog.Any
	Object   = slog.Any
)

func GroupAttrs(key string, attrs ...Attr) Attr {
	return New(key, slog.GroupValue(attrs...))
}

func Err(err error) Attr {
	return slog.Any("err", err)
}
func Error(err error) Attr {
	return slog.Any("error", err)
}
func Bytes(key string, bs []byte) Attr {
	return slog.String(key, string(bs))
}

func Stack() Attr {
	return StackKeySkipDepth("stack", 2, 32)
}
func StackKey(key string) Attr {
	return StackKeySkipDepth(key, 2, 32)
}
func StackSkip(skip int) Attr {
	return StackKeySkipDepth("stack", 2+skip, 32)
}
func StackKeySkip(key string, skip int) Attr {
	return StackKeySkipDepth(key, 2+skip, 32)
}
func StackSkipDepth(skip int, depth int) Attr {
	return StackKeySkipDepth("stack", 2+skip, depth)
}
func StackKeySkipDepth(key string, skip int, depth int) Attr {
	frames := callers.NewCallers(2+skip, depth).FramesSlice()
	attrs := make([]slog.Attr, 0, len(frames))
	for _, frame := range frames {
		attrs = append(attrs, slog.String(fmt.Sprintf("%s:%d", frame.File, frame.Line), frame.Function))
	}
	return GroupAttrs(key, attrs...)
}

func Caller() Attr              { return CallerKeySkip("caller", 1) }
func CallerKey(key string) Attr { return CallerKeySkip(key, 1) }
func CallerSkip(skip int) Attr  { return CallerKeySkip("caller", skip+1) }
func CallerKeySkip(key string, skip int) Attr {
	return StackKeySkipDepth(key, 2+skip, 1)
	//gattr := StackKeySkipDepth(key, 2+skip, 1)
	//return String(key, gattr.ValueJsonRawMessage.String())
}
