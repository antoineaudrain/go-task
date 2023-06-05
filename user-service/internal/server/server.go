package server

import (
	"context"
	"go-task/user-service/internal/handler"
	"go-task/user-service/pkg/utils"
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
		utils.Logger.Error("failed to listen", zap.Error(err))
		return err
	}
	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))

	userHandler := handler.NewHandler()
	userHandler.Register(s.grpcServer)

	utils.Logger.Info("Server started and listening on :50051")

	if err := s.grpcServer.Serve(lis); err != nil {
		utils.Logger.Error("failed to serve", zap.Error(err))
		return err
	}

	return nil
}

func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
		utils.Logger.Info("Server stopped")
	}
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	utils.Logger.Info("Received request", zap.String("method", info.FullMethod))
	return handler(ctx, req)
}
