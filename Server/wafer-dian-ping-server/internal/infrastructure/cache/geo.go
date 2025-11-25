package cache

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type GeoCache struct {
	rdb *redis.Client
}

func NewGeoCache(rdb *redis.Client) *GeoCache {
	return &GeoCache{rdb: rdb}
}

func (cache *GeoCache) GeoAdd(ctx context.Context, typeId int64, shopId int64, x, y float64) error {

	id := strconv.FormatInt(typeId, 10)
	key := "shop:geo:" + id
	shop := strconv.FormatInt(shopId, 10)

	location := redis.GeoLocation{
		Name:      shop,
		Latitude:  y,
		Longitude: x,
	}
	err := cache.rdb.GeoAdd(ctx, key, &location).Err()
	return err
}

type GeoDetail struct {
	ShopId int64
	Dist   float64
}

func (cache *GeoCache) GeoSearch(ctx context.Context, typeId int64, x, y float64, end int) ([]GeoDetail, error) {

	id := strconv.FormatInt(typeId, 10)
	key := "shop:geo:" + id

	query := redis.GeoSearchLocationQuery{
		GeoSearchQuery: redis.GeoSearchQuery{
			Longitude:  x,
			Latitude:   y,
			Radius:     5,
			RadiusUnit: "km",
			Count:      end,
			Sort:       "ASC",
		},
		WithDist: true,
	}
	val, err := cache.rdb.GeoSearchLocation(ctx, key, &query).Result()
	if err != nil {
		return nil, err
	}

	shops := make([]GeoDetail, len(val))

	for i, _ := range val {
		shopid, _ := strconv.ParseInt(val[i].Name, 10, 64)
		shops[i].ShopId = shopid
		shops[i].Dist = val[i].Dist
	}
	return shops, nil
}
