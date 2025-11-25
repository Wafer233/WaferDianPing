package application

import (
	"context"
	"strconv"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/infrastructure/cache"
)

type FollowService struct {
	repo    domain.FollowRepository
	cache   *cache.FollowCache
	userSvc *UserService
}

func NewFollowService(
	repo domain.FollowRepository,
	cache *cache.FollowCache,
	userSvc *UserService,
) *FollowService {
	return &FollowService{
		repo:    repo,
		cache:   cache,
		userSvc: userSvc,
	}

}

func (svc *FollowService) IsFollow(ctx context.Context,
	userId int64, curId int64) (bool, error) {

	ok, err := svc.repo.FindFollow(ctx, userId, curId)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (svc *FollowService) Follow(ctx context.Context,
	userId int64, curId int64, isFol bool) error {

	now := time.Now()
	follow := domain.Follow{
		UserId:       userId,
		FollowUserId: curId,
		CreateTime:   &now,
	}

	if !isFol {
		err := svc.repo.RemoveFollow(ctx, &follow)
		if err != nil {
			return err
		}

		err = svc.cache.SRem(ctx, userId, curId)
		if err != nil {
			return err
		}

		return nil
	} else {
		err := svc.repo.AddFollow(ctx, &follow)
		if err != nil {
			return err
		}

		err = svc.cache.SAdd(ctx, userId, curId)
		if err != nil {
			return err
		}
		return nil
	}

}

func (svc *FollowService) CommonFollow(ctx context.Context,
	userId int64, curId int64) ([]UserVO, error) {

	idsStr, err := svc.cache.SInter(ctx, userId, curId)
	if err != nil {
		return nil, err
	}
	if len(idsStr) == 0 {
		return nil, nil
	}

	ids := make([]int64, len(idsStr))
	for i, id := range idsStr {
		cur, er := strconv.ParseInt(id, 10, 64)
		if er != nil {
			return nil, er
		}
		ids[i] = cur
	}

	mapping, err := svc.userSvc.FindUserByIds(ctx, ids)

	vos := make([]UserVO, len(ids))

	for i, id := range ids {
		vos[i].Id = mapping[id].Id
		vos[i].Icon = mapping[id].Icon
		vos[i].NickName = mapping[id].NickName
	}
	return vos, nil
}
