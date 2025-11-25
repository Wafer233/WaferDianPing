package application

import (
	"context"
	"strconv"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/domain"
	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/pkg/times"

	"github.com/jinzhu/copier"

	"github.com/google/uuid"
)

type UserService struct {
	repo       domain.UserRepository
	sessionSvc *SessionService
}

func NewUserService(repo domain.UserRepository, sessionSvc *SessionService) *UserService {
	return &UserService{repo: repo, sessionSvc: sessionSvc}
}

func (svc *UserService) LoginService(ctx context.Context, dto *LoginDTO) (string, error) {

	// 获取当前id
	user, err := svc.repo.FindUserByPhone(ctx, dto.Phone)
	if err != nil || user == nil {
		return "", err
	}

	curId := user.Id
	curIdStr := strconv.Itoa(int(curId))

	//随便生成一个uuid作为sessionid
	sessionID := uuid.New().String()

	// 把session id存再redis里头
	err = svc.sessionSvc.Set(ctx, sessionID, curIdStr)
	if err != nil {
		return "", err
	}

	//返回sessionId 作为cookie的key
	return sessionID, nil
}

func (svc *UserService) FindUser(ctx context.Context, curId int64) (UserVO, error) {

	user, err := svc.repo.FindUserById(ctx, curId)
	if err != nil {
		return UserVO{}, err
	}

	vo := UserVO{}
	_ = copier.Copy(&vo, user)
	return vo, nil

}

func (svc *UserService) FindUserInfo(ctx context.Context, id int64) (UserInfoVO, error) {

	info, err := svc.repo.FindInfoByUserId(ctx, id)
	if err != nil {
		return UserInfoVO{}, err
	}

	vo := UserInfoVO{}
	_ = copier.Copy(&vo, info)
	vo.UpdateTime = times.FormatTime(info.UpdateTime, "2006-01-02 15:04:05")
	vo.CreateTime = times.FormatTime(info.CreateTime, "2006-01-02 15:04:05")
	vo.Birthday = times.FormatTime(info.Birthday, "2006-01-02")
	return vo, nil
}

func (svc *UserService) LogoutService(ctx context.Context, sessionId string) error {

	return svc.sessionSvc.Del(ctx, sessionId)
}

func (svc *UserService) FindUserByIds(ctx context.Context, ids []int64) (map[int64]*UserVO, error) {

	mapping, err := svc.repo.FindUserByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	mappingVO := make(map[int64]*UserVO)

	_ = copier.Copy(&mappingVO, &mapping)
	return mappingVO, nil
}
