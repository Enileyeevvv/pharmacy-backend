package controllers

import (
	"backend/app/models"
	"backend/database"
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

	data, err := database.GetMedicinalProduct(getMedicinalProduct)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     err,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"limit":   getMedicinalProduct.Limit,
		"offset":  getMedicinalProduct.Offset,
		"hasNext": "",
		"data":    data,
	})
}

func CreateMedicinalProduct(ctx *fiber.Ctx) error {
	req := &models.CreateMedicinalProduct{}

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	medicinalProduct := &models.MedicinalProduct{}

	medicinalProduct.Name = req.Name
	medicinalProduct.Description = req.Description
	medicinalProduct.MaxQuantity = req.MaxQuantity

	if req.Quantity != nil {
		medicinalProduct.Quantity = *req.Quantity
	} else {
		medicinalProduct.Quantity = req.MaxQuantity
	}

	err := database.CreateMedicinalProduct(medicinalProduct)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success":          true,
		"message":          nil,
		"medicinalProduct": medicinalProduct,
	})
}

func UpdateMedicinalProduct(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": nil,
	})
}

func DeleteMedicinalProduct(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": nil,
	})
}
