package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/internal/infrastructure/cache"
)

type SessionService struct {
	repo *cache.SessionRepository
}

func NewSessionService(repo *cache.SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

func (svc *SessionService) Get(ctx context.Context, sessionId string) (string, error) {

	key := "SessionId:" + sessionId
	return svc.repo.Get(ctx, key)

}

func (svc *SessionService) Set(ctx context.Context, sessionID string, userId string) error {

	key := "SessionId:" + sessionID
	ttl := 1 * time.Hour
	return svc.repo.Set(ctx, key, userId, ttl)
}

func (svc *SessionService) Del(ctx context.Context, sessionId string) error {
	key := "SessionId:" + sessionId
	return svc.repo.Del(ctx, key)
}
