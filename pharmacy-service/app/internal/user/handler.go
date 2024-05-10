package user

import (
	"github.com/gofiber/fiber/v2"
)

type Role uint8

const (
	DOCTOR Role = iota + 1
	PHARMACIST
	ADMIN
)

type Handler interface {
	UserSignUp() fiber.Handler
	UserSignIn() fiber.Handler
	UserSignOut() fiber.Handler
	AuthMW() fiber.Handler
	RoleMW(role Role) fiber.Handler
}
