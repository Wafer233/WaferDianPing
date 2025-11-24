package persistence

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"gorm.io/gorm"
)

type DefaultUserRepository struct {
	db *gorm.DB
}

func (repo *DefaultUserRepository) FindInfoByUserId(ctx context.Context, userId int64) (*domain.UserInfo, error) {

	entity := &domain.UserInfo{}

	err := repo.db.WithContext(ctx).Model(&domain.UserInfo{}).
		Where("user_id = ?", userId).
		Find(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (repo *DefaultUserRepository) FindUserById(ctx context.Context, id int64) (*domain.User, error) {

	entity := &domain.User{}

	err := repo.db.WithContext(ctx).Model(&domain.User{}).
		Where("id = ?", id).
		First(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (repo *DefaultUserRepository) FindUserByPhone(ctx context.Context, phone string) (*domain.User, error) {

	entity := &domain.User{}

	err := repo.db.WithContext(ctx).Model(&domain.User{}).
		Where("phone = ?", phone).
		First(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func NewDefaultUserRepository(db *gorm.DB) domain.UserRepository {
	return &DefaultUserRepository{db: db}
}
