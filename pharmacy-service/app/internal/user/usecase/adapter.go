package usecase

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
)

type PGAdapter interface {
	CheckIfUserExists(ctx context.Context, login string) (bool, *de.DomainError)
	CreateUser(ctx context.Context, login, passwordHash string) *de.DomainError
	GetPassword(ctx context.Context, login string) (string, *de.DomainError)
	GetUser(ctx context.Context, id int) (*User, *de.DomainError)
}
