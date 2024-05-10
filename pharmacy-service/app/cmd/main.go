package main

import (
	"fmt"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/config"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/adapter/postgres"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/delivery/http"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/usecase"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/layers"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/pkg/routes"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := initConfig()
	if err := runHTTPServer(cfg); err != nil {
		log.Fatal(err)
	}
}

func initConfig() *config.Config {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config.json")
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var c config.Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}
	err = validator.New().Struct(c)
	if err != nil {
		log.Fatal(err)
	}

	return &c
}

func initDB(cfg *config.Config) *sqlx.DB {
	dataSource := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DB)
	db, err := sqlx.Connect("pgx", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func initAdapters(cfg *config.Config) *layers.Adapters {
	db := initDB(cfg)

	userPGAdp := postgres.NewAdapter(db)

	return &layers.Adapters{
		UserPGAdp: userPGAdp,
	}
}

func initUseCases(cfg *config.Config) *layers.UseCases {
	adp := initAdapters(cfg)

	userUC := usecase.NewUseCase(adp.UserPGAdp)

	return &layers.UseCases{
		UserUC: userUC,
	}
}

func initHandlers(cfg *config.Config) *layers.Handlers {
	uc := initUseCases(cfg)

	userH := http.NewHandler(uc.UserUC)

	return &layers.Handlers{
		UserH: userH,
	}
}

func initHTTPServer(cfg *config.Config) *fiber.App {
	handlers := initHandlers(cfg)

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     cfg.HTTPServer.AllowOrigins,
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	http.MapUserRoots(app, handlers.UserH)

	return app
}

func runHTTPServer(cfg *config.Config) error {
	app := initHTTPServer(cfg)
	return app.Listen(fmt.Sprintf(":%d", cfg.HTTPServer.Port))
}
