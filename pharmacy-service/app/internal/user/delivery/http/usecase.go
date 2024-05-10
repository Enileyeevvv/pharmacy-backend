package http

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
)

type UseCase interface {
	SignUp(ctx context.Context, login, password string) *de.DomainError
	SignIn(ctx context.Context, login, password string) *de.DomainError
}
