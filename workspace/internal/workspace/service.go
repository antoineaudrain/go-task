package workspace

import (
	"github.com/google/uuid"
	"go-task/core/pkg/models"
	"log"
	"os"
)

type (
	service struct {
		store Store
	}

	Service interface {
		Create(name string) (*models.Workspace, error)
	}
)

var _ Service = (*service)(nil)

func NewService() Service {
	s, err := NewStore(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to create user store: %v", err)
	}

	return &service{
		store: s,
	}
}

func (s *service) Create(name string) (*models.Workspace, error) {
	workspace := &models.Workspace{
		ID:   uuid.New(),
		Name: name,
	}

	if err := s.store.CreateWorkspace(workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}
