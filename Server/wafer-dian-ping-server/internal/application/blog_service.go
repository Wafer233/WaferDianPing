package application

import (
	"context"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/times"
	"github.com/jinzhu/copier"
)

type BlogService struct {
	repo    domain.BlogRepository
	userSvc *UserService
}

func NewBlogService(repo domain.BlogRepository, userSvc *UserService) *BlogService {
	return &BlogService{repo: repo, userSvc: userSvc}
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
