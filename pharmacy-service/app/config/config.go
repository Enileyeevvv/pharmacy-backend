package config

type Config struct {
	Postgres   Postgres
	HTTPServer HTTPServer
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

type HTTPServer struct {
	Port         int
	AllowOrigins string
}
