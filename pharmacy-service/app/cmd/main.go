package main

import (
	"fmt"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/config"
	postgres2 "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/adapter/postgres"
	http2 "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/delivery/http"
	usecase2 "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/adapter/postgres"
	redis2 "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/adapter/redis"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/delivery/http"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user/usecase"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/layers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
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

func initRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
	})
}

func initAdapters(cfg *config.Config) *layers.Adapters {
	db := initDB(cfg)
	r := initRedis(cfg)

	userPGAdp := postgres.NewAdapter(db)
	userRedisAdp := redis2.NewAdapter(r)
	medicinePGAdp := postgres2.NewAdapter(db)

	return &layers.Adapters{
		UserPGAdp:         userPGAdp,
		UserRedisAdp:      userRedisAdp,
		MedicinePGAdapter: medicinePGAdp,
	}
}

func initUseCases(cfg *config.Config) *layers.UseCases {
	adp := initAdapters(cfg)

	userUC := usecase.NewUseCase(
		adp.UserPGAdp,
		adp.UserRedisAdp,
		cfg.User.SessionTTL,
		cfg.User.Secret,
	)

	medicineUC := usecase2.NewUseCase(adp.MedicinePGAdapter)

	return &layers.UseCases{
		UserUC:     userUC,
		MedicineUC: medicineUC,
	}
}

func initHandlers(cfg *config.Config) *layers.Handlers {
	uc := initUseCases(cfg)

	userH := http.NewHandler(uc.UserUC)
	medicineH := http2.NewHandler(uc.MedicineUC)

	return &layers.Handlers{
		UserH:     userH,
		MedicineH: medicineH,
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

	app.Use(logger.New())

	http.MapUserRoots(app, handlers.UserH)
	http2.MapMedicineRoots(app, handlers.UserH, handlers.MedicineH)

	return app
}

func runHTTPServer(cfg *config.Config) error {
	app := initHTTPServer(cfg)
	return app.Listen(fmt.Sprintf(":%d", cfg.HTTPServer.Port))
}
