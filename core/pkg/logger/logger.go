package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type (
	logger struct {
		zap *zap.Logger
	}

	Logger interface {
		Sync() error
		Debug(msg string, keysAndValues ...interface{})
		Info(msg string, keysAndValues ...interface{})
		Warn(msg string, keysAndValues ...interface{})
		Error(msg string, keysAndValues ...interface{})
	}
)

var _ Logger = (*logger)(nil)

func New(namespace string) Logger {
	var core zapcore.Core
	logFilename := fmt.Sprintf("logs/%s.log", namespace)

	switch os.Getenv("ENV") {
	case "production":
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
	default:
		consoleConfig := zap.NewDevelopmentEncoderConfig()
		consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)
		core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.DebugLevel)
	}

	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &logger{
		zap: l,
	}
}

func (l *logger) Sync() error {
	return l.zap.Sync()
}

func (l *logger) Debug(msg string, keysAndValues ...interface{}) {
	fields := make([]zap.Field, len(keysAndValues)/2)
	for i := 0; i < len(keysAndValues); i += 2 {
		key := keysAndValues[i].(string)
		value := keysAndValues[i+1]
		fields[i/2] = zap.Any(key, value)
	}
	l.zap.Debug(msg, fields...)
}

func (l *logger) Info(msg string, keysAndValues ...interface{}) {
	fields := make([]zap.Field, len(keysAndValues)/2)
	for i := 0; i < len(keysAndValues); i += 2 {
		key := keysAndValues[i].(string)
		value := keysAndValues[i+1]
		fields[i/2] = zap.Any(key, value)
	}
	l.zap.Info(msg, fields...)
}

func (l *logger) Warn(msg string, keysAndValues ...interface{}) {
	fields := make([]zap.Field, len(keysAndValues)/2)
	for i := 0; i < len(keysAndValues); i += 2 {
		key := keysAndValues[i].(string)
		value := keysAndValues[i+1]
		fields[i/2] = zap.Any(key, value)
	}
	l.zap.Warn(msg, fields...)
}

func (l *logger) Error(msg string, keysAndValues ...interface{}) {
	l.zap.Error(msg, zap.Any("error", keysAndValues))
}
