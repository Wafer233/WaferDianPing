package application

import (
	"context"
	"errors"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/times"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
)

type BlogService struct {
	repo    domain.BlogRepository
	userSvc *UserService
	likeSvc *LikeService
}

func NewBlogService(
	repo domain.BlogRepository,
	userSvc *UserService,
	likeSvc *LikeService,
) *BlogService {
	return &BlogService{
		repo:    repo,
		userSvc: userSvc,
		likeSvc: likeSvc,
	}
}

func (svc *BlogService) FindBlogByUserId(ctx context.Context, userId int64) ([]BlogVO, error) {

	user, err := svc.userSvc.FindUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	blogs, err := svc.repo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	vos := make([]BlogVO, len(blogs))
	_ = copier.Copy(&vos, &blogs)

	for i, _ := range vos {
		vos[i].Icon = user.Icon
		vos[i].Name = user.NickName
		vos[i].CreateTime = times.FormatTime(blogs[i].CreateTime, "2006-01-02 15:04:05")
		vos[i].UpdateTime = times.FormatTime(blogs[i].UpdateTime, "2006-01-02 15:04:05")
	}
	return vos, nil

}

func (svc *BlogService) Create(ctx context.Context, dto *BlogDTO, userId int64) (int64, error) {

	blog := domain.Blog{}
	_ = copier.Copy(&blog, dto)

	blog.UserId = userId
	id, err := svc.repo.Create(ctx, &blog)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc *BlogService) FindPageHot(ctx context.Context, page int, pageSize int) ([]BlogVO, error) {

	blogs, err := svc.repo.FindPageHot(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}
	if len(blogs) == 0 {
		return []BlogVO{}, nil
	}

	vos := make([]BlogVO, len(blogs))
	_ = copier.Copy(&vos, &blogs)

	// 查user
	ids := make([]int64, len(blogs))
	for i, _ := range vos {
		ids[i] = blogs[i].UserId
	}

	mapping, err := svc.userSvc.FindUserByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	for i, _ := range vos {
		vos[i].CreateTime = times.FormatTime(blogs[i].CreateTime, "2006-01-02 15:04:05")
		vos[i].UpdateTime = times.FormatTime(blogs[i].UpdateTime, "2006-01-02 15:04:05")
		vos[i].Name = mapping[vos[i].UserId].NickName
		vos[i].Icon = mapping[vos[i].UserId].Icon
	}

	return vos, nil
}

func (svc *BlogService) FindById(ctx context.Context, id int64,
	curId int64) (BlogVO, error) {

	blog, err := svc.repo.FindById(ctx, id)
	if err != nil || blog == nil {
		return BlogVO{}, err
	}
	vo := BlogVO{}
	_ = copier.Copy(&vo, &blog)

	vo.UpdateTime = times.FormatTime(blog.UpdateTime, "2006-01-02 15:04:05")
	vo.CreateTime = times.FormatTime(blog.CreateTime, "2006-01-02 15:04:05")

	user, err := svc.userSvc.FindUser(ctx, blog.UserId)
	if err != nil {
		return BlogVO{}, err
	}
	vo.Name = user.NickName
	vo.Icon = user.Icon

	_, err = svc.likeSvc.ZScore(ctx, curId, id)
	if err == nil {
		vo.IsLike = true
	}

	return vo, nil
}

func (svc *BlogService) LikeBlog(ctx context.Context, userId int64, blogId int64) error {

	_, err := svc.likeSvc.ZScore(ctx, userId, blogId)

	// err == redis.Nil -> 没 score -> 没点赞需要点赞
	if errors.Is(err, redis.Nil) {

		er := svc.repo.LikeById(ctx, blogId)
		if er != nil {
			return er
		}

		er = svc.likeSvc.ZAdd(ctx, userId, blogId)
		if er != nil {
			return er
		}
		return nil
	}

	// 已经点赞了
	if err == nil {
		er := svc.repo.DislikeById(ctx, blogId)
		if er != nil {
			return er
		}

		er = svc.likeSvc.ZRem(ctx, userId, blogId)
		if er != nil {
			return er
		}
		return nil
	}

	// 其他奇奇怪怪的错
	return err

}

func (svc *BlogService) TopLikes(ctx context.Context, blogId int64) ([]UserVO, error) {

	ids, err := svc.likeSvc.ZRange(ctx, blogId)
	if err != nil || len(ids) == 0 {
		return nil, err
	}

	vos := make([]UserVO, len(ids))
	for index, id := range ids {

		user, er := svc.userSvc.FindUser(ctx, id)
		if er != nil {
			return nil, er
		}
		vos[index] = user
	}
	return vos, nil
}
