package application

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/infrastructure/cache"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/times"
	"github.com/jinzhu/copier"
)

type ShopService struct {
	repo  domain.ShopRepository
	cache *cache.GeoCache
}

func NewShopService(
	repo domain.ShopRepository,
	cache *cache.GeoCache,
) *ShopService {
	return &ShopService{repo: repo,
		cache: cache}
}

func (svc *ShopService) FindById(ctx context.Context, id int64) (ShopVO, error) {

	shop, err := svc.repo.FindById(ctx, id)
	if err != nil {
		return ShopVO{}, err
	}
	vo := ShopVO{}
	_ = copier.Copy(&vo, shop)

	vo.CreateTime = times.FormatTime(shop.CreateTime, "2006-01-02 15:04:05")
	vo.UpdateTime = times.FormatTime(shop.UpdateTime, "2006-01-02 15:04:05")
	return vo, nil
}

func (svc *ShopService) FindTypePage(ctx context.Context, typeId int64,
	page int, size int, x, y float64) ([]ShopVO, error) {

	//=========== 1. 没有坐标 → 普通分页 ============
	if x == 0 || y == 0 {
		shops, err := svc.repo.FindTypePage(ctx, typeId, page, size)
		if err != nil || len(shops) == 0 {
			return nil, err
		}
		vos := make([]ShopVO, len(shops))

		//
		//for i := range shops {
		//	svc.cache.GeoAdd(ctx, typeId, shops[i].Id, shops[i].X, shops[i].Y)
		//}

		_ = copier.Copy(&vos, &shops)

		for i, _ := range vos {
			vos[i].UpdateTime = times.FormatTime(shops[i].UpdateTime, "2006-01-02 15:04:05")
			vos[i].CreateTime = times.FormatTime(shops[i].CreateTime, "2006-01-02 15:04:05")
		}

		return vos, nil
	} else {
		// 2. 计算分页范围
		from := (page - 1) * size
		end := page * size

		//	查redis
		geoRet, err := svc.cache.GeoSearch(ctx, typeId, x, y, end)

		if err != nil {
			return nil, err
		}
		if len(geoRet) == 0 || len(geoRet) < from {
			return nil, nil
		}

		details := geoRet[from:]

		ids := make([]int64, len(details))
		dests := make([]float64, len(details))

		for i, _ := range details {
			ids[i] = details[i].ShopId
			dests[i] = details[i].Dist
		}

		shops, err := svc.repo.FindByIds(ctx, ids)

		vos := make([]ShopVO, len(shops))

		for i, id := range ids {
			curShop := shops[id]
			_ = copier.Copy(&vos[i], curShop)
			vos[i].UpdateTime = times.FormatTime(shops[id].UpdateTime, "2006-01-02 15:04:05")
			vos[i].CreateTime = times.FormatTime(shops[id].CreateTime, "2006-01-02 15:04:05")
			vos[i].Distance = int64(dests[i])
		}
		return vos, nil

	}

}

func (svc *ShopService) FindNamePage(ctx context.Context, name string,
	page int, size int) ([]*ShopVO, error) {

	shops, err := svc.repo.FindNamePage(ctx, name, page, size)
	if err != nil || len(shops) == 0 {
		return nil, err
	}
	vos := make([]*ShopVO, len(shops))

	_ = copier.Copy(&vos, &shops)

	for i, _ := range vos {
		vos[i].UpdateTime = times.FormatTime(shops[i].UpdateTime, "2006-01-02 15:04:05")
		vos[i].CreateTime = times.FormatTime(shops[i].CreateTime, "2006-01-02 15:04:05")
	}
	return vos, nil

}
