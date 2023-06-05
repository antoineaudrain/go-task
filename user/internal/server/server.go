package server

import (
	"context"
	"go-task/core/pkg/logger"
	"go-task/user/internal/handler"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Logger.Error("failed to listen", zap.Error(err))
		return err
	}
	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))

	userHandler := handler.NewHandler()
	userHandler.Register(s.grpcServer)

	logger.Logger.Info("Server started and listening on :50051")

	if err := s.grpcServer.Serve(lis); err != nil {
		logger.Logger.Error("failed to serve", zap.Error(err))
		return err
	}

	return nil
}

func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
		logger.Logger.Info("Server stopped")
	}
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Logger.Info("Received request", zap.String("method", info.FullMethod))
	return handler(ctx, req)
}
