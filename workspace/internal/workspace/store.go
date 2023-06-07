package workspace

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-task/core/pkg/logger"
	"go-task/core/pkg/models"
)

type (
	store struct {
		db  *pgxpool.Pool
		log logger.Logger
	}

	Store interface {
		CreateWorkspace(ctx context.Context, tx pgx.Tx, workspace *models.Workspace) error
		CreateWorkspaceUser(ctx context.Context, tx pgx.Tx, workspaceUser *models.WorkspaceUser) error
		BeginTx(ctx context.Context) (pgx.Tx, error)
		RollbackTx(ctx context.Context, tx pgx.Tx) error
	}
)

var _ Store = (*store)(nil)

func NewStore(connStr string, log logger.Logger) (Store, error) {
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &store{db: db, log: log}, nil
}

func (s *store) BeginTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		s.log.Error("could not begin transaction", "error", err)
		return nil, err
	}
	return tx, nil
}

func (s *store) RollbackTx(ctx context.Context, tx pgx.Tx) error {
	err := tx.Rollback(ctx)
	if err != nil {
		s.log.Error("could not rollback transaction", "error", err)
		return err
	}
	return nil
}

func (s *store) CreateWorkspace(ctx context.Context, tx pgx.Tx, workspace *models.Workspace) error {
	sqlStatement := `
		INSERT INTO workspaces (id, name)
		VALUES ($1, $2)
	`

	_, err := tx.Exec(ctx, sqlStatement, workspace.ID, workspace.Name)
	if err != nil {
		s.log.Error("failed to create workspace", "error", err)
		return err
	}

	return nil
}

func (s *store) CreateWorkspaceUser(ctx context.Context, tx pgx.Tx, workspaceUser *models.WorkspaceUser) error {
	sqlStatement := `
		INSERT INTO workspace_users (id, workspace_id, user_id, status)
		VALUES ($1, $2, $3, $4)
	`

	_, err := tx.Exec(ctx, sqlStatement, workspaceUser.ID, workspaceUser.WorkspaceID, workspaceUser.UserID, workspaceUser.Status)
	if err != nil {
		s.log.Error("failed to create workspace user", "error", err)
		return err
	}

	return nil
}
