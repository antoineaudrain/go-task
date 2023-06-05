package auth

import (
	"context"
	"go-task/auth-service/internal/event"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	pb "go-task/auth-service/pkg/pb"
)

func NewServer() *Server {
	return &Server{Service: ServiceImpl{}}
}

func (s *Server) RegisterUser(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Register Action Received: %v", in.GetEmail())

	token, err := s.Service.Create(in.GetEmail(), in.GetPassword())
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to register user")
	}

	event.PublishUserRegistered("123", in.GetEmail())

	return &pb.RegisterResponse{Token: token}, nil
}

func (s *Server) LoginUser(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Login Action Received: %v", in.GetEmail())

	token, err := s.Service.Login(in.GetEmail(), in.GetPassword())
	if err != nil {
		log.Printf("Error signing in user: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to login user")
	}

	return &pb.LoginResponse{Token: token}, nil
}
