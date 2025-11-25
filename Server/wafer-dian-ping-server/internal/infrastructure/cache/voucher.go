package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type VoucherCache struct {
	rdb *redis.Client
}

func NewVoucherCache(rdb *redis.Client) *VoucherCache {
	return &VoucherCache{rdb: rdb}
}

func (cache *VoucherCache) Set(ctx context.Context,
	voucherId int64, stock int) error {

	vouId := strconv.FormatInt(voucherId, 10)
	key := "seckill:stock:" + vouId
	ttl := 1 * time.Hour
	err := cache.rdb.Set(ctx, key, stock, ttl).Err()
	return err
}
