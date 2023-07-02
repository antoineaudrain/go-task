package persistence

import (
	"context"
	"errors"
	"github.com/antoineaudrain/go-task/workspace/internal/domain/member"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{
		db: db,
	}
}

func (r *MemberRepository) Save(ctx context.Context, member *member.Member) error {
	if err := r.db.WithContext(ctx).Create(member).Error; err != nil {
		// Handle the error appropriately
		return err
	}
	return nil
}

func (r *MemberRepository) FindByID(ctx context.Context, id uuid.UUID) (*member.Member, error) {
	m := &member.Member{}
	result := r.db.WithContext(ctx).Where("id = ?", id).First(m)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Handle the case when no m with the given ID is found
			return nil, nil
		}
		return nil, result.Error
	}
	return m, nil
}

func (r *MemberRepository) ListByWorkspaceID(ctx context.Context, workspaceID uuid.UUID) ([]*member.Member, error) {
	var members []*member.Member
	result := r.db.WithContext(ctx).Where("workspace_id = ?", workspaceID).Find(&members)
	if result.Error != nil {
		// Handle the error appropriately
		return nil, result.Error
	}
	return members, nil
}

func (r *MemberRepository) Delete(ctx context.Context, memberID uuid.UUID) error {
	m := &member.Member{ID: memberID}
	if err := r.db.WithContext(ctx).Delete(m).Error; err != nil {
		return err
	}

	return nil
}
