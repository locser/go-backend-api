package logger

import (
	"go-backend/pkg/setting"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Log_level // "debug"
	// debug -> info -> warn -> error -> fatal -> panic
	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "default":
		level = zapcore.InfoLevel
	}

	encoder := getEnCodeLog()

	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook),
	), level)

	// logger := zap.New(core, zap.AddCaller()) //theem mau sac:  zap.AddCaller()

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEnCodeLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// {"level":"info","ts":1758374452.2692091,"caller":"cli/main.log.go:35","msg":"Hello, World"}
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 1758374452.2692091 -> 2025-09-20T20:20:52.268+0700

	encodeConfig.TimeKey = "time"

	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// caller : hien thi cli/main.log.go:31
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}
