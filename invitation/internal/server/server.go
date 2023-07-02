package server

import (
	"context"
	"fmt"
	"github.com/antoineaudrain/go-task/core/pkg/auth"
	pb "github.com/antoineaudrain/go-task/invitation/api"
	"github.com/antoineaudrain/go-task/invitation/internal/app/handlers"
	"github.com/antoineaudrain/go-task/invitation/internal/infrastructure/persistence"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	logger     *zap.Logger
	db         *gorm.DB
}

func NewServer(logger *zap.Logger, db *gorm.DB) *Server {
	return &Server{
		db:     db,
		logger: logger,
	}
}

func (s *Server) Start(port int) {
	conn, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to listen on port %d", port), zap.Error(err))
		return
	}

	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(s.unaryInterceptor))

	invitationRepo := persistence.NewInvitationRepository(s.db)

	handler := handlers.NewInvitationHandler(s.logger, invitationRepo)

	pb.RegisterInvitationServiceServer(s.grpcServer, handler)

	s.logger.Info(fmt.Sprintf("Server started and listening on %s", conn.Addr().String()))

	if err := s.grpcServer.Serve(conn); err != nil {
		s.logger.Error("Failed to serve", zap.Error(err))
		return
	}
}

func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
		s.logger.Info("Server stopped gracefully.")
	}
}

func (s *Server) unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.logger.Warn("Failed to extract metadata")
		return handler(ctx, req)
	}

	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		s.logger.Warn("Authorization header not found")
		return handler(ctx, req)
	}

	authToken := auth.ExtractAuthenticationToken(authHeaders[0])

	s.logger.Info("New request received",
		zap.String("method", info.FullMethod),
		zap.String("authToken", authToken),
	)

	userID, err := auth.ValidateAccessToken(authToken)
	if len(authHeaders) == 0 {
		s.logger.Warn("Invalid auth token")
		return handler(ctx, req)
	}

	ctx = context.WithValue(ctx, "authenticatedUserID", userID)

	resp, err := handler(ctx, req)

	s.logger.Info("Request completed",
		zap.String("method", info.FullMethod),
		zap.String("authToken", authToken),
		zap.String("status", status.Code(err).String()),
	)

	return resp, err
}
