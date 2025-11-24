package domain

import "context"

type ShopRepository interface {
	FindById(context.Context, int64) (*Shop, error)
	FindPage(context.Context, int64, int, int) ([]*Shop, error)
}
