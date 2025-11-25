package cache

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type FollowCache struct {
	rdb *redis.Client
}

func NewFollowCache(rdb *redis.Client) *FollowCache {
	return &FollowCache{rdb: rdb}
}

func (cache *FollowCache) SAdd(ctx context.Context, userId int64, curId int64) error {

	curIdStr := strconv.FormatInt(curId, 10)
	key := "follow:" + curIdStr

	err := cache.rdb.SAdd(ctx, key, userId).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cache *FollowCache) SRem(ctx context.Context, userId int64, curId int64) error {
	curIdStr := strconv.FormatInt(curId, 10)
	key := "follow:" + curIdStr

	err := cache.rdb.SRem(ctx, key, userId).Err()
	if err != nil {
		return err
	}
	return nil

}

func (cache *FollowCache) SInter(ctx context.Context, userId int64, curId int64) ([]string, error) {
	userIdStr := strconv.FormatInt(userId, 10)
	key1 := "follow:" + userIdStr

	curIdStr := strconv.FormatInt(curId, 10)
	key2 := "follow:" + curIdStr

	ids, err := cache.rdb.SInter(ctx, key1, key2).Result()
	if err != nil {
		return nil, err
	}
	return ids, nil

}
