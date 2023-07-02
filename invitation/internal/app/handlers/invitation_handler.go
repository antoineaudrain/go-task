package handlers

import (
	"context"
	pb "github.com/antoineaudrain/go-task/invitation/api"
	"github.com/antoineaudrain/go-task/invitation/internal/app/services"
	"github.com/antoineaudrain/go-task/invitation/internal/domain/invitation"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InvitationHandler struct {
	*pb.UnimplementedInvitationServiceServer
	invitationService *services.InvitationService
	logger            *zap.Logger
}

func NewInvitationHandler(logger *zap.Logger, invitationRepo invitation.Repository) *InvitationHandler {
	invitationService := services.NewInvitationService(invitationRepo)

	return &InvitationHandler{
		invitationService: invitationService,
		logger:            logger,
	}
}

func (h *InvitationHandler) SendWorkspaceInvitation(ctx context.Context, req *pb.SendWorkspaceInvitationRequest) (*pb.SendWorkspaceInvitationResponse, error) {
	authenticatedUserID := ctx.Value("authenticatedUserID").(*uuid.UUID)
	if authenticatedUserID == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	email := req.GetEmail()
	workspaceID, err := uuid.Parse(req.GetWorkspaceId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid workspace ID")
	}

	_, err = h.invitationService.CreateInvitation(ctx, email, workspaceID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create invitation")
	}

	return &pb.SendWorkspaceInvitationResponse{
		Success: true,
	}, nil
}
