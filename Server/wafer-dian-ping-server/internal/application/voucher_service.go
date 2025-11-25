package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/infrastructure/cache"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/times"
	"github.com/jinzhu/copier"
)

type VoucherService struct {
	repo  domain.VoucherRepository
	cache *cache.VoucherCache
}

func NewVoucherService(
	repo domain.VoucherRepository,
	cache *cache.VoucherCache,
) *VoucherService {
	return &VoucherService{
		repo:  repo,
		cache: cache,
	}
}

func (svc *VoucherService) CreateVoucher(ctx context.Context, dto *VoucherDTO) (int64, error) {

	vou := domain.Voucher{}
	_ = copier.Copy(&vou, dto)

	layout := "2006-01-02 15:04:05"
	createTime, _ := time.Parse(layout, dto.CreateTime)
	updateTime, _ := time.Parse(layout, dto.UpdateTime)
	vou.CreateTime = &createTime
	vou.UpdateTime = &updateTime

	id, err := svc.repo.CreateVoucher(ctx, &vou)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc *VoucherService) CreateSeckillVoucher(ctx context.Context, dto *VoucherDTO) (int64, error) {
	//转换时间
	createTime := times.ToTime(dto.CreateTime)
	updateTime := times.ToTime(dto.UpdateTime)
	beginTime := times.ToTime(dto.BeginTime)
	endTime := times.ToTime(dto.EndTime)

	// 保存
	vou := domain.Voucher{}
	_ = copier.Copy(&vou, dto)
	vou.CreateTime = createTime
	vou.UpdateTime = updateTime

	id, err := svc.repo.CreateVoucher(ctx, &vou)
	if err != nil {
		return 0, err
	}

	//保存秒杀全
	seckill := domain.SecKillVoucher{}
	seckill.BeginTime = beginTime
	seckill.EndTime = endTime
	seckill.CreateTime = createTime
	seckill.UpdateTime = updateTime
	seckill.Stock = dto.Stock
	seckill.VoucherId = id
	//seckill.VoucherId = snow.NewId(1)
	id, err = svc.repo.CreateSeckillVoucher(ctx, &seckill)
	if err != nil {
		return 0, err
	}

	//保存到redis里面
	err = svc.cache.Set(ctx, id, dto.Stock)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (svc *VoucherService) FindVoucher(ctx context.Context,
	shopId int64) ([]VoucherVO, error) {

	vos := make([]VoucherVO, 0)

	vouchers, ids, err := svc.repo.FindByShopId(ctx, shopId)

	if err != nil {
		return nil, err
	}
	if len(vouchers) == 0 {
		return vos, nil
	}

	_ = copier.Copy(&vos, &vouchers)

	// 如果有 ids，则查秒杀信息
	if len(ids) > 0 {
		mapping, er := svc.repo.FindSecKillByIds(ctx, ids)
		if er != nil {
			return nil, er
		}

		// 将 mapping 信息填到 vos 里（假设你需要）
		for i := range vos {
			if sk, ok := mapping[vos[i].Id]; ok {
				layout := "2006-01-02 15:04:05"
				begin := times.FormatTime(sk.BeginTime, layout)
				end := times.FormatTime(sk.EndTime, layout)

				vos[i].BeginTime = begin
				vos[i].EndTime = end
				vos[i].Stock = sk.Stock
			}
		}
	}

	return vos, nil
}
