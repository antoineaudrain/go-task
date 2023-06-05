package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-task/workspace/pkg/models"
)

type WorkspaceStore struct {
	db *pgxpool.Pool
}

func NewWorkspaceStore(connStr string) (*WorkspaceStore, error) {
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &WorkspaceStore{db: db}, nil
}

func (s *WorkspaceStore) CreateWorkspace(workspace *models.Workspace) error {
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
