package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

var (
	logger *zap.Logger
)

func Init(namespace string) {
	var core zapcore.Core
	logFilename := fmt.Sprintf("logs/%s.log", namespace)

	if os.Getenv("ENV") != "production" {
		consoleConfig := zap.NewDevelopmentEncoderConfig()
		consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)
		core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.DebugLevel)
	} else {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logFilename,
			MaxSize:    500,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		})
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zap.InfoLevel,
		)
	}

	logger = zap.New(core)

	log.SetOutput(zap.NewStdLog(logger).Writer())
}

func Close() {
	_ = logger.Sync()
}

func Debug(message string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(message, args...))
}

func Info(message string, args ...interface{}) {
	logger.Info(fmt.Sprintf(message, args...))
}

func Warn(message string, args ...interface{}) {
	logger.Warn(fmt.Sprintf(message, args...))
}

func Error(message string, args ...interface{}) {
	logger.Error(fmt.Sprintf(message, args...))
}
