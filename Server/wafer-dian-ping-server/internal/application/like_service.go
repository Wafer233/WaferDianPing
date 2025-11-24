package application

import (
	"context"
	"strconv"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/infrastructure/cache"
)

type LikeService struct {
	repo *cache.LikeRepository
}

func NewLikeService(repo *cache.LikeRepository) *LikeService {
	return &LikeService{repo: repo}
}

func (svc *LikeService) ZScore(ctx context.Context,
	userId int64, blogId int64) (float64, error) {

	blogStr := strconv.FormatInt(blogId, 10)
	key := "blog:like:" + blogStr

	member := strconv.FormatInt(userId, 10)
	return svc.repo.ZScore(ctx, key, member)

}

func (svc *LikeService) ZAdd(ctx context.Context,
	userId int64, blogId int64) error {
	blogStr := strconv.FormatInt(blogId, 10)
	key := "blog:like:" + blogStr
	member := strconv.FormatInt(userId, 10)

	score := float64(time.Now().UnixMilli())
	_, err := svc.repo.ZAdd(ctx, key, member, score)
	return err

}

func (svc *LikeService) ZRem(ctx context.Context,
	userId int64, blogId int64) error {
	blogStr := strconv.FormatInt(blogId, 10)
	key := "blog:like:" + blogStr
	member := strconv.FormatInt(userId, 10)

	_, err := svc.repo.ZRem(ctx, key, member)
	return err

}

func (svc *LikeService) ZRange(ctx context.Context, blogId int64) ([]int64, error) {
	blogStr := strconv.FormatInt(blogId, 10)
	key := "blog:like:" + blogStr

	idsStr, err := svc.repo.ZRange(ctx, key, 0, 4)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(idsStr))
	for i, idStr := range idsStr {
		id, er := strconv.ParseInt(idStr, 10, 64)
		if er != nil {
			return nil, er
		}
		ids[i] = id
	}
	return ids, nil
}
