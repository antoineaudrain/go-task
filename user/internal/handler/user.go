package handler

import (
	"context"
	pb "go-task/user/api"
	"go-task/user/internal/service"
	"go-task/user/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserService
}

func NewHandler() *Handler {
	return &Handler{
		userService: service.NewUserService(),
	}
}

func (h *Handler) Register(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, h)
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	user, err := h.userService.Create(req.GetEmail(), req.GetPassword(), req.GetFullName())
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
	user, err := h.userService.GetUserByEmail(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), "secretKey")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate access token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String(), "secretKey")
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
	claims, err := utils.ValidateRefreshToken(req.RefreshToken, "secretKey")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid refresh token")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Invalid refresh token")
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	// Generate a new access token
	accessToken, err := utils.GenerateAccessToken(user.ID.String(), "secretKey")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate access token")
	}

	return &pb.RefreshTokenResponse{AccessToken: accessToken}, nil
}
