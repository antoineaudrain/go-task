package member

import (
	"github.com/google/uuid"
	"time"
)

type Role string

const (
	RoleOwner  Role = "owner"
	RoleMember Role = "member"
)

type Member struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	WorkspaceID  uuid.UUID `gorm:"type:uuid"`
	UserID       uuid.UUID `gorm:"type:uuid"`
	Role         Role      `gorm:"type:varchar(20)"`
	InvitationId uuid.UUID `gorm:"type:uuid"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	//DeletedAt    *time.Time
}

func NewMember(workspaceID, userID uuid.UUID, role Role) *Member {
	return &Member{
		ID:          uuid.New(),
		WorkspaceID: workspaceID,
		UserID:      userID,
		Role:        role,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
