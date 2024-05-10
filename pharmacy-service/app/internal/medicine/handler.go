package medicine

import "github.com/gofiber/fiber/v2"

type Handler interface {
	FetchMedicinalProducts() fiber.Handler
}
