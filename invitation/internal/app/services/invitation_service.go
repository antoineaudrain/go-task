package services

import (
	"context"
	"github.com/antoineaudrain/go-task/invitation/internal/domain/invitation"
	"github.com/antoineaudrain/go-task/invitation/internal/shared/invite_code"
	"github.com/google/uuid"
)

type InvitationService struct {
	invitationRepo invitation.Repository
}

func NewInvitationService(invitationRepo invitation.Repository) *InvitationService {
	return &InvitationService{
		invitationRepo: invitationRepo,
	}
}

func (s *InvitationService) CreateInvitation(ctx context.Context, email string, workspaceID uuid.UUID) (*invitation.Invitation, error) {
	code := invite_code.GenerateInvitationCode()
	i := invitation.NewInvitation(email, code, workspaceID)

	err := s.invitationRepo.Save(ctx, i)
	if err != nil {
		// Handle the error appropriately
		return nil, err
	}

	return i, nil
}
