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
	medProductRoutes.Post("/", uH.AuthMW(), uH.RoleMW(user.ADMIN), mH.CreateMedicinalProduct())

	patient := v1.Group("/patient")
	patient.Get("/", mH.FetchPatients())
	patient.Get("/:id", mH.GetPatient())

	prescription := v1.Group("/prescription")
	prescription.Get("/", mH.FetchPrescriptions())
	prescription.Get("/:id", mH.GetPrescription())
	prescription.Post("/single/create", uH.AuthMW(), uH.RoleMW(user.DOCTOR), mH.CreateSinglePrescription())
	prescription.Post("/multiple/create", uH.AuthMW(), uH.RoleMW(user.DOCTOR), mH.CreateMultiplePrescription())
	prescription.Post("/submit", uH.AuthMW(), uH.RoleMW(user.PHARMACIST), mH.SubmitPrescription())
	prescription.Post("/cancel", uH.AuthMW(), uH.RoleMW(user.PHARMACIST), mH.CancelPrescription())
}
