package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/antoineaudrain/go-task/workspace/internal/domain/workspace"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkspaceRepository struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) *WorkspaceRepository {
	return &WorkspaceRepository{
		db: db,
	}
}

func (r *WorkspaceRepository) Save(ctx context.Context, workspace *workspace.Workspace) error {
	if err := r.db.WithContext(ctx).Save(workspace).Error; err != nil {
		// Handle the error appropriately
		fmt.Println(err)
		return err
	}
	fmt.Println("NO ERROR")
	return nil
}

func (r *WorkspaceRepository) FindByID(ctx context.Context, id uuid.UUID) (*workspace.Workspace, error) {
	w := &workspace.Workspace{}
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case when no w with the given ID is found
			return nil, nil
		}
		return nil, err
	}

	return w, nil
}

func (r *WorkspaceRepository) Delete(ctx context.Context, workspace workspace.Workspace) error {
	if err := r.db.WithContext(ctx).Delete(workspace).Error; err != nil {
		return err
	}

	return nil
}

func (r *WorkspaceRepository) IsOwner(ctx context.Context, workspaceID, userID uuid.UUID) (bool, error) {
	w := &workspace.Workspace{}
	if err := r.db.WithContext(ctx).Where("id = ? AND owner_id = ?", workspaceID, userID).First(w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// The w or the owner was not found
			return false, nil
		}
		return false, err
	}

	return true, nil
}
