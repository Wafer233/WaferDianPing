package domain

import "context"

type FollowRepository interface {
	FindFollow(context.Context, int64, int64) (bool, error)
	AddFollow(context.Context, *Follow) error
	RemoveFollow(context.Context, *Follow) error
}
