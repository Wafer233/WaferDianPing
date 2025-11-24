package persistence

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"gorm.io/gorm"
)

type DefaultShopTypeRepository struct {
	db *gorm.DB
}

func NewDefaultShopTypeRepository(db *gorm.DB) domain.ShopTypeRepository {
	return &DefaultShopTypeRepository{db: db}
}

func (repo *DefaultShopTypeRepository) FindShopTypeList(ctx context.Context) ([]*domain.ShopType, error) {

	types := make([]*domain.ShopType, 0)

	err := repo.db.WithContext(ctx).Model(&domain.ShopType{}).Find(&types).Error

	return types, err
}
