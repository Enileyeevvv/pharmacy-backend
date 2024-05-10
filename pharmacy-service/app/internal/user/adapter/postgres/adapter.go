package postgres

import (
	"context"
	"database/sql"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/usecase"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type adapter struct {
	db *sqlx.DB
}

func NewAdapter(db *sqlx.DB) usecase.PGAdapter {
	return &adapter{
		db: db,
	}
}

func (a *adapter) CheckIfUserExists(ctx context.Context, login string) (bool, *de.DomainError) {
	var exists bool
	if err := a.db.GetContext(ctx, &exists, queryCheckIfUserExists, login); err != nil {
		log.Error(err)
		return false, de.ErrCheckIfUserExists
	}

	return exists, nil
}

func (a *adapter) CreateUser(ctx context.Context, login, passwordHash string) *de.DomainError {
	if _, err := a.db.ExecContext(ctx, queryCreateUser, login, passwordHash); err != nil {
		log.Error(err)
		return de.ErrCreateUser
	}

	return nil
}

func (a *adapter) GetPassword(ctx context.Context, login string) (string, *de.DomainError) {
	var password string
	err := a.db.GetContext(ctx, &password, queryGetPassword, login)

	if errors.Is(err, sql.ErrNoRows) {
		return "", de.ErrUserNotFound
	}

	if err != nil {
		return "", de.ErrGetPassword
	}

	return password, nil
}

func (a *adapter) GetUserID(ctx context.Context, login string) (int, *de.DomainError) {
	var userID int
	err := a.db.GetContext(ctx, &userID, queryGetUserID, login)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, de.ErrUserNotFound
	}

	if err != nil {
		return 0, de.ErrGetUserID
	}

	return userID, nil
}

func (a *adapter) GetUser(ctx context.Context, id int) (*usecase.User, *de.DomainError) {
	var user User
	err := a.db.GetContext(ctx, &user, queryGetUser, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		log.Error(err)
		return nil, de.ErrGetUser
	}

	return MapGetUser(&user), nil
}
