package persistence

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"gorm.io/gorm"
)

type DefaultVoucherRepository struct {
	db *gorm.DB
}

func (repo *DefaultVoucherRepository) FindSecKillByIds(ctx context.Context,
	ids []int64) (map[int64]*domain.SecKillVoucher, error) {

	seckills := make([]*domain.SecKillVoucher, 0)

	err := repo.db.WithContext(ctx).Model(&domain.SecKillVoucher{}).
		Where("voucher_id in ?", ids).
		Find(&seckills).Error
	if err != nil {
		return nil, err
	}

	mapping := make(map[int64]*domain.SecKillVoucher)
	for _, seckill := range seckills {
		mapping[seckill.VoucherId] = seckill
	}
	return mapping, nil
}

func (repo *DefaultVoucherRepository) FindByShopId(ctx context.Context, shopId int64) ([]*domain.Voucher, []int64, error) {

	var vouchers []*domain.Voucher
	var ids []int64

	err := repo.db.WithContext(ctx).
		Where("shop_id = ?", shopId).
		Find(&vouchers).Error
	if err != nil {
		return nil, nil, err
	}

	// 再查 IDs
	err = repo.db.WithContext(ctx).
		Model(&domain.Voucher{}).
		Where("shop_id = ? AND type = ?", shopId, 1).
		Pluck("id", &ids).Error
	if err != nil {
		return nil, nil, err
	}

	return vouchers, ids, nil

}

func (repo *DefaultVoucherRepository) CreateVoucher(ctx context.Context,
	voucher *domain.Voucher) (int64, error) {

	err := repo.db.WithContext(ctx).Model(&domain.Voucher{}).
		Create(voucher).Error
	if err != nil {
		return 0, err
	}
	return voucher.Id, nil
}

func (repo DefaultVoucherRepository) CreateSeckillVoucher(ctx context.Context, voucher *domain.SecKillVoucher) (int64, error) {

	err := repo.db.WithContext(ctx).Model(&domain.SecKillVoucher{}).
		Create(voucher).Error
	if err != nil {
		return 0, err
	}
	return voucher.VoucherId, nil
}

func NewDefaultVoucherRepository(db *gorm.DB) domain.VoucherRepository {
	return &DefaultVoucherRepository{db: db}
}
