package config

import "time"

type Config struct {
	Postgres   Postgres
	Redis      Redis
	User       User
	HTTPServer HTTPServer
}

type Postgres struct {
	Host     string `validate:"required"`
	Port     int    `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DB       string `validate:"required"`
}

type Redis struct {
	Host string `validate:"required"`
	Port int    `validate:"required"`
}

type User struct {
	SessionTTL time.Duration `validate:"required"`
	Secret     string        `validate:"required"`
}

type HTTPServer struct {
	Port         int    `validate:"required"`
	AllowOrigins string `validate:"required"`
}
