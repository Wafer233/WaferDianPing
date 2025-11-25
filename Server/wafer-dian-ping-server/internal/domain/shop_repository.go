package domain

import "context"

type ShopRepository interface {
	FindById(context.Context, int64) (*Shop, error)
	FindTypePage(context.Context, int64, int, int) ([]*Shop, error)
	FindNamePage(context.Context, string, int, int) ([]*Shop, error)
	FindByIds(context.Context, []int64) (map[int64]*Shop, error)
}
