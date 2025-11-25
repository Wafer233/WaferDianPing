package domain

import "context"

type VoucherRepository interface {
	CreateVoucher(context.Context, *Voucher) (int64, error)
	CreateSeckillVoucher(context.Context, *SecKillVoucher) (int64, error)
	FindByShopId(context.Context, int64) ([]*Voucher, []int64, error)
	FindSecKillByIds(context.Context, []int64) (map[int64]*SecKillVoucher, error)
}
