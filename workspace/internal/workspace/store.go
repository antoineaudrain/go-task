package workspace

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-task/core/pkg/models"
)

type (
	store struct {
		db *pgxpool.Pool
	}

	Store interface {
		CreateWorkspace(workspace *models.Workspace) error
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
