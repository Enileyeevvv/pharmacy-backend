package queries

import (
	"backend/app/models"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) CreateUser(u *models.User) error {
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := q.Exec(
		query,
		u.ID, u.Login, u.Password, u.Status, u.CreatedAt, u.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
