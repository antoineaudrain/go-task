package workspace

import (
	"github.com/google/uuid"
	customErrors "go-task/core/pkg/errors"
	"go-task/core/pkg/models"
	"log"
	"os"
)

type (
	service struct {
		store Store
	}

	Service interface {
		CreateWorkspaceWithUser(name string, userId uuid.UUID) (*models.Workspace, error)
	}
)

var _ Service = (*service)(nil)

func NewService() Service {
	s, err := NewStore(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to create user store: %v", err)
	}

	return &service{
		store: s,
	}
}

func (s *service) CreateWorkspaceWithUser(name string, userId uuid.UUID) (*models.Workspace, error) {
	workspace := &models.Workspace{
		ID:   uuid.New(),
		Name: name,
	}

	if err := s.store.CreateWorkspace(workspace); err != nil {
		return nil, customErrors.NewDatabaseError("failed to create workspace", err)
	}

	workspaceUser := &models.WorkspaceUser{
		ID:          uuid.New(),
		WorkspaceID: workspace.ID,
		UserID:      userId,
		Status:      models.WorkspaceStatusPending,
	}

	if err := s.store.CreateWorkspaceUser(workspaceUser); err != nil {
		return nil, customErrors.NewDatabaseError("failed to create workspace user", err)
	}

	return workspace, nil
}
