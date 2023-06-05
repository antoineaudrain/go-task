package handler

import (
	"context"
	pb "go-task/workspace/api"
	"go-task/workspace/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedWorkspaceServiceServer
	userService *service.WorkspaceService
}

func NewHandler() *Handler {
	return &Handler{
		userService: service.NewWorkspaceService(),
	}
}

func (h *Handler) Register(s *grpc.Server) {
	pb.RegisterWorkspaceServiceServer(s, h)
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	user, err := h.userService.Create(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	return &pb.CreateResponse{
		Workspace: &pb.Workspace{
			Id:          user.ID.String(),
			Name:        user.Name,
			Description: user.Description,
		},
	}, nil
}
