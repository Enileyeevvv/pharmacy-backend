package database

import (
	"backend/app/queries"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*queries.UserQueries
}

func OpenConnection() (*Queries, error) {
	var db *sqlx.DB
	var err error

	db, err = PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	return &Queries{
		UserQueries: &queries.UserQueries{DB: db}, // from User model
	}, nil
}
