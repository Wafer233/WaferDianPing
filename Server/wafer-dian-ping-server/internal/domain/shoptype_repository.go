package domain

import "context"

type ShopTypeRepository interface {
	FindShopTypeList(ctx context.Context) ([]*ShopType, error)
}
