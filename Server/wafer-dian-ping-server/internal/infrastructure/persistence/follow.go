package persistence

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"gorm.io/gorm"
)

type DefaultFollowRepository struct {
	db *gorm.DB
}

func (repo *DefaultFollowRepository) AddFollow(ctx context.Context, follow *domain.Follow) error {

	err := repo.db.WithContext(ctx).
		Model(&domain.Follow{}).
		Create(follow).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultFollowRepository) RemoveFollow(ctx context.Context, follow *domain.Follow) error {

	err := repo.db.WithContext(ctx).
		Where("user_id = ? AND follow_user_id = ?", follow.UserId, follow.FollowUserId).
		Delete(&domain.Follow{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultFollowRepository) FindFollow(ctx context.Context, userId int64, curId int64) (bool, error) {

	total := int64(0)

	err := repo.db.WithContext(ctx).Model(&domain.Follow{}).
		Where("user_id = ? and follow_user_id = ?", userId, curId).
		Count(&total).Error
	if err != nil {
		return false, err
	}
	return total > 0, nil
}

func NewDefaultFollowRepository(db *gorm.DB) domain.FollowRepository {
	return &DefaultFollowRepository{db: db}
}
