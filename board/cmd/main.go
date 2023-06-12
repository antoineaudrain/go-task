package main

import (
	"go-task/board/internal/server"
	"go-task/core/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.New("board")
	defer func() {
		if err := log.Sync(); err != nil {
			log.Error("Error syncing log", "error", err)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	s := server.NewServer(port, log)

	go func() {
		if err := s.Run(); err != nil {
			log.Error("Server failed to start", "error", err)
		}
	}()

	waitForTerminationSignal(log)

	s.Shutdown()
}

func waitForTerminationSignal(log logger.Logger) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	log.Info("Shutting down server gracefully...")
}
