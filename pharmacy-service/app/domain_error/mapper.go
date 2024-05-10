package domain_error

import "github.com/gofiber/fiber/v2"

func (de *DomainError) ToHTTPError(ctx *fiber.Ctx) error {
	errorBody := struct {
		Message string         `json:"message"`
		Params  map[string]any `json:"params,omitempty"`
	}{
		Message: de.message,
		Params:  de.params,
	}
	return ctx.Status(de.code.ToHTTPCode()).JSON(errorBody)
}
