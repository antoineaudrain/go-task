package workspace

import (
	"context"
	"github.com/google/uuid"
	customErrors "go-task/core/pkg/errors"
	"go-task/core/pkg/logger"
	"go-task/core/pkg/models"
	"os"
)

type (
	service struct {
		store Store
		log   logger.Logger
	}

	Service interface {
		CreateWorkspaceWithUser(name string, userId uuid.UUID) (*models.Workspace, error)
	}
)

var _ Service = (*service)(nil)

func NewService(log logger.Logger) Service {
	s, err := NewStore(os.Getenv("DATABASE_URL"), log)
	if err != nil {
		log.Error("Failed to create user store", "error", err)
		os.Exit(1)
	}

	return &service{
		store: s,
		log:   log,
	}
}

func (s *service) CreateWorkspaceWithUser(name string, userId uuid.UUID) (*models.Workspace, error) {
	ctx := context.Background()
	tx, err := s.store.BeginTx(ctx)
	if err != nil {
		return nil, customErrors.NewDatabaseError("failed to start transaction", err)
	}

	workspace := &models.Workspace{
		ID:   uuid.New(),
		Name: name,
	}

	if err := s.store.CreateWorkspace(ctx, tx, workspace); err != nil {
		_ = s.store.RollbackTx(ctx, tx)
		s.log.Error("failed to create workspace", "error", err)
		return nil, customErrors.NewDatabaseError("failed to create workspace", err)
	}

	workspaceUser := &models.WorkspaceUser{
		ID:          uuid.New(),
		WorkspaceID: workspace.ID,
		UserID:      userId,
		Status:      models.WorkspaceStatusPending,
	}

	if err := s.store.CreateWorkspaceUser(ctx, tx, workspaceUser); err != nil {
		_ = s.store.RollbackTx(ctx, tx)
		s.log.Error("failed to create workspace user", "error", err)
		return nil, customErrors.NewDatabaseError("failed to create workspace user", err)
	}

	if err := tx.Commit(ctx); err != nil {
		s.log.Error("failed to commit transaction", "error", err)
		return nil, customErrors.NewDatabaseError("failed to commit transaction", err)
	}

	return workspace, nil
}
