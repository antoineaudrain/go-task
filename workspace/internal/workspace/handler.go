package workspace

import (
	"context"
	"github.com/google/uuid"
	"go-task/core/pkg/auth"
	customErrors "go-task/core/pkg/errors"
	"go-task/core/pkg/logger"
	pb "go-task/workspace/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedWorkspaceServiceServer
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
	pb.RegisterWorkspaceServiceServer(s, h)
}

func (h *Handler) CreateWorkspace(ctx context.Context, req *pb.CreateWorkspaceRequest) (*pb.CreateWorkspaceResponse, error) {
	h.log.Info("CreateWorkspace called", "workspace_name", req.GetName())

	userId, err := auth.Authenticate(ctx)
	if err != nil {
		h.log.Error("Authentication failed", "error", err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token is missing: %v", err)
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		h.log.Error("Failed to parse user ID", "error", err)
		return nil, status.Errorf(codes.Internal, "Invalid user ID: %v", err)
	}

	user, err := h.service.CreateWorkspaceWithUser(req.GetName(), parsedUserId)
	if err != nil {
		switch err := err.(type) {
		case *customErrors.DatabaseError:
			h.log.Error("Database error while creating workspace", "error", err)
			return nil, status.Errorf(codes.Internal, "Failed to create workspace: %v", err)
		default:
			h.log.Error("Unexpected error while creating workspace", "error", err)
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

	h.log.Info("Workspace created successfully", "workspace_id", user.ID.String(), "workspace_name", user.Name)

	return createResponse, nil
}
