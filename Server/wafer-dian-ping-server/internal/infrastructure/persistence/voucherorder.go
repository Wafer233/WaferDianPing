package persistence

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"gorm.io/gorm"
)

type DefaultVoucherOrderRepository struct {
	db *gorm.DB
}

func (repo *DefaultVoucherOrderRepository) Exists(ctx context.Context,
	userId int64, voucherId int64) (bool, error) {

	var count int64

	err := repo.db.WithContext(ctx).Model(&domain.VoucherOrder{}).
		Where("user_id = ? AND voucher_id = ?", userId, voucherId).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *DefaultVoucherOrderRepository) Create(ctx context.Context, order *domain.VoucherOrder) (int64, error) {

	err := repo.db.WithContext(ctx).Model(&domain.VoucherOrder{}).Create(order).Error
	if err != nil {
		return 0, err
	}
	return order.Id, nil
}

func NewDefaultVoucherOrderRepository(db *gorm.DB) domain.VoucherOrderRepository {
	return &DefaultVoucherOrderRepository{db: db}
}
