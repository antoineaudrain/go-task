package services

import (
	"context"
	"fmt"
	"github.com/antoineaudrain/go-task/workspace/internal/domain/member"
	"github.com/antoineaudrain/go-task/workspace/internal/domain/workspace"
	"github.com/google/uuid"
)

type WorkspaceService struct {
	workspaceRepo workspace.Repository
	memberRepo    member.Repository
}

func NewWorkspaceService(workspaceRepo workspace.Repository, memberRepo member.Repository) *WorkspaceService {
	return &WorkspaceService{
		workspaceRepo: workspaceRepo,
		memberRepo:    memberRepo,
	}
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, name string, ownerID uuid.UUID) (*workspace.Workspace, error) {
	w := workspace.NewWorkspace(name, ownerID)

	err := s.workspaceRepo.Save(ctx, w)
	if err != nil {
		// Handle the error appropriately
		return nil, err
	}

	workspaceMember := member.NewMember(w.ID, ownerID, member.RoleOwner)

	err = s.memberRepo.Save(ctx, workspaceMember)
	if err != nil {
		// Handle the error appropriately
		return nil, err
	}

	return w, nil
}

func (s *WorkspaceService) GetWorkspace(ctx context.Context, workspaceID uuid.UUID) (*workspace.Workspace, error) {
	w, err := s.workspaceRepo.FindByID(ctx, workspaceID)
	if err != nil {
		// Handle the error appropriately
		return nil, err
	}

	return w, nil
}

func (s *WorkspaceService) UpdateWorkspace(ctx context.Context, workspaceID uuid.UUID, name string, ownerID uuid.UUID) (*workspace.Workspace, error) {
	w, err := s.workspaceRepo.FindByID(ctx, workspaceID)
	if err != nil {
		// Handle the error appropriately
		return nil, err
	}

	if w.OwnerID != ownerID {
		// Handle the error appropriately
		return nil, err
	}

	w.Name = name

	fmt.Println(w)

	err = s.workspaceRepo.Save(ctx, w)
	if err != nil {
		// Handle the error appropriately
		return nil, err
	}

	return w, nil
}

func (s *WorkspaceService) DeleteWorkspace(ctx context.Context, workspaceID uuid.UUID, userID uuid.UUID) error {
	w, err := s.workspaceRepo.FindByID(ctx, workspaceID)
	if err != nil {
		// Handle the error appropriately
		return err
	}

	if w.OwnerID != userID {
		// Handle the error appropriately
		return err
	}

	err = s.workspaceRepo.Delete(ctx, *w)
	if err != nil {
		// Handle the error appropriately
		return err
	}

	return nil
}
