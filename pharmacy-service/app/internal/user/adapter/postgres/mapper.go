package postgres

import "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/usecase"

func MapGetUser(user *User) *usecase.User {
	if user == nil {
		return nil
	}

	return &usecase.User{
		ID:           user.ID,
		Login:        user.Login,
		PasswordHash: user.Password,
		Status:       user.Status,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
