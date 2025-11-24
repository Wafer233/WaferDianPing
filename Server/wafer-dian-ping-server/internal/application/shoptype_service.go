package application

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/jinzhu/copier"
)

type ShopTypeService struct {
	repo domain.ShopTypeRepository
}

func NewShopTypeService(repo domain.ShopTypeRepository) *ShopTypeService {
	return &ShopTypeService{repo: repo}
}

func (svc *ShopTypeService) FindShopTypeList(ctx context.Context) ([]ShopTypeVO, error) {

	entities, err := svc.repo.FindShopTypeList(ctx)
	if err != nil || len(entities) == 0 {
		return []ShopTypeVO{}, err
	}

	vos := make([]ShopTypeVO, len(entities))

	_ = copier.Copy(&vos, &entities)

	for i := range vos {
		vos[i].CreateTime = entities[i].CreateTime.Format("2006-01-02 15:04:05")
		vos[i].UpdateTime = entities[i].UpdateTime.Format("2006-01-02 15:04:05")
	}

	return vos, nil

}
