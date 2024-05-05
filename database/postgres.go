package database

import (
	"backend/pkg/utils"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

func PostgreSQLConnection() (*sqlx.DB, error) {
	maxConn, _ := strconv.Atoi("100")
	maxIdleConn, _ := strconv.Atoi("10")
	maxLifetimeConn, _ := strconv.Atoi("10")

	postgresConnURL, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("pgx", postgresConnURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
