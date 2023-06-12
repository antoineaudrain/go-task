package models

import (
	"github.com/google/uuid"
	"time"
)

type WorkspaceStatus string

const (
	StatusPending  WorkspaceStatus = "pending"
	StatusSent     WorkspaceStatus = "sent"
	StatusAccepted WorkspaceStatus = "accepted"
	StatusRejected WorkspaceStatus = "rejected"
)

type Workspace struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WorkspaceUser struct {
	ID          uuid.UUID
	WorkspaceID uuid.UUID
	UserID      uuid.UUID
	Status      WorkspaceStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
