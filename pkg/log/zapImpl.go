package log

import (
	"context"
	"io"

	"github.com/cloudwego/kitex/pkg/klog"
	"go.uber.org/zap"
)

var logger klog.FullLogger = &zapLogger{
	stdlog: initLog(),
	level:  klog.LevelInfo,
}

// SetOutput sets the output of default logger. By default, it is stderr.
func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

// SetLevel sets the level of logs below which logs will not be output.
// The default log level is LevelTrace.
// Note that this method is not concurrent-safe.
func SetLevel(lv klog.Level) {
	logger.SetLevel(lv)
}

// Fatal calls the default logger's Fatal method and then os.Exit(1).
func Fatal(v ...any) {
	logger.Fatal(v...)
}

// Error calls the default logger's Error method.
func Error(v ...any) {
	logger.Error(v...)
}

// Warn calls the default logger's Warn method.
func Warn(v ...any) {
	logger.Warn(v...)
}

// Notice calls the default logger's Notice method.
func Notice(v ...any) {
	logger.Notice(v...)
}

// Info calls the default logger's Info method.
func Info(v ...any) {
	logger.Info(v...)
}

// Debug calls the default logger's Debug method.
func Debug(v ...any) {
	logger.Debug(v...)
}

// Trace calls the default logger's Trace method.
func Trace(v ...any) {
	logger.Trace(v...)
}

// Fatalf calls the default logger's Fatalf method and then os.Exit(1).
func Fatalf(format string, v ...any) {
	logger.Fatalf(format, v...)
}

// Errorf calls the default logger's Errorf method.
func Errorf(format string, v ...any) {
	logger.Errorf(format, v...)
}

// Warnf calls the default logger's Warnf method.
func Warnf(format string, v ...any) {
	logger.Warnf(format, v...)
}

// Noticef calls the default logger's Noticef method.
func Noticef(format string, v ...any) {
	logger.Noticef(format, v...)
}

// Infof calls the default logger's Infof method.
func Infof(format string, v ...any) {
	logger.Infof(format, v...)
}

// Debugf calls the default logger's Debugf method.
func Debugf(format string, v ...any) {
	logger.Debugf(format, v...)
}

// Tracef calls the default logger's Tracef method.
func Tracef(format string, v ...any) {
	logger.Tracef(format, v...)
}

// CtxFatalf calls the default logger's CtxFatalf method and then os.Exit(1).
func CtxFatalf(ctx context.Context, format string, v ...any) {
	logger.CtxFatalf(ctx, format, v...)
}

// CtxErrorf calls the default logger's CtxErrorf method.
func CtxErrorf(ctx context.Context, format string, v ...any) {
	logger.CtxErrorf(ctx, format, v...)
}

// CtxWarnf calls the default logger's CtxWarnf method.
func CtxWarnf(ctx context.Context, format string, v ...any) {
	logger.CtxWarnf(ctx, format, v...)
}

// CtxNoticef calls the default logger's CtxNoticef method.
func CtxNoticef(ctx context.Context, format string, v ...any) {
	logger.CtxNoticef(ctx, format, v...)
}

// CtxInfof calls the default logger's CtxInfof method.
func CtxInfof(ctx context.Context, format string, v ...any) {
	logger.CtxInfof(ctx, format, v...)
}

// CtxDebugf calls the default logger's CtxDebugf method.
func CtxDebugf(ctx context.Context, format string, v ...any) {
	logger.CtxDebugf(ctx, format, v...)
}

// CtxTracef calls the default logger's CtxTracef method.
func CtxTracef(ctx context.Context, format string, v ...any) {
	logger.CtxTracef(ctx, format, v...)
}

type zapLogger struct {
	stdlog *zap.Logger
	level  klog.Level
}

/**
	Control
**/
func (ll *zapLogger) SetOutput(w io.Writer) {
	return
}

func (ll *zapLogger) SetLevel(lv klog.Level) {
	ll.level = lv
}

/**
	Logger
**/
func (ll *zapLogger) Fatal(v ...any) {
	ll.stdlog.Sugar().Fatal(v)
}

func (ll *zapLogger) Error(v ...any) {
	ll.stdlog.Sugar().Error(v)
}

func (ll *zapLogger) Warn(v ...any) {
	ll.stdlog.Sugar().Warn(v)
}

func (ll *zapLogger) Notice(v ...any) {
	ll.stdlog.Sugar().DPanic(v)
}

func (ll *zapLogger) Info(v ...any) {
	ll.stdlog.Sugar().Info(v)
}

func (ll *zapLogger) Debug(v ...any) {
	ll.stdlog.Sugar().Debug(v)
	println(1111)
}

func (ll *zapLogger) Trace(v ...any) {
	ll.stdlog.Sugar().Info(v)
}

/**
	FormatLogger
**/
func (ll *zapLogger) Fatalf(format string, v ...any) {
	ll.stdlog.Sugar().Fatalf(format, v)
}

func (ll *zapLogger) Errorf(format string, v ...any) {
	ll.stdlog.Sugar().Errorf(format, v)
}

func (ll *zapLogger) Warnf(format string, v ...any) {
	ll.stdlog.Sugar().Warnf(format, v)
}

func (ll *zapLogger) Noticef(format string, v ...any) {
	ll.stdlog.Sugar().DPanicf(format, v)
}

func (ll *zapLogger) Infof(format string, v ...any) {
	ll.stdlog.Sugar().Infof(format, v)
}

func (ll *zapLogger) Debugf(format string, v ...any) {
	ll.stdlog.Sugar().Debugf(format, v)
}

func (ll *zapLogger) Tracef(format string, v ...any) {
	ll.stdlog.Sugar().Infof(format, v)
}

/**
	CtxLogger
**/
func (ll *zapLogger) CtxFatalf(ctx context.Context, format string, v ...any) {
	ll.stdlog.Sugar().With("ctx", ctx).Fatalw(format, v...)
}

func (ll *zapLogger) CtxErrorf(ctx context.Context, format string, v ...any) {
	ll.stdlog.Sugar().With("ctx", ctx).Errorw(format, v...)
}

func (ll *zapLogger) CtxWarnf(ctx context.Context, format string, v ...any) {
	ll.stdlog.Sugar().With("ctx", ctx).Warnw(format, v...)
}

func (ll *zapLogger) CtxNoticef(ctx context.Context, format string, v ...any) {
	ll.stdlog.Sugar().With("ctx", ctx).DPanicw(format, v...)
}

func (ll *zapLogger) CtxInfof(ctx context.Context, format string, v ...any) {
	ll.stdlog.Sugar().With("ctx", ctx).Infow(format, v...)
}

func (ll *zapLogger) CtxDebugf(ctx context.Context, format string, v ...any) {
	ll.stdlog.Sugar().With("ctx", ctx).Debugw(format, v...)
}

func (ll *zapLogger) CtxTracef(ctx context.Context, format string, v ...any) {
	ll.stdlog.Sugar().With("ctx", ctx).Infow(format, v...)
}
