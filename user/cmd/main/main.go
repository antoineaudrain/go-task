package main

import (
	"go-task/core/pkg/logger"
	"go-task/user/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.Init("user", os.Getenv("ENV") != "production")
	defer logger.Close()

	s := server.NewServer()

	go func() {
		if err := s.Run(); err != nil {
			logger.Error("Server failed to start: %v", err)
		}
	}()

	waitForTerminationSignal()

	s.Stop()
}

func waitForTerminationSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	logger.Error("Termination signal received. Shutting down gracefully.")
}
