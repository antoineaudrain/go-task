package models

import (
	"github.com/google/uuid"
	"time"
)

const (
	WorkspaceStatusPending  = "pending"
	WorkspaceStatusSent     = "sent"
	WorkspaceStatusAccepted = "accepted"
	WorkspaceStatusRejected = "rejected"
)

type Workspace struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type WorkspaceUser struct {
	ID          uuid.UUID
	WorkspaceID uuid.UUID
	UserID      uuid.UUID
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
