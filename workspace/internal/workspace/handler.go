package workspace

import (
	"context"
	"fmt"
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

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	userId := ctx.Value("userId").(string)

	fmt.Println(userId)

	user, err := h.service.Create(req.GetName())
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
