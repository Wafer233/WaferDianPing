package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type LikeRepository struct {
	rdb *redis.Client
}

func NewLikeRepository(rdb *redis.Client) *LikeRepository {
	return &LikeRepository{rdb: rdb}
}

func (repo *LikeRepository) ZScore(ctx context.Context, key, member string) (float64, error) {

	return repo.rdb.ZScore(ctx, key, member).Result()

}

func (repo *LikeRepository) ZAdd(ctx context.Context, key, member string, score float64) (int64, error) {
	zset := redis.Z{
		Score:  score,
		Member: member,
	}
	return repo.rdb.ZAdd(ctx, key, zset).Result()
}

func (repo *LikeRepository) ZRem(ctx context.Context, key, member string) (int64, error) {
	return repo.rdb.ZRem(ctx, key, member).Result()
}

func (repo *LikeRepository) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {

	ids, err := repo.rdb.ZRange(ctx, key, start, stop).Result()
	if err != nil || len(ids) == 0 {
		return nil, err
	}
	return ids, nil

}
