package http

import (
	"github.com/Enileyeevvv/pharmacy-backend/internal/user"
	"github.com/gofiber/fiber/v2"
)

func MapUserRoots(a *fiber.App, h user.Handler) {
	v1 := a.Group("/api/v1")

	userRoutes := v1.Group("/user")

	auth := userRoutes.Group("/sign")

	auth.Post("/up", h.UserSignUp())
	auth.Post("/in", h.UserSignIn())
	auth.Post("/out", h.UserSignOut())
}
