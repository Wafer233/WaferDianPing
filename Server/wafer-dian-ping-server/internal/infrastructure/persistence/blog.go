package persistence

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"gorm.io/gorm"
)

type DefaultBlogRepository struct {
	db *gorm.DB
}

func (repo *DefaultBlogRepository) LikeById(ctx context.Context, id int64) error {

	err := repo.db.WithContext(ctx).Model(domain.Blog{}).
		Where("id = ?", id).
		Update("liked", gorm.Expr("liked + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultBlogRepository) DislikeById(ctx context.Context, id int64) error {
	err := repo.db.WithContext(ctx).Model(domain.Blog{}).
		Where("id = ?", id).
		Update("liked", gorm.Expr("liked - ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultBlogRepository) FindById(ctx context.Context, id int64) (*domain.Blog, error) {

	blog := domain.Blog{}
	err := repo.db.WithContext(ctx).Model(&domain.Blog{}).
		Where("id = ?", id).
		First(&blog).Error

	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (repo *DefaultBlogRepository) FindPageHot(ctx context.Context, page int, pageSize int) ([]*domain.Blog, error) {

	blogs := make([]*domain.Blog, 0)
	offset := (page - 1) * pageSize
	limit := pageSize

	err := repo.db.WithContext(ctx).Model(&domain.Blog{}).
		Offset(offset).Limit(limit).Find(&blogs).Order("liked desc").Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (repo *DefaultBlogRepository) Create(ctx context.Context, blog *domain.Blog) (int64, error) {

	err := repo.db.WithContext(ctx).Model(&domain.Blog{}).Create(blog).Error
	if err != nil {
		return 0, err
	}

	return blog.Id, nil

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
