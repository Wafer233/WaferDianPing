package persistence

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"gorm.io/gorm"
)

type DefaultBlogRepository struct {
	db *gorm.DB
}

func (repo *DefaultBlogRepository) FindByUserId(ctx context.Context, id int64) ([]*domain.Blog, error) {

	blogs := make([]*domain.Blog, 0)

	err := repo.db.WithContext(ctx).Model(&domain.Blog{}).Where("user_id = ?", id).
		Find(&blogs).Error
	return blogs, err

}

func (repo *DefaultBlogRepository) FindByShopId(ctx context.Context, id int64) ([]*domain.Blog, error) {
	blogs := make([]*domain.Blog, 0)

	err := repo.db.WithContext(ctx).Model(&domain.Blog{}).Where("shop_id = ?", id).
		Find(&blogs).Error
	return blogs, err
}

func NewDefaultBlogRepository(db *gorm.DB) domain.BlogRepository {
	return &DefaultBlogRepository{db: db}
}
