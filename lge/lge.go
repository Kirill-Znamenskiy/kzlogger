// Package lge = logger error
package lge

import (
	"github.com/Kirill-Znamenskiy/kzerror"

	"github.com/Kirill-Znamenskiy/kzlogger/lga"
)

type (
	ErrorWithMsg     = kzerror.ErrorWithAttrs
	ErrorWithAttrs   = kzerror.ErrorWithAttrs
	ErrorUnwrappable = kzerror.ErrorUnwrappable
)

var NewErr = kzerror.NewErr

func NewErrWithStack(msg string, attrs ...lga.Attr) *kzerror.Error {
	attrs = append(attrs, lga.StackSkip(1))
	return kzerror.NewErr(msg, attrs...)
}

func NewErrWithCaller(msg string, attrs ...lga.Attr) *kzerror.Error {
	attrs = append(attrs, lga.CallerSkip(1))
	return kzerror.NewErr(msg, attrs...)
}

var WrapErr = kzerror.WrapErr

func WrapErrWithStack(err error, attrs ...lga.Attr) *kzerror.Error {
	attrs = append(attrs, lga.StackSkip(1))
	return kzerror.WrapErr(err, attrs...)
}

func WrapErrWithCaller(err error, attrs ...lga.Attr) *kzerror.Error {
	attrs = append(attrs, lga.CallerSkip(1))
	return kzerror.WrapErr(err, attrs...)
}

var WrapErrMsg = kzerror.WrapErrMsg

func WrapErrMsgWithStack(err error, msg string, attrs ...lga.Attr) *kzerror.Error {
	attrs = append(attrs, lga.StackSkip(1))
	return kzerror.WrapErrMsg(err, msg, attrs...)
}

func WrapErrMsgWithCaller(err error, msg string, attrs ...lga.Attr) *kzerror.Error {
	attrs = append(attrs, lga.CallerSkip(1))
	return kzerror.WrapErrMsg(err, msg, attrs...)
}

var (
	New           = NewErr
	NewWithStack  = NewErrWithStack
	NewWithCaller = NewErrWithCaller

	Wrap           = WrapErr
	WrapWithStack  = WrapErrWithStack
	WrapWithCaller = WrapErrWithCaller

	WrapMsg           = WrapErrMsg
	WrapMsgWithStack  = WrapErrMsgWithStack
	WrapMsgWithCaller = WrapErrMsgWithCaller
)

var (
	IsNewErrAutoWithStack  = false
	IsNewErrAutoWithCaller = true
)

func NewErrAuto(msg string, attrs ...lga.Attr) *kzerror.Error {
	if IsNewErrAutoWithStack {
		attrs = append(attrs, lga.StackSkip(1))
	}
	if IsNewErrAutoWithCaller {
		attrs = append(attrs, lga.CallerSkip(1))
	}
	return kzerror.NewErr(msg, attrs...)
}

var NewAuto = NewErrAuto

var (
	IsWrapErrAutoWithStack  = false
	IsWrapErrAutoWithCaller = true
)

func WrapErrAuto(err error, attrs ...lga.Attr) *kzerror.Error {
	if IsWrapErrAutoWithStack {
		attrs = append(attrs, lga.StackSkip(1))
	}
	if IsWrapErrAutoWithCaller {
		attrs = append(attrs, lga.CallerSkip(1))
	}
	return kzerror.WrapErrMsg(err, "", attrs...)
}

var WrapAuto = WrapErrAuto

func WrapErrMsgAuto(err error, msg string, attrs ...lga.Attr) *kzerror.Error {
	if IsWrapErrAutoWithStack {
		attrs = append(attrs, lga.StackSkip(1))
	}
	if IsWrapErrAutoWithCaller {
		attrs = append(attrs, lga.CallerSkip(1))
	}
	return kzerror.WrapErrMsg(err, msg, attrs...)
}

var WrapMsgAuto = WrapErrMsgAuto
