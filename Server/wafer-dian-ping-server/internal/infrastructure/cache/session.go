package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionRepository struct {
	rdb *redis.Client
}

func NewSessionRepository(rdb *redis.Client) *SessionRepository {
	return &SessionRepository{rdb: rdb}
}

func (cache *SessionRepository) Get(ctx context.Context, key string) (string, error) {

	rdb := cache.rdb.Get(ctx, key)
	return rdb.Result()
}

func (cache *SessionRepository) Set(ctx context.Context, key string, value string,
	ttl time.Duration) error {

	rdb := cache.rdb.Set(ctx, key, value, ttl)
	return rdb.Err()
}

func (cache *SessionRepository) Del(ctx context.Context, key string) error {
	rdb := cache.rdb.Del(ctx, key)
	return rdb.Err()
}
