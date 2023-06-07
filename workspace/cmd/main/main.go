package main

import (
	"go-task/core/pkg/logger"
	"go-task/workspace/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.Init("workspace")
	defer logger.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	s := server.NewServer(port)

	go func() {
		if err := s.Run(); err != nil {
			logger.Error("Server failed to start: %v", err)
		}
	}()

	waitForTerminationSignal()

	s.Shutdown()
}

func waitForTerminationSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	logger.Info("Shutting down server gracefully...")
}
