package log

import (
	"douyin-micro/pkg/constants"
	"douyin-micro/pkg/utils"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	level   zapcore.Level // 日志等级
	options []zap.Option  // 配置项
)

func initLog() *zap.Logger {

	createRootDir() // 创建日志目录

	setLogLevel() // 设置日志等级

	if constants.LogShowLine {
		options = append(options, zap.AddCaller())
	}
	Logger := zap.New(getZapCore(), options...)
	Logger.Info("log init success!")
	return Logger
}

func createRootDir() {
	rootDir := constants.LogRootDir
	if ok, _ := utils.PathExists(rootDir); !ok {
		_ = os.Mkdir(rootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch constants.LogLevel {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[2006-01-02 15:04:05.000]"))
	}

	encodeConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(fmt.Sprintf("%s.%s", constants.AppEnv, l.String()))
	}
	if constants.LogFormat == "json" {
		encoder = zapcore.NewJSONEncoder(encodeConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encodeConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   constants.LogRootDir + "/" + constants.LogFileName,
		MaxSize:    constants.LogMaxSize,
		MaxBackups: constants.LogMaxBackups,
		MaxAge:     constants.LogMaxAge,
		Compress:   constants.LogCompress,
	}

	return zapcore.AddSync(file)
}
