package usecase

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UseCase struct {
	pgAdp      PGAdapter
	redisAdp   RedisAdapter
	sessionTTL time.Duration
	secret     string
}

func NewUseCase(
	pgAdp PGAdapter,
	redisAdp RedisAdapter,
	sessionTTL time.Duration,
	secret string,
) *UseCase {
	return &UseCase{
		pgAdp:      pgAdp,
		redisAdp:   redisAdp,
		sessionTTL: sessionTTL,
		secret:     secret,
	}
}

func (uc *UseCase) CreateUser(ctx context.Context, login, password string) *de.DomainError {
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

func (uc *UseCase) CheckPassword(ctx context.Context, login, password string) *de.DomainError {
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

func (uc *UseCase) CreateSession(ctx context.Context, login string) (string, time.Time, *de.DomainError) {
	payload := jwt.MapClaims{
		"login": login,
		"ttl":   uc.sessionTTL,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString([]byte(uc.secret))
	if err != nil {
		log.Error(err)
		return "", time.Time{}, de.ErrCreateSession
	}

	userID, dErr := uc.pgAdp.GetUserID(ctx, login)
	if err != nil {
		return "", time.Time{}, dErr
	}

	dErr = uc.redisAdp.SaveSession(ctx, t, userID, uc.sessionTTL)
	if dErr != nil {
		return "", time.Time{}, dErr
	}

	expireAt := time.Now().Add(uc.sessionTTL)

	return t, expireAt, nil
}

func (uc *UseCase) GetUserIDFromSession(ctx context.Context, token string) (int, *de.DomainError) {
	return uc.redisAdp.GetSession(ctx, token)
}

func (uc *UseCase) CheckUserRole(ctx context.Context, userID, roleID int) (bool, *de.DomainError) {
	user, err := uc.pgAdp.GetUser(ctx, userID)
	if err != nil {
		return false, err
	}

	roleMatch := user.RoleID == roleID
	return roleMatch, nil
}

func (uc *UseCase) DeleteSession(ctx context.Context, token string) *de.DomainError {
	return uc.redisAdp.DeleteSession(ctx, token)
}

func (uc *UseCase) GetUserLoginAndRoleID(ctx context.Context, userID int) (string, int, *de.DomainError) {
	user, err := uc.pgAdp.GetUser(ctx, userID)
	if err != nil {
		return "", 0, err
	}

	return user.Login, user.RoleID, nil
}
