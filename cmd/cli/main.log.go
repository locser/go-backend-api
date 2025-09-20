package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()

	// sugar.Infof("Hello, %s", "World")
	// sugar.Infof("Hello, %s, age: %d", "World", 30)

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello, World", zap.String(
	// 	"name", "John",
	// ), zap.Int("age", 30))

	// logger.Info("Hello, World")

	// 	{"level":"info","msg":"Hello, World"}
	// {"level":"info","msg":"Hello, World, age: 30"}
	// {"level":"info","msg":"Hello, World","name":"John","age":30}
	// {"level":"info","msg":"Hello, World"}

	// LOGGER vs SUGAR:
	// Logger (structured): Type-safe, zero allocation, performance tốt, dùng cho production
	// Sugar (printf-style): Dễ đọc, có overhead, dùng cho development/debugging

	// development
	// logger, _ := zap.NewDevelopment()

	// logger.Info("Hello, World")

	// // production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello, World")

	// 	2025-09-20T20:20:52.268+0700    INFO    cli/main.log.go:31      Hello, World
	// {"level":"info","ts":1758374452.2692091,"caller":"cli/main.log.go:35","msg":"Hello, World"}

	// 3.

	encoder := getEnCodeLog()
	sync := getWriterSync()

	core := zapcore.NewCore(encoder, sync, zap.InfoLevel)

	logger := zap.New(core, zap.AddCaller()) //theem mau sac:  zap.AddCaller()

	logger.Info("Info log", zap.Int("line INFO", 1))
	logger.Error("Error log", zap.Int("line ERROR", 2))

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

// ghi vao file
func getWriterSync() zapcore.WriteSyncer {
	// Create log directory if it doesn't exist
	if err := os.MkdirAll("./log", 0755); err != nil {
		panic(err)
	}

	file, err := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
