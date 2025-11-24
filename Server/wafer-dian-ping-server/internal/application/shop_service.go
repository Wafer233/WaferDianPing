package application

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/times"
	"github.com/jinzhu/copier"
)

type ShopService struct {
	repo domain.ShopRepository
}

func NewShopService(repo domain.ShopRepository) *ShopService {
	return &ShopService{repo: repo}
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

func (svc *ShopService) FindPage(ctx context.Context, typeId int64,
	page int, size int) ([]*ShopVO, error) {

	shops, err := svc.repo.FindPage(ctx, typeId, page, size)
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
