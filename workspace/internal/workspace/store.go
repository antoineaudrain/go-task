package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	customErrors "go-task/core/pkg/errors"
	"go-task/core/pkg/logger"
	"go-task/core/pkg/models"
)

type (
	store struct {
		db  *pgxpool.Pool
		log logger.Logger
	}

	Store interface {
		BeginTx(ctx context.Context) (pgx.Tx, error)
		RollbackTx(ctx context.Context, tx pgx.Tx) error
		CreateWorkspace(ctx context.Context, tx pgx.Tx, name string) (*models.Workspace, error)
		GetWorkspace(ctx context.Context, workspaceID uuid.UUID) (*models.Workspace, error)
		CreateWorkspaceUser(ctx context.Context, tx pgx.Tx, workspaceID uuid.UUID) (*models.WorkspaceUser, error)
		GetWorkspaceUser(ctx context.Context, workspaceID, userID uuid.UUID) (*models.WorkspaceUser, error)
		ListWorkspaces(ctx context.Context) ([]*models.Workspace, error)
		UpdateWorkspace(ctx context.Context, tx pgx.Tx, workspaceID, userID uuid.UUID, name string) (*models.Workspace, error)
		DeleteWorkspace(ctx context.Context, tx pgx.Tx, workspaceID, userID uuid.UUID) (*models.Workspace, error)
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

func (s *store) CreateWorkspace(ctx context.Context, tx pgx.Tx, name string) (*models.Workspace, error) {
	workspaceID := uuid.New()

	_, err := tx.Exec(ctx, "INSERT INTO workspaces (id, name) VALUES ($1, $2)", workspaceID, name)
	if err != nil {
		return nil, err
	}

	workspace, err := s.getWorkspaceInTransaction(ctx, tx, workspaceID)
	if err != nil {
		return nil, err
	}

	return workspace, nil
}

func (s *store) getWorkspaceInTransaction(ctx context.Context, tx pgx.Tx, workspaceID uuid.UUID) (*models.Workspace, error) {
	var workspace models.Workspace
	sqlStatement := `
		SELECT id, name, created_at, updated_at FROM workspaces WHERE id = $1;
	`
	err := tx.QueryRow(ctx, sqlStatement, workspaceID).Scan(&workspace.ID, &workspace.Name, &workspace.CreatedAt, &workspace.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, customErrors.NewNotFoundError("workspace not found", err)
		}
		return nil, customErrors.NewDatabaseError("failed to get workspace", err)
	}

	return &workspace, nil
}

func (s *store) GetWorkspace(ctx context.Context, workspaceID uuid.UUID) (*models.Workspace, error) {
	userID := ctx.Value("userID").(uuid.UUID)
	var workspace models.Workspace
	sqlStatement := `
		SELECT w.id, w.name, w.created_at, w.updated_at
		FROM workspaces w
		JOIN workspace_users u ON w.id = u.workspace_id
		WHERE w.id = $1 AND u.user_id = $2;
	`
	err := s.db.QueryRow(ctx, sqlStatement, workspaceID, userID).Scan(&workspace.ID, &workspace.Name, &workspace.CreatedAt, &workspace.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, customErrors.NewNotFoundError("workspace not found", err)
		}
		return nil, customErrors.NewDatabaseError("failed to get workspace", err)
	}

	return &workspace, nil
}

func (s *store) CreateWorkspaceUser(ctx context.Context, tx pgx.Tx, workspaceID uuid.UUID) (*models.WorkspaceUser, error) {
	userID := ctx.Value("userID").(uuid.UUID)

	_, err := tx.Exec(ctx, "INSERT INTO workspace_users (workspace_id, user_id, status) VALUES ($1, $2, $3)", workspaceID, userID, models.StatusPending)
	if err != nil {
		return nil, err
	}

	workspaceUser, err := s.getWorkspaceUserInTransaction(ctx, tx, workspaceID, userID)
	if err != nil {
		return nil, err
	}

	return workspaceUser, nil
}

func (s *store) getWorkspaceUserInTransaction(ctx context.Context, tx pgx.Tx, workspaceID, userID uuid.UUID) (*models.WorkspaceUser, error) {
	var workspaceUser models.WorkspaceUser
	sqlStatement := `
		SELECT workspace_users.id, workspace_users.workspace_id, workspace_users.user_id, workspace_users.status, workspace_users.created_at, workspace_users.updated_at
		FROM workspace_users
		WHERE workspace_id = $1 AND user_id = $2;
	`
	err := tx.QueryRow(ctx, sqlStatement, workspaceID, userID).Scan(&workspaceUser.ID, &workspaceUser.WorkspaceID, &workspaceUser.UserID, &workspaceUser.Status, &workspaceUser.CreatedAt, &workspaceUser.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, customErrors.NewNotFoundError("workspace user not found", err)
		}
		return nil, customErrors.NewDatabaseError("failed to get workspace user", err)
	}

	return &workspaceUser, nil
}

func (s *store) GetWorkspaceUser(ctx context.Context, workspaceID, userID uuid.UUID) (*models.WorkspaceUser, error) {
	var workspaceUser models.WorkspaceUser
	sqlStatement := `
		SELECT workspace_users.id, workspace_users.workspace_id, workspace_users.user_id, workspace_users.status, workspace_users.created_at, workspace_users.updated_at
		FROM workspace_users
		WHERE workspace_id = $1 AND user_id = $2;
	`
	err := s.db.QueryRow(ctx, sqlStatement, workspaceID, userID).Scan(&workspaceUser.ID, &workspaceUser.WorkspaceID, &workspaceUser.UserID, &workspaceUser.Status, &workspaceUser.CreatedAt, &workspaceUser.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, customErrors.NewNotFoundError("workspace user not found", err)
		}
		return nil, customErrors.NewDatabaseError("failed to get workspace user", err)
	}

	return &workspaceUser, nil
}

func (s *store) ListWorkspaces(ctx context.Context) ([]*models.Workspace, error) {
	userID := ctx.Value("userID").(uuid.UUID)
	sqlStatement := `
		SELECT w.id, w.name, w.created_at, w.updated_at
		FROM workspaces w
		JOIN workspace_users u ON w.id = u.workspace_id
		WHERE u.user_id = $1
	`
	rows, err := s.db.Query(ctx, sqlStatement, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workspaces []*models.Workspace

	for rows.Next() {
		var workspace models.Workspace
		err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.CreatedAt, &workspace.UpdatedAt)
		if err != nil {
			return nil, err
		}

		workspaces = append(workspaces, &workspace)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (s *store) UpdateWorkspace(ctx context.Context, tx pgx.Tx, workspaceID, userID uuid.UUID, name string) (*models.Workspace, error) {
	updateStatement := `
		UPDATE workspaces
		SET name = $1, updated_at = NOW()
		WHERE id = $2
		AND id IN (
			SELECT workspace_id
			FROM workspace_users
			WHERE user_id = $3
		);
	`
	_, err := tx.Exec(ctx, updateStatement, name, workspaceID, userID)
	if err != nil {
		return nil, err
	}

	workspace, err := s.getWorkspaceInTransaction(ctx, tx, workspaceID)
	if err != nil {
		return nil, err
	}

	return workspace, nil
}

func (s *store) DeleteWorkspace(ctx context.Context, tx pgx.Tx, workspaceID, userID uuid.UUID) (*models.Workspace, error) {
	deleteStatement := `
		DELETE FROM workspaces
		WHERE id = $1
		AND id IN (
			SELECT workspace_id
			FROM workspace_users
			WHERE user_id = $2
		);
	`

	workspace, err := s.getWorkspaceInTransaction(ctx, tx, workspaceID)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, deleteStatement, workspaceID, userID)
	if err != nil {
		return nil, err
	}

	return workspace, nil
}
