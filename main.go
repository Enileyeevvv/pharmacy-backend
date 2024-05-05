package main

import (
	"backend/database"
	"backend/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Init()

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	app.Listen(":9000")

}
