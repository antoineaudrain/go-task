package main

import (
	"go-task/user-service/internal/server"
	"go-task/user-service/pkg/utils"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	utils.InitLogger(os.Getenv("MODE") != "production")
	defer utils.CloseLogger()

	log.SetOutput(zap.NewStdLog(utils.Logger).Writer())

	s := server.NewServer()

	go func() {
		if err := s.Run(); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	waitForTerminationSignal()

	s.Stop()
}

func waitForTerminationSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	log.Println("Termination signal received. Shutting down gracefully.")
}
