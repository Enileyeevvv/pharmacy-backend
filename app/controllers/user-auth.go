package controllers

import (
	"backend/app/models"
	"backend/database"
	"backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

func UserSignUp(ctx *fiber.Ctx) error {
	signUp := &models.SignUp{}

	if err := ctx.BodyParser(signUp); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(signUp); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     utils.ValidatorErrors(err),
		})
	}

	var isUserSignedUp = database.IsUserSignedUpByLogin(signUp.Login)

	if isUserSignedUp {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     "This Login is already signed up",
		})
	}

	db, err := database.OpenConnection()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}

	user := &models.User{}

	// Set initialized default data for user:
	user.ID = 2
	user.Login = signUp.Login
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	user.Password = utils.GeneratePassword(signUp.Password)
	user.Status = 1 // 0 == blocked, 1 == active

	if err := validate.Struct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     utils.ValidatorErrors(err),
		})
	}

	if err := db.CreateUser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"check":   true,
			"success": false,
			"msg":     err.Error(),
		})
	}

	user.Password = ""

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": nil,
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
