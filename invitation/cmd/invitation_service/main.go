package main

import (
	"github.com/antoineaudrain/go-task/invitation/internal/server"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := NewLogger("invitation")
	defer func() { _ = logger.Sync() }()

	db, err := NewDatabase()
	if err != nil {
		logger.Error("Server failed to start", zap.Error(err))
		return
	}

	srv := server.NewServer(logger, db)
	go func() { srv.Start(50051) }()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	logger.Info("Shutting down server gracefully...")

	srv.Stop()
}

func NewLogger(service string) *zap.Logger {
	env := os.Getenv("ENV")
	logger, _ := zap.NewProduction(zap.Fields(
		zap.String("env", env),
		zap.String("service", service),
	))

	//if env == "" || env == "development" {
	//	logger, _ = zap.NewDevelopment()
	//}

	return logger
}

func NewDatabase() (*gorm.DB, error) {
	logger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,            // Don't include params in the SQL log
			Colorful:                  false,           // Disable color
		},
	)

	config := &gorm.Config{
		Logger: logger,
	}
	return gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), config)
}
