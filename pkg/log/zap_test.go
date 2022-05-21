package log

import (
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

// test package level functions without format
func normalOutput(t *testing.T, testLevel klog.Level, want string, args ...any) {
	// buf := new(bytes.Buffer)
	// SetOutput(buf)
	// defer SetOutput(os.Stderr)
	// switch testLevel {
	// case LevelTrace:
	// 	Trace(args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelDebug:
	// 	Debug(args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelInfo:
	// 	Info(args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelNotice:
	// 	Notice(args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelWarn:
	// 	Warn(args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelError:
	// 	Error(args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelFatal:
	// 	t.Fatal("fatal method cannot be tested")
	// default:
	// 	t.Errorf("unknow level: %d", testLevel)
	// }
}

// test package level Ctx-related functions with 'format'
func ctxOutput(t *testing.T, testLevel klog.Level, want, format string, args ...any) {
	// buf := new(bytes.Buffer)
	// SetOutput(buf)
	// defer SetOutput(os.Stderr)

	// // the default logger implementation of CtxLogger is same as FormatLogger, no context handle now
	// ctx := context.Background()

	// switch testLevel {
	// case LevelTrace:
	// 	CtxTracef(ctx, format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelDebug:
	// 	CtxDebugf(ctx, format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelInfo:
	// 	CtxInfof(ctx, format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelNotice:
	// 	CtxNoticef(ctx, format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelWarn:
	// 	CtxWarnf(ctx, format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelError:
	// 	CtxErrorf(ctx, format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelFatal:
	// 	t.Fatal("fatal method cannot be tested")
	// default:
	// 	t.Errorf("unknow level: %d", testLevel)
	// }
}

// test package level functions with 'format'
func formatOutput(t *testing.T, testLevel klog.Level, want, format string, args ...any) {
	// buf := new(bytes.Buffer)
	// SetOutput(buf)
	// defer SetOutput(os.Stderr)
	// switch testLevel {
	// case LevelTrace:
	// 	Tracef(format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelDebug:
	// 	Debugf(format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelInfo:
	// 	Infof(format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelNotice:
	// 	Noticef(format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelWarn:
	// 	Warnf(format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelError:
	// 	Errorf(format, args...)
	// 	test.Assert(t, buf.String() == want)
	// case LevelFatal:
	// 	t.Fatal("fatal method cannot be tested")
	// default:
	// 	t.Errorf("unknow level: %d", testLevel)
	// }
}

var strs = []string{
	"[Trace] ",
	"[Debug] ",
	"[Info] ",
	"[Notice] ",
	"[Warn] ",
	"[Error] ",
	"[Fatal] ",
}

func TestOutput(t *testing.T) {
	// l := DefaultLogger().(*defaultLogger)
	// oldFlags := l.stdlog.Flags()
	// l.stdlog.SetFlags(0)
	// defer l.stdlog.SetFlags(oldFlags)
	// defer SetLevel(klog.LevelInfo)
	InitLog()
	tests := []struct {
		format      string
		args        []any
		testLevel   klog.Level
		loggerLevel klog.Level
		want        string
	}{
		{"%s", []any{"LevelNotice test"}, klog.LevelNotice, klog.LevelInfo, strs[klog.LevelNotice] + "LevelNotice test\n"},
		{"%s %s", []any{"LevelInfo", "test"}, klog.LevelInfo, klog.LevelWarn, ""},
		{"%s%s", []any{"LevelDebug", "Test"}, klog.LevelDebug, klog.LevelDebug, strs[klog.LevelDebug] + "LevelDebugTest\n"},
		{"%s", []any{"LevelTrace test"}, klog.LevelTrace, klog.LevelTrace, strs[klog.LevelTrace] + "LevelTrace test\n"},
		{"%s", []any{"LevelError test"}, klog.LevelError, klog.LevelInfo, strs[klog.LevelError] + "LevelError test\n"},
		{"%s", []any{"LevelWarn test"}, klog.LevelWarn, klog.LevelWarn, strs[klog.LevelWarn] + "LevelWarn test\n"},
	}

	for _, tt := range tests {
		klog.Error(tt.args...)
		klog.Info(tt.args...)
		klog.Debug(tt.args...)
		// SetLevel(tt.loggerLevel)
		// normalOutput(t, tt.testLevel, tt.want, tt.args...)
		// formatOutput(t, tt.testLevel, tt.want, tt.format, tt.args...)
		// ctxOutput(t, tt.testLevel, tt.want, tt.format, tt.args...)
	}
}
