package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisDatabase() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "",
		DB:       0,
	})

	ping := rdb.Ping(context.Background())
	if err := ping.Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
