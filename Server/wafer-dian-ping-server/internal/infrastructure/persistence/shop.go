package persistence

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"gorm.io/gorm"
)

type DefaultShopRepository struct {
	db *gorm.DB
}

func (repo *DefaultShopRepository) FindById(ctx context.Context, id int64) (*domain.Shop, error) {

	shop := domain.Shop{}
	err := repo.db.WithContext(ctx).Model(&domain.Shop{}).
		Where("id = ?", id).First(&shop).Error
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

func (repo DefaultShopRepository) FindPage(ctx context.Context, typeId int64, page int, pageSize int) ([]*domain.Shop, error) {

	offset := (page - 1) * pageSize
	shops := make([]*domain.Shop, 0)
	err := repo.db.WithContext(ctx).Model(&domain.Shop{}).
		Where("type_id = ?", typeId).
		Limit(pageSize).
		Offset(offset).
		Find(&shops).Error
	if err != nil {
		return nil, err
	}
	return shops, nil

}

func NewDefaultShopRepository(db *gorm.DB) *DefaultShopRepository {
	return &DefaultShopRepository{db: db}
}
