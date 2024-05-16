package http

import (
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type handler struct {
	uc UseCase
	v  *validator.Validate
}

func NewHandler(uc UseCase) medicine.Handler {
	return &handler{
		uc: uc,
		v:  validator.New(),
	}
}

func (h *handler) FetchMedicinalProducts() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req FetchMedicinalProductsRequest

		if err := ctx.QueryParser(&req); err != nil {
			return de.ErrParseRequestBody.ToHTTPError(ctx)
		}

		if err := h.v.Struct(&req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		mps, hasNext, dErr := h.uc.FetchMedicinalProducts(ctx.Context(), req.Limit, req.Offset)
		if dErr != nil {
			return dErr.ToHTTPError(ctx)
		}

		return ctx.Status(fiber.StatusOK).JSON(MapFetchMedicinalProductsResponse(mps, hasNext))
	}
}

func (h *handler) CreateMedicinalProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req CreateMedicinalProductRequest

		if err := ctx.BodyParser(&req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		if err := h.v.Struct(&req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		err := h.uc.CreateMedicinalProduct(ctx.Context(), MapCreateMedicinalProductRequest(req))
		if err != nil {
			return err.ToHTTPError(ctx)
		}

		return de.OK.ToHTTPError(ctx)
	}
}

func (h *handler) FetchPatients() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req FetchPatientsRequest

		if err := ctx.QueryParser(&req); err != nil {
			return de.ErrParseRequestBody.ToHTTPError(ctx)
		}

		if err := h.v.Struct(&req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		ps, hasNext, dErr := h.uc.FetchPatients(ctx.Context(), req.Limit, req.Offset, req.Name)
		if dErr != nil {
			return dErr.ToHTTPError(ctx)
		}

		return ctx.Status(fiber.StatusOK).JSON(MapFetchPatientsResponse(ps, hasNext))
	}
}

func (h *handler) GetPatient() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		patientIDParam := ctx.Params("id")
		patientID, err := strconv.Atoi(patientIDParam)
		if err != nil {
			return de.ErrIncorrectPathParam.ToHTTPError(ctx)
		}

		p, dErr := h.uc.GetPatient(ctx.Context(), patientID)
		if dErr != nil {
			return dErr.ToHTTPError(ctx)
		}

		return ctx.Status(fiber.StatusOK).JSON(MapGetPatientResponse(p))
	}
}

func (h *handler) FetchPrescriptions() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req FetchPrescriptionsRequest

		if err := ctx.QueryParser(&req); err != nil {
			return de.ErrParseRequestBody.ToHTTPError(ctx)
		}

		if err := h.v.Struct(&req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		ps, hasNext, dErr := h.uc.FetchPrescriptions(ctx.Context(), req.Limit, req.Offset)
		if dErr != nil {
			return dErr.ToHTTPError(ctx)
		}

		return ctx.Status(fiber.StatusOK).JSON(MapFetchPrescriptionsResponse(ps, hasNext))
	}
}

func (h *handler) GetPrescription() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		prescriptionIDParam := ctx.Params("id")
		prescriptionID, err := strconv.Atoi(prescriptionIDParam)
		if err != nil {
			return de.ErrIncorrectPathParam.ToHTTPError(ctx)
		}

		p, dErr := h.uc.GetPrescription(ctx.Context(), prescriptionID)
		if dErr != nil {
			return dErr.ToHTTPError(ctx)
		}

		return ctx.Status(fiber.StatusOK).JSON(MapGetPrescriptionResponse(p))
	}
}

func (h *handler) CreateSinglePrescription() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req CreateSinglePrescriptionRequest

		if err := ctx.BodyParser(&req); err != nil {
			return de.ErrParseRequestBody.ToHTTPError(ctx)
		}

		if err := h.v.Struct(&req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		userID, ok := (ctx.Locals("userID")).(int)
		if !ok {
			return de.ErrInvalidUserID.ToHTTPError(ctx)
		}

		dErr := h.uc.CreatePrescription(ctx.Context(), MapCreateSinglePrescriptionRequest(req, userID))
		if dErr != nil {
			return dErr.ToHTTPError(ctx)
		}

		return de.OK.ToHTTPError(ctx)
	}
}

func (h *handler) CreateMultiplePrescription() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req CreateMultiplePrescriptionRequest

		if err := ctx.BodyParser(&req); err != nil {
			return de.ErrParseRequestBody.ToHTTPError(ctx)
		}

		if err := h.v.Struct(&req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		userID, ok := (ctx.Locals("userID")).(int)
		if !ok {
			return de.ErrInvalidUserID.ToHTTPError(ctx)
		}

		dErr := h.uc.CreatePrescription(ctx.Context(), MapCreateMultiplePrescriptionRequest(req, userID))
		if dErr != nil {
			return dErr.ToHTTPError(ctx)
		}

		return de.OK.ToHTTPError(ctx)
	}
}

func (h *handler) SubmitPrescription() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req SubmitPrescriptionRequest

		if err := ctx.BodyParser(&req); err != nil {
			return de.ErrParseRequestBody.ToHTTPError(ctx)
		}

		if err := h.v.Struct(&req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		userID, ok := (ctx.Locals("userID")).(int)
		if !ok {
			return de.ErrInvalidUserID.ToHTTPError(ctx)
		}

		dErr := h.uc.CheckoutPrescription(ctx.Context(), MapSubmitPrescriptionRequest(req, userID))
		if dErr != nil {
			return dErr.ToHTTPError(ctx)
		}

		return de.OK.ToHTTPError(ctx)
	}
}

func (h *handler) CancelPrescription() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req CancelPrescriptionRequest

		if err := ctx.BodyParser(&req); err != nil {
			return de.ErrParseRequestBody.ToHTTPError(ctx)
		}

		if err := h.v.Struct(&req); err != nil {
			return de.ErrRequestBodyInvalid.ToHTTPError(ctx)
		}

		userID, ok := (ctx.Locals("userID")).(int)
		if !ok {
			return de.ErrInvalidUserID.ToHTTPError(ctx)
		}

		dErr := h.uc.CheckoutPrescription(ctx.Context(), MapCancelPrescriptionRequest(req, userID))
		if dErr != nil {
			return dErr.ToHTTPError(ctx)
		}

		return de.OK.ToHTTPError(ctx)
	}
}
