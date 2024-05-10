package http

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	v  *validator.Validate
	uc UseCase
}

func NewHandler(uc UseCase) user.Handler {
	return &handler{
		uc: uc,
		v:  validator.New(),
	}
}

func (h *handler) UserSignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req SignUpRequest

		if err := ctx.BodyParser(&req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(SignUpResponse{
				Success: false,
				Message: err.Error(),
			})
		}

		if err := h.v.Struct(req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(SignUpResponse{
				Success: false,
				Msg:     utils.ValidatorErrors(err),
			})
		}

		err := h.uc.SignUp(ctx.Context(), req.Login, req.Password)
		if err != nil {
			return ctx.Status(err.Code().ToHTTPCode()).JSON(SignUpResponse{
				Success: false,
				Message: err.Message(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(SignUpResponse{
			Success: true,
		})
	}
}

//func (h *handler) UserSignIn() fiber.Handler {
//	return func(ctx *fiber.Ctx) error {
//		return ctx.Status(200).JSON(fiber.Map{
//			"success": true,
//			"message": "Hello world",
//		})
//	}
//}

//func (h *handler) UserSignOut() fiber.Handler {
//	return func(ctx *fiber.Ctx) error {
//		ctx.ClearCookie("access-token")
//		ctx.ClearCookie("refresh-token")
//		return ctx.Status(fiber.StatusOK).JSON(// todo)
//	}
//}
