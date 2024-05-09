package user

import "github.com/gofiber/fiber/v2"

type Handler interface {
	UserSignUp() fiber.Handler
	UserSignIn() fiber.Handler
	UserSignOut() fiber.Handler
}
