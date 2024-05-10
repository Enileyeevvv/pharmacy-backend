package http

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"time"
)

type UseCase interface {
	CreateUser(ctx context.Context, login, password string) *de.DomainError
	CheckPassword(ctx context.Context, login, password string) *de.DomainError
	CreateSession(ctx context.Context, login string) (string, time.Time, *de.DomainError)
	GetUserIDFromSession(ctx context.Context, token string) (int, *de.DomainError)
	CheckUserRole(ctx context.Context, userID, roleID int) (bool, *de.DomainError)
	DeleteSession(ctx context.Context, token string) *de.DomainError
	GetUserLoginAndRoleID(ctx context.Context, userID int) (string, int, *de.DomainError)
}
