package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/redis/go-redis/v9"
)

type CachedShopRepository struct {
	repo  *DefaultShopRepository
	cache *redis.Client
}

func (repo *CachedShopRepository) FindById(ctx context.Context, id int64) (*domain.Shop, error) {

	idStr := strconv.FormatInt(id, 10)
	key := "shop:shop:" + idStr
	val, err := repo.cache.Get(ctx, key).Result()
	if err == nil {
		var shopCache domain.Shop
		er := json.Unmarshal([]byte(val), &shopCache)
		if er != nil {
			return nil, er
		}
		return &shopCache, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	shop, err := repo.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	// 如果数据库也查不到 → 写一个短暂空值避免穿透
	if shop == nil {
		repo.cache.Set(ctx, key, "", time.Minute*2)
		return nil, nil
	}

	bytes, _ := json.Marshal(shop)
	ttl := time.Hour
	repo.cache.Set(ctx, key, string(bytes), ttl)
	return shop, nil
}

func (repo *CachedShopRepository) FindPage(ctx context.Context,
	typeId int64, page int, pageSize int) ([]*domain.Shop, error) {

	//key: shop:shops:{typeId}:{page}:
	typeIdStr := strconv.FormatInt(typeId, 10)
	pageStr := strconv.Itoa(page)

	key := "shop:shops:" + typeIdStr + ":" + pageStr
	val, err := repo.cache.Get(ctx, key).Result()
	if err == nil {
		var shopsCache []*domain.Shop
		er := json.Unmarshal([]byte(val), &shopsCache)
		if er != nil {
			return nil, er
		}
		return shopsCache, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	shops, err := repo.repo.FindPage(ctx, typeId, page, pageSize)
	if err != nil {
		return nil, err
	}

	if len(shops) == 0 {
		repo.cache.Set(ctx, key, "[]", time.Minute*2)
		return nil, nil
	}

	ttl := time.Hour
	bytes, _ := json.Marshal(shops)
	repo.cache.Set(ctx, key, string(bytes), ttl)

	return shops, nil
}

func NewCachedShopRepository(repo *DefaultShopRepository, cache *redis.Client) domain.ShopRepository {
	return &CachedShopRepository{repo, cache}
}
