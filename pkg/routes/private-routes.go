package routes

import (
	"backend/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	v1 := a.Group("/api/v1")

	medProductRoutes := v1.Group("/medicinal_product")

	medProductRoutes.Get("/", controllers.GetMedicinalProductList)
	medProductRoutes.Post("/", controllers.CreateMedicinalProduct)
	medProductRoutes.Patch("/", controllers.UpdateMedicinalProduct)
	medProductRoutes.Delete("/", controllers.DeleteMedicinalProduct)
}
