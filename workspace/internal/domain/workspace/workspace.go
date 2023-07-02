package workspace

import (
	"github.com/google/uuid"
	"time"
)

type Workspace struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	OwnerID   uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt *time.Time
}

func NewWorkspace(name string, ownerID uuid.UUID) *Workspace {
	return &Workspace{
		ID:        uuid.New(),
		Name:      name,
		OwnerID:   ownerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
