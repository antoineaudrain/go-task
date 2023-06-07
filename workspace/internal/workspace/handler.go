package workspace

import (
	"context"
	"github.com/google/uuid"
	"go-task/core/pkg/auth"
	customErrors "go-task/core/pkg/errors"
	pb "go-task/workspace/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedWorkspaceServiceServer
	service Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) Register(s *grpc.Server) {
	pb.RegisterWorkspaceServiceServer(s, h)
}

func (h *Handler) CreateWorkspace(ctx context.Context, req *pb.CreateWorkspaceRequest) (*pb.CreateWorkspaceResponse, error) {
	userId, err := auth.Authenticate(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token is missing: %v", err)
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Invalid user ID: %v", err)
	}

	user, err := h.service.CreateWorkspaceWithUser(req.GetName(), parsedUserId)
	if err != nil {
		switch err := err.(type) {
		case *customErrors.DatabaseError:
			return nil, status.Errorf(codes.Internal, "Failed to create workspace: %v", err)
		default:
			return nil, status.Errorf(codes.Internal, "Unexpected error: %v", err)
		}
	}

	createResponse := &pb.CreateWorkspaceResponse{
		Workspace: &pb.Workspace{
			Id:          user.ID.String(),
			Name:        user.Name,
			Description: user.Description,
		},
	}

	return createResponse, nil
}
