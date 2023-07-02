package invitation

import (
	"github.com/google/uuid"
	"time"
)

type Invitation struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email       string
	Code        string
	WorkspaceID uuid.UUID `gorm:"type:uuid"`
	LastSentAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	//DeletedAt *time.Time
}

func NewInvitation(email, code string, workspaceID uuid.UUID) *Invitation {
	return &Invitation{
		ID:          uuid.New(),
		Email:       email,
		Code:        code,
		WorkspaceID: workspaceID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
