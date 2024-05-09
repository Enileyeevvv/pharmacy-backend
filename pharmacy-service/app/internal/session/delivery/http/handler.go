package http

import "github.com/gofiber/fiber/v2"

type handler struct {
}

func (h *handler) CreateSession() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
