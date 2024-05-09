package http

import (
	"github.com/Enileyeevvv/pharmacy-backend/database"
	"github.com/Enileyeevvv/pharmacy-backend/models"
	"github.com/Enileyeevvv/pharmacy-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

type handler struct {
}

func (h *handler) UserSignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := &models.SignUp{}
		if err := ctx.BodyParser(req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		validate := utils.NewValidator()
		if err := validate.Struct(req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"msg":     utils.ValidatorErrors(err),
			})
		}

		isUserSignedUp := database.IsUserCreatedByLogin(req.Login)
		if isUserSignedUp {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"msg":     "This login is already signed up",
			})
		}

		user := &models.User{}

		user.Login = req.Login
		user.CreatedAt = time.Now().Unix()
		user.UpdatedAt = time.Now().Unix()
		user.Password = utils.GeneratePassword(req.Password)
		user.Status = 1 // 0 == blocked, 1 == active

		err := database.CreateUser(user)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"msg":     err.Error(),
			})
		}

		user.Password = ""

		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"user":    user,
		})
	}
}

func (h *handler) UserSignIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Hello world",
		})
	}
}

func (h *handler) UserSignOut() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Hello world",
		})
	}
}
