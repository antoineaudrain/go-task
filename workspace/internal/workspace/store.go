package workspace

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-task/core/pkg/models"
)

type (
	store struct {
		db *pgxpool.Pool
	}

	Store interface {
		CreateWorkspace(workspace *models.Workspace) error
		CreateWorkspaceUser(workspaceUser *models.WorkspaceUser) error
	}
)

var _ Store = (*store)(nil)

func NewStore(connStr string) (Store, error) {
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &store{db: db}, nil
}

func (s *store) CreateWorkspace(workspace *models.Workspace) error {
	sqlStatement := `
		INSERT INTO workspaces (id, name)
		VALUES ($1, $2)
	`

	_, err := s.db.Exec(context.Background(), sqlStatement, workspace.ID, workspace.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) CreateWorkspaceUser(workspaceUser *models.WorkspaceUser) error {
	sqlStatement := `
		INSERT INTO workspace_users (id, workspace_id, user_id, status)
		VALUES ($1, $2, $3, $4)
	`

	_, err := s.db.Exec(context.Background(), sqlStatement, workspaceUser.ID, workspaceUser.WorkspaceID, workspaceUser.UserID, workspaceUser.Status)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}
