package server

import (
	"context"
	"go-task/core/pkg/logger"
	"go-task/user/internal/handler"
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
	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Error("failed to listen", err)
		return err
	}
	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))

	_handler := handler.NewHandler()
	_handler.Register(s.grpcServer)

	logger.Info("Server started and listening on :50051")

	if err := s.grpcServer.Serve(conn); err != nil {
		logger.Error("failed to serve", err)
		return err
	}

	return nil
}

func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
		logger.Info("Server stopped")
	}
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Info("Received request", info.FullMethod)
	return handler(ctx, req)
}
