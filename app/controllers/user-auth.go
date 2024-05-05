package controllers

import (
	"backend/app/models"
	"backend/database"
	"backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

func UserSignUp(ctx *fiber.Ctx) error {
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

func UserSignIn(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
	})
}

func UserSignOut(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Hello world",
	})
}
