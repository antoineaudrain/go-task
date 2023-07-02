package invitation

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, workspace *Invitation) error
	FindByEmailAndCode(ctx context.Context, email, code string) (*Invitation, error)
	Delete(ctx context.Context, workspace Invitation) error
}
