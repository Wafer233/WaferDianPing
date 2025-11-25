package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type VoucherOrderCache struct {
	rdb *redis.Client
}

func NewVoucherOrderCache(rdb *redis.Client) *VoucherOrderCache {
	return &VoucherOrderCache{rdb: rdb}
}

func (cache *VoucherOrderCache) Eval(ctx context.Context,
	script string, keys []string, args ...interface{}) {

	cache.rdb.Eval(ctx, script, keys, args...)
	return
}
