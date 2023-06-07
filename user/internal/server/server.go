package server

import (
	"context"
	"fmt"
	"go-task/core/pkg/logger"
	"go-task/user/internal/user"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	port       string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Run() error {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to listen on port %s", s.port), err)
		return err
	}

	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(s.interceptor))

	userHandler := user.NewHandler()
	userHandler.Register(s.grpcServer)

	logger.Info(fmt.Sprintf("Server started and listening on %s", conn.Addr().String()))

	if err := s.grpcServer.Serve(conn); err != nil {
		logger.Error("failed to serve", err)
		return err
	}

	return nil
}

func (s *Server) Shutdown() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
		logger.Info("Server stopped gracefully.")
	}
}

func (s *Server) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Info(fmt.Sprintf("Received request for method: %s", info.FullMethod))
	return handler(ctx, req)
}
