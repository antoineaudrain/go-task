package service

import (
	"github.com/google/uuid"
	"go-task/workspace/internal/store/postgres"
	"go-task/workspace/pkg/models"
	"log"
)

type WorkspaceService struct {
	workspaceStore *postgres.WorkspaceStore
}

type Service interface {
	Create(name string) (*models.Workspace, error)
}

var _ Service = (*WorkspaceService)(nil)

func NewWorkspaceService() *WorkspaceService {
	store, err := postgres.NewWorkspaceStore("postgresql://go_task:go_task@localhost:5433/workspace_db")
	if err != nil {
		log.Fatalf("Failed to create user store: %v", err)
	}

	return &WorkspaceService{
		workspaceStore: store,
	}
}

func (s *WorkspaceService) Create(name string) (*models.Workspace, error) {
	workspace := &models.Workspace{
		ID:   uuid.New(),
		Name: name,
	}

	if err := s.workspaceStore.CreateWorkspace(workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}
