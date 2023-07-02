package workspace

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, workspace *Workspace) error
	FindByID(ctx context.Context, id uuid.UUID) (*Workspace, error)
	IsOwner(ctx context.Context, workspaceID, userID uuid.UUID) (bool, error)
	Delete(ctx context.Context, workspace Workspace) error
}
