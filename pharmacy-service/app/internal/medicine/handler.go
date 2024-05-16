package medicine

import "github.com/gofiber/fiber/v2"

type Handler interface {
	FetchMedicinalProducts() fiber.Handler
	CreateMedicinalProduct() fiber.Handler
	FetchPatients() fiber.Handler
	GetPatient() fiber.Handler
	FetchPrescriptions() fiber.Handler
	GetPrescription() fiber.Handler
	CreateSinglePrescription() fiber.Handler
	CreateMultiplePrescription() fiber.Handler
	SubmitPrescription() fiber.Handler
	CancelPrescription() fiber.Handler
}
