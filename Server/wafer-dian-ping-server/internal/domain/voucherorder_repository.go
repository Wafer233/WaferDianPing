package domain

import "context"

type VoucherOrderRepository interface {
	Create(context.Context, *VoucherOrder) (int64, error)
	Exists(context.Context, int64, int64) (bool, error)
}
