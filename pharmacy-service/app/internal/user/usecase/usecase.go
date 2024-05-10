package usecase

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/gofiber/fiber/v2/log"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	pgAdp PGAdapter
}

func NewUseCase(pgAdp PGAdapter) *UseCase {
	return &UseCase{
		pgAdp: pgAdp,
	}
}

func (uc *UseCase) SignUp(ctx context.Context, login, password string) *de.DomainError {
	exists, dErr := uc.pgAdp.CheckIfUserExists(ctx, login)
	if dErr != nil {
		return dErr
	}

	if exists {
		return de.ErrUserAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
		return de.ErrSignUp
	}

	return uc.pgAdp.CreateUser(ctx, login, string(passwordHash))
}

func (uc *UseCase) SignIn(ctx context.Context, login, password string) *de.DomainError {
	exists, dErr := uc.pgAdp.CheckIfUserExists(ctx, login)
	if dErr != nil {
		return dErr
	}

	if !exists {
		return de.ErrUserNotFound
	}

	passwordHash, dErr := uc.pgAdp.GetPassword(ctx, login)
	if dErr != nil {
		return dErr
	}

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return de.ErrWrongPassword
	}

	if err != nil {
		return de.ErrSignIn
	}

	return nil
}
