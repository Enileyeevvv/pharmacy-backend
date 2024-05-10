package usecase

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"time"
)

type PGAdapter interface {
	CheckIfUserExists(ctx context.Context, login string) (bool, *de.DomainError)
	CreateUser(ctx context.Context, login, passwordHash string) *de.DomainError
	GetPassword(ctx context.Context, login string) (string, *de.DomainError)
	GetUserID(ctx context.Context, login string) (int, *de.DomainError)
	GetUser(ctx context.Context, id int) (*User, *de.DomainError)
}

type RedisAdapter interface {
	SaveSession(ctx context.Context, sessionID string, userID int, ttl time.Duration) *de.DomainError
	GetSession(ctx context.Context, sessionID string) (int, *de.DomainError)
	DeleteSession(ctx context.Context, sessionID string) *de.DomainError
}
