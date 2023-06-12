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
		CreateWorkspaceWithUser(userID uuid.UUID, name string) (*models.Workspace, error)
		GetWorkspace(userID, workspaceID uuid.UUID) (*models.Workspace, error)
		ListWorkspaces(userID uuid.UUID) ([]*models.Workspace, error)
		UpdateWorkspace(userID, workspaceID uuid.UUID, name string) (*models.Workspace, error)
		DeleteWorkspace(userID, workspaceID uuid.UUID) (*models.Workspace, error)
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

func (s *service) CreateWorkspaceWithUser(userID uuid.UUID, name string) (*models.Workspace, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", userID)

	tx, err := s.store.BeginTx(ctx)
	if err != nil {
		s.log.Error("failed to begin transaction", "error", err)
		return nil, customErrors.NewDatabaseError("failed to begin transaction", err)
	}

	workspace, err := s.store.CreateWorkspace(ctx, tx, name)
	if err != nil {
		_ = tx.Rollback(ctx)
		s.log.Error("failed to create workspace", "error", err)
		return nil, customErrors.NewDatabaseError("failed to create workspace", err)
	}

	_, err = s.store.CreateWorkspaceUser(ctx, tx, workspace.ID)
	if err != nil {
		_ = tx.Rollback(ctx)
		s.log.Error("failed to create workspace user", "error", err)
		return nil, customErrors.NewDatabaseError("failed to create workspace user", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		s.log.Error("failed to commit transaction", "error", err)
		return nil, customErrors.NewDatabaseError("failed to commit transaction", err)
	}

	return workspace, nil
}

func (s *service) GetWorkspace(userID, workspaceID uuid.UUID) (*models.Workspace, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", userID)

	workspace, err := s.store.GetWorkspace(ctx, workspaceID)
	if err != nil {
		s.log.Error("failed to get workspaces", "error", err)
		return nil, customErrors.NewDatabaseError("failed to get workspaces", err)
	}

	return workspace, nil
}

func (s *service) ListWorkspaces(userID uuid.UUID) ([]*models.Workspace, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", userID)

	workspaces, err := s.store.ListWorkspaces(ctx)
	if err != nil {
		s.log.Error("failed to list workspaces", "error", err)
		return nil, customErrors.NewDatabaseError("failed to list workspaces", err)
	}

	return workspaces, nil
}

func (s *service) UpdateWorkspace(userID, workspaceID uuid.UUID, name string) (*models.Workspace, error) {
	ctx := context.Background()

	tx, err := s.store.BeginTx(ctx)
	if err != nil {
		return nil, customErrors.NewDatabaseError("Failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	workspace, err := s.store.UpdateWorkspace(ctx, tx, workspaceID, userID, name)
	if err != nil {
		s.log.Error("Failed to update workspace", "error", err)
		return nil, customErrors.NewDatabaseError("failed to update workspace", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, customErrors.NewDatabaseError("Failed to commit transaction", err)
	}

	return workspace, nil
}

func (s *service) DeleteWorkspace(userID, workspaceID uuid.UUID) (*models.Workspace, error) {
	ctx := context.Background()

	tx, err := s.store.BeginTx(ctx)
	if err != nil {
		return nil, customErrors.NewDatabaseError("Failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	workspace, err := s.store.DeleteWorkspace(ctx, tx, workspaceID, userID)
	if err != nil {
		s.log.Error("Failed to delete workspace", "error", err)
		return nil, customErrors.NewDatabaseError("failed to delete workspace", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, customErrors.NewDatabaseError("Failed to commit transaction", err)
	}

	return workspace, nil
}
