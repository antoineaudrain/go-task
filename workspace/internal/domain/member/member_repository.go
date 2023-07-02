package member

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, member *Member) error
	FindByID(ctx context.Context, id uuid.UUID) (*Member, error)
	ListByWorkspaceID(ctx context.Context, workspaceID uuid.UUID) ([]*Member, error)
	Delete(ctx context.Context, memberID uuid.UUID) error
}
