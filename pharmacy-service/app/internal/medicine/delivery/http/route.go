package http

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/user"
	"github.com/gofiber/fiber/v2"
)

func MapMedicineRoots(a *fiber.App, uH user.Handler, mH medicine.Handler) {
	v1 := a.Group("/api/v1")

	medProductRoutes := v1.Group("/medicinal_product")

	medProductRoutes.Get("/", uH.AuthMW(), mH.FetchMedicinalProducts())
	//medProductRoutes.Post("/", controllers.CreateMedicinalProduct)
	//medProductRoutes.Patch("/", controllers.UpdateMedicinalProduct)
	//medProductRoutes.Delete("/", controllers.DeleteMedicinalProduct)
}
