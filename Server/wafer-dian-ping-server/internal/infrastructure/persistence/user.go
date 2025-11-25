package persistence

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"gorm.io/gorm"
)

type DefaultUserRepository struct {
	db *gorm.DB
}

func (repo *DefaultUserRepository) FindUserByIds(ctx context.Context, ids []int64) (map[int64]*domain.User, error) {

	users := make([]*domain.User, 0)

	err := repo.db.WithContext(ctx).
		Model(domain.User{}).Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		return nil, err
	}

	mapping := make(map[int64]*domain.User)
	for _, user := range users {
		mapping[user.Id] = user
	}
	return mapping, nil
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
