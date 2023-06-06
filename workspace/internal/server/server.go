package server

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go-task/core/pkg/logger"
	"go-task/workspace/internal/workspace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"strings"
	"time"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	conn, err := net.Listen("tcp", ":50052")
	if err != nil {
		logger.Error("failed to listen", err)
		return err
	}

	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(interceptor))

	_handler := workspace.NewHandler()
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

func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Info("Received request", info.FullMethod)

	accessToken, err := extractAccessTokenFromMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if accessToken == "" {
		return nil, status.Error(codes.Unauthenticated, "access token is missing")
	}

	userId, err := decodeAccessToken(accessToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	ctx = context.WithValue(ctx, "userId", userId)
	return handler(ctx, req)
}

func extractAccessTokenFromMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("failed to get metadata from context")
	}

	authorizationHeader := md.Get("authorization")
	if len(authorizationHeader) == 0 {
		return "", fmt.Errorf("authorization header is missing")
	}

	authValue := strings.TrimSpace(authorizationHeader[0])
	if !strings.HasPrefix(authValue, "Bearer ") {
		return "", fmt.Errorf("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authValue, "Bearer ")
	return accessToken, nil
}

func decodeAccessToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse access token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid access token")
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	fmt.Println(expirationTime)
	if time.Now().After(expirationTime) {
		return "", fmt.Errorf("access token has expired")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return "", fmt.Errorf("invalid userId claim in access token")
	}

	return userId, nil
}
