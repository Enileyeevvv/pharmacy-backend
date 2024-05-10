package domain_code

import "github.com/gofiber/fiber/v2"

type DomainCode uint16

const (
	OK = 200

	BadRequest   = 400
	Unauthorized = 401
	Forbidden    = 403
	NotFound     = 404

	Internal      = 500
	Unimplemented = 501
)

func (dc DomainCode) ToHTTPCode() int {
	switch dc {
	case OK:
		return fiber.StatusOK

	case BadRequest:
		return fiber.StatusBadRequest

	case Unauthorized:
		return fiber.StatusUnauthorized

	case Forbidden:
		return fiber.StatusForbidden

	case NotFound:
		return fiber.StatusNotFound

	case Internal:
		return fiber.StatusInternalServerError

	case Unimplemented:
		return fiber.StatusNotImplemented

	default:
		return fiber.StatusInternalServerError
	}
}
