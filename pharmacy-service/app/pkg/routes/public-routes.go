package routes

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	v1 := a.Group("/api/v1")

	userRoutes := v1.Group("/user")

	auth := userRoutes.Group("/sign")

	auth.Post("/up", controllers.UserSignUp)
	auth.Post("/in", controllers.UserSignIn)
	auth.Post("/out", controllers.UserSignOut)
}
