package main

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/database"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	database.Init()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "http://localhost:3005",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	app.Listen(":9000")

}
