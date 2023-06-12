package user

import (
	"context"
	"go-task/core/pkg/auth"
	customErrors "go-task/core/pkg/errors"
	"go-task/core/pkg/logger"
	pb "go-task/user/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedUserServiceServer
	service Service
	log     logger.Logger
}

func NewHandler(log logger.Logger) *Handler {
	return &Handler{
		service: NewService(log),
		log:     log,
	}
}

func (h *Handler) Register(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, h)
}

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	h.log.Info("CreateUser called", "email", req.GetEmail())

	user, err := h.service.CreateUserWithHashedPassword(req.GetEmail(), req.GetPassword(), req.GetFullName())
	if err != nil {
		switch err := err.(type) {
		case *customErrors.DatabaseError:
			h.log.Error("Database error while creating user", "error", err)
			return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
		default:
			h.log.Error("Unexpected error while creating user", "error", err)
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

	h.log.Info("User created successfully", "userId", user.ID.String())

	return createResponse, nil
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	h.log.Info("Login called", "email", req.GetEmail())

	user, err := h.service.AuthenticateUser(req.GetEmail(), req.GetPassword())
	if err != nil {
		h.log.Error("Failed to authenticate user", "email", req.GetEmail(), "error", err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password: %v", err)
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID.String())
	if err != nil {
		h.log.Error("Failed to generate refresh token", "userID", user.ID.String(), "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to generate refresh token: %v", err)
	}

	accessToken, err := auth.GenerateAccessToken(refreshToken)
	if err != nil {
		h.log.Error("Failed to generate access token", "error", err)
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

	h.log.Info("Login successful", "email", req.GetEmail())

	return loginResponse, nil
}

func (h *Handler) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	h.log.Info("RefreshAccessToken called")

	accessToken, err := auth.GenerateAccessToken(req.GetRefreshToken())
	if err != nil {
		h.log.Error("Failed to generate access token", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to generate access token: %v", err)
	}

	refreshTokenResponse := &pb.RefreshAccessTokenResponse{AccessToken: accessToken}

	h.log.Info("Access token generated successfully")

	return refreshTokenResponse, nil
}
