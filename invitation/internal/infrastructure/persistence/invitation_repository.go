package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/antoineaudrain/go-task/invitation/internal/domain/invitation"
	"gorm.io/gorm"
)

type InvitationRepository struct {
	db *gorm.DB
}

func NewInvitationRepository(db *gorm.DB) *InvitationRepository {
	return &InvitationRepository{
		db: db,
	}
}

func (r *InvitationRepository) Save(ctx context.Context, invitation *invitation.Invitation) error {
	if err := r.db.WithContext(ctx).Save(invitation).Error; err != nil {
		// Handle the error appropriately
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *InvitationRepository) FindByEmailAndCode(ctx context.Context, email, code string) (*invitation.Invitation, error) {
	w := &invitation.Invitation{}
	if err := r.db.WithContext(ctx).Where("email = ? AND code = ?", email, code).First(w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case when no w with the given ID is found
			return nil, nil
		}
		return nil, err
	}

	return w, nil
}

func (r *InvitationRepository) Delete(ctx context.Context, invitation invitation.Invitation) error {
	if err := r.db.WithContext(ctx).Delete(invitation).Error; err != nil {
		return err
	}

	return nil
}
