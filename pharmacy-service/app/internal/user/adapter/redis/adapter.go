package redis

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/usecase"
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type adapter struct {
	r *redis.Client
}

func NewAdapter(r *redis.Client) usecase.RedisAdapter {
	return &adapter{
		r: r,
	}
}

func (a *adapter) SaveSession(ctx context.Context, sessionID string, userID int, ttl time.Duration) *de.DomainError {
	err := a.r.Set(ctx, sessionID, userID, ttl).Err()
	if err != nil {
		log.Error(err)
		return de.ErrSaveSession
	}

	return nil
}

func (a *adapter) GetSession(ctx context.Context, sessionID string) (int, *de.DomainError) {
	user, err := a.r.Get(ctx, sessionID).Result()
	if err == redis.Nil {
		return 0, de.ErrUnauthorized
	}

	if err != nil {
		log.Error(err)
		return 0, de.ErrGetSession
	}

	userID, err := strconv.Atoi(user)
	if err != nil {
		log.Error(err)
		return 0, de.ErrGetSession
	}

	return userID, nil
}

func (a *adapter) DeleteSession(ctx context.Context, sessionID string) *de.DomainError {
	if _, err := a.r.Del(ctx, sessionID).Result(); err != nil {
		log.Error(err)
		return de.ErrDeleteSession
	}

	return nil
}
