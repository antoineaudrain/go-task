package workspace

import (
	"context"
	"github.com/google/uuid"
	"go-task/core/pkg/auth"
	customErrors "go-task/core/pkg/errors"
	"go-task/core/pkg/format"
	"go-task/core/pkg/logger"
	pb "go-task/workspace/pkg/proto"
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

	userID, err := auth.Authenticate(ctx)
	if err != nil {
		h.log.Error("Authentication failed", "error", err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token is missing: %v", err)
	}

	workspace, err := h.service.CreateWorkspaceWithUser(userID, req.GetName())
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
			Id:        workspace.ID.String(),
			Name:      workspace.Name,
			CreatedAt: format.TimeAsString(workspace.CreatedAt),
			UpdatedAt: format.TimeAsString(workspace.UpdatedAt),
		},
	}

	h.log.Info("Workspace created successfully", "workspace_id", workspace.ID.String(), "workspace_name", workspace.Name)

	return createResponse, nil
}

func (h *Handler) GetWorkspace(ctx context.Context, req *pb.GetWorkspaceRequest) (*pb.GetWorkspaceResponse, error) {
	h.log.Info("GetWorkspace called", "workspace_id", req.GetWorkspaceID())

	userID, err := auth.Authenticate(ctx)
	if err != nil {
		h.log.Error("Authentication failed", "error", err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token is missing: %v", err)
	}

	workspaceID, err := uuid.Parse(req.GetWorkspaceID())
	if err != nil {
		h.log.Error("Failed to parse workspace ID", "error", err)
		return nil, status.Errorf(codes.Internal, "Invalid workspace ID: %v", err)
	}

	workspace, err := h.service.GetWorkspace(userID, workspaceID)
	if err != nil {
		switch err := err.(type) {
		case *customErrors.DatabaseError:
			h.log.Error("Database error while getting workspace", "error", err)
			return nil, status.Errorf(codes.Internal, "Failed to get workspace: %v", err)
		default:
			h.log.Error("Unexpected error while creating workspace", "error", err)
			return nil, status.Errorf(codes.Internal, "Unexpected error: %v", err)
		}
	}

	getResponse := &pb.GetWorkspaceResponse{
		Workspace: &pb.Workspace{
			Id:        workspace.ID.String(),
			Name:      workspace.Name,
			CreatedAt: format.TimeAsString(workspace.CreatedAt),
			UpdatedAt: format.TimeAsString(workspace.UpdatedAt),
		},
	}

	return getResponse, nil
}

func (h *Handler) ListWorkspaces(ctx context.Context, req *pb.ListWorkspaceRequest) (*pb.ListWorkspaceResponse, error) {
	h.log.Info("ListWorkspaces called")

	userID, err := auth.Authenticate(ctx)
	if err != nil {
		h.log.Error("Authentication failed", "error", err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token is missing: %v", err)
	}

	workspaces, err := h.service.ListWorkspaces(userID)
	if err != nil {
		switch err := err.(type) {
		case *customErrors.DatabaseError:
			h.log.Error("Database error while getting workspace", "error", err)
			return nil, status.Errorf(codes.Internal, "Failed to get workspace: %v", err)
		default:
			h.log.Error("Unexpected error while creating workspace", "error", err)
			return nil, status.Errorf(codes.Internal, "Unexpected error: %v", err)
		}
	}

	var pbWorkspaces []*pb.Workspace
	for _, workspace := range workspaces {
		pbWorkspace := &pb.Workspace{
			Id:        workspace.ID.String(),
			Name:      workspace.Name,
			CreatedAt: format.TimeAsString(workspace.CreatedAt),
			UpdatedAt: format.TimeAsString(workspace.UpdatedAt),
		}
		pbWorkspaces = append(pbWorkspaces, pbWorkspace)
	}

	listResponse := &pb.ListWorkspaceResponse{
		Workspaces: pbWorkspaces,
	}

	return listResponse, nil
}

func (h *Handler) UpdateWorkspace(ctx context.Context, req *pb.UpdateWorkspaceRequest) (*pb.UpdateWorkspaceResponse, error) {
	h.log.Info("UpdateWorkspace called", "workspace_id", req.GetWorkspaceID())

	userID, err := auth.Authenticate(ctx)
	if err != nil {
		h.log.Error("Authentication failed", "error", err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token is missing: %v", err)
	}

	workspaceID, err := uuid.Parse(req.GetWorkspaceID())
	if err != nil {
		h.log.Error("Failed to parse workspace ID", "error", err)
		return nil, status.Errorf(codes.Internal, "Invalid workspace ID: %v", err)
	}

	workspace, err := h.service.UpdateWorkspace(userID, workspaceID, req.GetName())
	if err != nil {
		switch err := err.(type) {
		case *customErrors.DatabaseError:
			h.log.Error("Database error while updating workspace", "error", err)
			return nil, status.Errorf(codes.Internal, "Failed to update workspace: %v", err)
		case *customErrors.NotFoundError:
			h.log.Error("Workspace not found", "error", err)
			return nil, status.Errorf(codes.NotFound, "Workspace not found: %v", err)
		default:
			h.log.Error("Unexpected error while updating workspace", "error", err)
			return nil, status.Errorf(codes.Internal, "Unexpected error: %v", err)
		}
	}

	updateResponse := &pb.UpdateWorkspaceResponse{
		Workspace: &pb.Workspace{
			Id:        workspace.ID.String(),
			Name:      workspace.Name,
			CreatedAt: format.TimeAsString(workspace.CreatedAt),
			UpdatedAt: format.TimeAsString(workspace.UpdatedAt),
		},
	}

	h.log.Info("Workspace updated successfully", "workspace_id", workspace.ID.String(), "workspace_name", workspace.Name)

	return updateResponse, nil
}

func (h *Handler) DeleteWorkspace(ctx context.Context, req *pb.DeleteWorkspaceRequest) (*pb.DeleteWorkspaceResponse, error) {
	h.log.Info("DeleteWorkspace called", "workspace_id", req.GetWorkspaceID())

	userID, err := auth.Authenticate(ctx)
	if err != nil {
		h.log.Error("Authentication failed", "error", err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token is missing: %v", err)
	}

	workspaceID, err := uuid.Parse(req.GetWorkspaceID())
	if err != nil {
		h.log.Error("Failed to parse workspace ID", "error", err)
		return nil, status.Errorf(codes.Internal, "Invalid workspace ID: %v", err)
	}

	workspace, err := h.service.DeleteWorkspace(userID, workspaceID)
	if err != nil {
		switch err := err.(type) {
		case *customErrors.DatabaseError:
			h.log.Error("Database error while deleting workspace", "error", err)
			return nil, status.Errorf(codes.Internal, "Failed to delete workspace: %v", err)
		case *customErrors.NotFoundError:
			h.log.Error("Workspace not found", "error", err)
			return nil, status.Errorf(codes.NotFound, "Workspace not found: %v", err)
		default:
			h.log.Error("Unexpected error while deleting workspace", "error", err)
			return nil, status.Errorf(codes.Internal, "Unexpected error: %v", err)
		}
	}

	deleteResponse := &pb.DeleteWorkspaceResponse{
		Workspace: &pb.Workspace{
			Id:        workspace.ID.String(),
			Name:      workspace.Name,
			CreatedAt: format.TimeAsString(workspace.CreatedAt),
			UpdatedAt: format.TimeAsString(workspace.UpdatedAt),
		},
	}

	h.log.Info("Workspace deleted successfully", "workspace_id", workspace.ID.String(), "workspace_name", workspace.Name)

	return deleteResponse, nil
}
