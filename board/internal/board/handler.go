package board

import (
	pb "go-task/board/pkg/proto"
	"go-task/core/pkg/logger"
	"google.golang.org/grpc"
)

type Handler struct {
	pb.UnimplementedBoardServiceServer
	log logger.Logger
}

func NewHandler(log logger.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) Register(s *grpc.Server) {
	pb.RegisterBoardServiceServer(s, h)
}
