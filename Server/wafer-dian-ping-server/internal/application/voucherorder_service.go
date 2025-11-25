package application

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/infrastructure/cache"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/snow"
	"github.com/google/uuid"
)

type VoucherOrderService struct {
	repo  domain.VoucherOrderRepository
	cache *cache.VoucherOrderCache
	svc   *VoucherService
}

func NewVoucherOrderService(
	repo domain.VoucherOrderRepository,
	svc *VoucherService,
	cache *cache.VoucherOrderCache,

) *VoucherOrderService {
	return &VoucherOrderService{
		repo:  repo,
		svc:   svc,
		cache: cache,
	}
}

func (svc *VoucherOrderService) SeckillService(ctx context.Context,
	voucherId int64, userId int64,
) (int64, error) {
	//// ====================查询秒杀券
	seckill, err := svc.svc.FindSecKillById(ctx, voucherId)
	stock := seckill.Stock
	if err != nil {
		return 0, err
	}

	//============.判断是否开始&&结束
	now := time.Now()
	//begin := times.ToTime(seckill.BeginTime)
	//end := times.ToTime(seckill.EndTime)

	//if begin.After(now) {
	//	return 0, errors.New("秒杀还未开始")
	//}
	//if end.Before(now) {
	//	return 0, errors.New("秒杀已经结束")
	//}

	//==========2.判断库存是否充足
	if seckill.Stock <= 0 {
		return 0, errors.New("被抢光了")
	}

	// 4. 用户级 分布式锁
	vou := strconv.FormatInt(voucherId, 10)
	key := "seckill:" + vou
	lockKey := key + ":lock"
	lockId := uuid.New().String()
	luaByte, err := os.ReadFile("./script/redis/unlock.lua")
	if err != nil {
		return 0, err
	}
	luaReleaseLock := string(luaByte)
	defer svc.cache.Eval(ctx, luaReleaseLock, []string{lockKey}, lockId)

	//==============判断用户是否购买过
	exist, err := svc.repo.Exists(ctx, userId, voucherId)
	if err != nil {
		return 0, err
	}
	if exist {
		return 0, errors.New("请勿重复购买")
	}

	// =================扣减库存（MySQL）
	ok, err := svc.svc.DecrStock(ctx, voucherId, stock)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("库存不足")
	}

	// 7. 创建订单（MySQL）
	orderId := snow.NewId(1) // 雪花算法
	order := &domain.VoucherOrder{
		Id:         orderId,
		UserId:     userId,
		VoucherId:  voucherId,
		Status:     1,
		PayType:    1,
		CreateTime: &now,
		UpdateTime: &now,
	}

	orderId, err = svc.repo.Create(ctx, order)
	if err != nil {
		return 0, err
	}

	return orderId, nil
}
