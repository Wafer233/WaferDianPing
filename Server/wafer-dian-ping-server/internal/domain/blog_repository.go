package domain

import "context"

type BlogRepository interface {
	FindByUserId(context.Context, int64) ([]*Blog, error)
	FindByShopId(context.Context, int64) ([]*Blog, error)
	Create(context.Context, *Blog) (int64, error)
	FindPageHot(context.Context, int, int) ([]*Blog, error)
	FindById(context.Context, int64) (*Blog, error)
	LikeById(context.Context, int64) error
	DislikeById(context.Context, int64) error
}
