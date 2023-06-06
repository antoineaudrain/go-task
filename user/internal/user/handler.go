package user

import (
	"context"
	"go-task/core/pkg/auth"
	pb "go-task/user/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

type Handler struct {
	pb.UnimplementedUserServiceServer
	service Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) Register(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, h)
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	user, err := h.service.Create(req.GetEmail(), req.GetPassword(), req.GetFullName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	return &pb.CreateResponse{
		User: &pb.User{
			Id:       user.ID.String(),
			Email:    user.Email,
			FullName: user.FullName,
		},
	}, nil
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.service.Authenticate(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password")
	}

	accessToken, err := auth.GenerateAccessToken(user.ID.String(), os.Getenv("SECRET_KEY"))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate access token")
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID.String(), os.Getenv("SECRET_KEY"))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate refresh token")
	}

	return &pb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: &pb.User{
			Id:       user.ID.String(),
			Email:    user.Email,
			FullName: user.FullName,
		},
	}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	claims, err := auth.ValidateRefreshToken(req.RefreshToken, "secretKey")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid refresh token")
	}

	userID, ok := claims["userId"].(string)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Invalid refresh token")
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	accessToken, err := auth.GenerateAccessToken(user.ID.String(), "secretKey")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate access token")
	}

	return &pb.RefreshTokenResponse{AccessToken: accessToken}, nil
}
