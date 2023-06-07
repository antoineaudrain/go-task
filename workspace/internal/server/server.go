package server

import (
	"context"
	"fmt"
	"go-task/core/pkg/auth"
	"go-task/core/pkg/logger"
	pb "go-task/user/api"
	"go-task/workspace/internal/workspace"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	userClient pb.UserServiceClient
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

	workspaceHandler := workspace.NewHandler()
	workspaceHandler.Register(s.grpcServer)

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
	logger.Info("Received request", info.FullMethod)

	accessToken := auth.ExtractAccessTokenFromContext(ctx)
	ctx = context.WithValue(ctx, "accessToken", accessToken)
	return handler(ctx, req)
}
