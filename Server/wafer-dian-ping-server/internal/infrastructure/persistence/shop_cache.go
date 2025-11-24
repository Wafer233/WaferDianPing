package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/google/uuid"
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

	// 拼接 key
	//key: shop:shops:{typeId}:{page}:
	typeIdStr := strconv.FormatInt(typeId, 10)
	pageStr := strconv.Itoa(page)
	key := "shop:shops:" + typeIdStr + ":" + pageStr

	// =============== 1. 从 Redis 查询商铺缓存 ===============
	val, err := repo.cache.Get(ctx, key).Result()
	if err == nil {
		//缓存命中 → 判断是否为空值
		if val == "[]" {
			// 空值命中 → 返回 nil，不查数据库
			return nil, nil
		}

		// 正常缓存命中 → 反序列化并返回
		var shopsCache []*domain.Shop
		er := json.Unmarshal([]byte(val), &shopsCache)
		if er != nil {
			return nil, er
		}
		return shopsCache, nil
	}

	//排除未命中之外的其他可能性
	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	// =============== 2. 缓存未命中 → 尝试获取互斥锁 ===============
	lockKey := key + ":lock"
	lockId := uuid.New().String()
	luaByte, err := os.ReadFile("./script/redis/unlock.lua")
	if err != nil {
		return nil, err
	}
	luaReleaseLock := string(luaByte)
	for {

		ok, _ := repo.cache.SetNX(ctx, lockKey, lockId, time.Second*5).Result()

		// =============== 4. 获取到锁 ===============
		if ok {
			// 竞争到了锁
			// =============== 7. 释放互斥锁 ===============
			//repo.cache.Del(ctx, lockKey)
			defer repo.cache.Eval(ctx, luaReleaseLock, []string{lockKey}, lockId)
			break
		}

		// =============== 3. 未获得锁 → 休眠后继续查缓存 ===============
		time.Sleep(time.Millisecond * 50)

		//再查缓存
		// 用2代表不是第一轮的
		val2, err2 := repo.cache.Get(ctx, key).Result()

		if err2 == nil {
			if val2 == "[]" {
				return nil, nil
			}
			var shopsCache2 []*domain.Shop
			_ = json.Unmarshal([]byte(val2), &shopsCache2)
			return shopsCache2, nil
		}

		//排除未命中之外的其他可能性
		if !errors.Is(err2, redis.Nil) {
			return nil, err2
		}

	}

	// =============== 5. 得到锁 → 查询数据库 ===============
	shops, err := repo.repo.FindPage(ctx, typeId, page, pageSize)
	if err != nil {
		return nil, err
	}

	// =============== 6. 将数据写入 Redis ===============
	if len(shops) == 0 {
		//数据库不存在 → 写入空值到 Redis（防缓存穿透）
		repo.cache.Set(ctx, key, "[]", time.Minute*2)
		return nil, nil
	}

	//数据存在 → 序列化写入 Redis（正常缓存）
	ttl := time.Hour
	bytes, _ := json.Marshal(shops)
	repo.cache.Set(ctx, key, string(bytes), ttl)

	return shops, nil
}

func NewCachedShopRepository(repo *DefaultShopRepository, cache *redis.Client) domain.ShopRepository {
	return &CachedShopRepository{repo, cache}
}
