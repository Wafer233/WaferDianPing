package domain

import "context"

type BlogRepository interface {
	FindByUserId(context.Context, int64) ([]*Blog, error)
	FindByShopId(context.Context, int64) ([]*Blog, error)
}
