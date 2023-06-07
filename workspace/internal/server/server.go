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
	log        logger.Logger
	port       string
}

func NewServer(port string, log logger.Logger) *Server {
	return &Server{
		port: port,
		log:  log,
	}
}

func (s *Server) Run() error {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to listen on port %s", s.port), "error", err)
		return err
	}

	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(s.interceptor))

	workspaceHandler := workspace.NewHandler(s.log)
	workspaceHandler.Register(s.grpcServer)

	s.log.Info(fmt.Sprintf("Server started and listening on %s", conn.Addr().String()))

	if err := s.grpcServer.Serve(conn); err != nil {
		s.log.Error("failed to serve", "error", err)
		return err
	}

	return nil
}

func (s *Server) Shutdown() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
		s.log.Info("Server stopped gracefully.")
	}
}

func (s *Server) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	s.log.Info(fmt.Sprintf("Received request for method: %s", info.FullMethod))

	accessToken := auth.ExtractAccessTokenFromContext(ctx)
	ctx = context.WithValue(ctx, "accessToken", accessToken)
	return handler(ctx, req)
}
