package services

import (
	"context"
	"github.com/antoineaudrain/go-task/workspace/internal/domain/member"
	"github.com/antoineaudrain/go-task/workspace/internal/domain/workspace"
	customErrors "github.com/antoineaudrain/go-task/workspace/internal/shared/errors"
	"github.com/google/uuid"
)

type MemberService struct {
	workspaceRepo workspace.Repository
	memberRepo    member.Repository
}

func NewMemberService(workspaceRepo workspace.Repository, memberRepo member.Repository) *MemberService {
	return &MemberService{
		workspaceRepo: workspaceRepo,
		memberRepo:    memberRepo,
	}
}

func (s *MemberService) ListByWorkspaceID(ctx context.Context, workspaceID uuid.UUID) ([]*member.Member, error) {
	_, err := s.workspaceRepo.FindByID(ctx, workspaceID)
	if err != nil {
		return nil, err
	}

	members, err := s.memberRepo.ListByWorkspaceID(ctx, workspaceID)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (s *MemberService) InviteWorkspaceMember(ctx context.Context, workspaceID uuid.UUID, email string, inviterID uuid.UUID) (*member.Member, error) {
	_, err := s.workspaceRepo.FindByID(ctx, workspaceID)
	if err != nil {
		return nil, err
	}

	m := &member.Member{
		WorkspaceID:  workspaceID,
		Role:         member.RoleMember,
		InvitationId: uuid.New(),
	}

	err = s.memberRepo.Save(ctx, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *MemberService) RemoveWorkspaceMember(ctx context.Context, workspaceID, memberID, removerID uuid.UUID) error {
	_, err := s.workspaceRepo.FindByID(ctx, workspaceID)
	if err != nil {
		return err
	}

	_, err = s.memberRepo.FindByID(ctx, memberID)
	if err != nil {
		return err
	}

	isOwner, err := s.workspaceRepo.IsOwner(ctx, workspaceID, removerID)
	if err != nil {
		return err
	}
	if !isOwner {
		return customErrors.ErrUnauthorized
	}

	err = s.memberRepo.Delete(ctx, memberID)
	if err != nil {
		return err
	}

	return nil
}
