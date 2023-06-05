package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

var (
	Logger *zap.Logger
)

func Init(debug bool) {
	var core zapcore.Core
	if debug {
		consoleConfig := zap.NewDevelopmentEncoderConfig()
		consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)
		core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.DebugLevel)
	} else {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "user-service.log",
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

	Logger = zap.New(core)

	log.SetOutput(zap.NewStdLog(Logger).Writer())
}

func Close() {
	_ = Logger.Sync()
}
