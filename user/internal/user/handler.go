package user

import (
	"context"
	"go-task/core/pkg/auth"
	customErrors "go-task/core/pkg/errors"
	pb "go-task/user/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := h.service.CreateUserWithHashedPassword(req.GetEmail(), req.GetPassword(), req.GetFullName())
	if err != nil {
		switch err := err.(type) {
		case *customErrors.DatabaseError:
			return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
		default:
			return nil, status.Errorf(codes.Internal, "Unexpected error: %v", err)
		}
	}

	createResponse := &pb.CreateUserResponse{
		User: &pb.User{
			Id:       user.ID.String(),
			Email:    user.Email,
			FullName: user.FullName,
		},
	}

	return createResponse, nil
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.service.AuthenticateUser(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password: %v", err)
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate refresh token: %v", err)
	}

	accessToken, err := auth.GenerateAccessToken(refreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate access token: %v", err)
	}

	loginResponse := &pb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: &pb.User{
			Id:       user.ID.String(),
			Email:    user.Email,
			FullName: user.FullName,
		},
	}

	return loginResponse, nil
}

func (h *Handler) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	accessToken, err := auth.GenerateAccessToken(req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate access token: %v", err)
	}

	refreshTokenResponse := &pb.RefreshAccessTokenResponse{AccessToken: accessToken}

	return refreshTokenResponse, nil
}
