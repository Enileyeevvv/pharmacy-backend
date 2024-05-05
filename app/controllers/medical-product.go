package controllers

import (
	"backend/app/models"
	"backend/database"
	"backend/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetMedicinalProductList(ctx *fiber.Ctx) error {
	getMedicinalProduct := &models.GetMedicinalProduct{}

	if err := ctx.QueryParser(getMedicinalProduct); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": nil,
	})
}

func CreateMedicinalProduct(ctx *fiber.Ctx) error {
	medicinalProduct := &models.CreateMedicinalProduct{}

	if err := ctx.BodyParser(medicinalProduct); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	fmt.Println(medicinalProduct)

	new := &models.MedicinalProduct{Name: "medicinalProduct", Description: "medicinalProduct"}

	result := database.GetDB().Create(new)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     utils.ValidatorErrors(result.Error),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": nil,
	})
}
