package http

import (
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

		m := usecase.MedicinalProduct{
			Name:        req.Name,
			SellName:    req.SellName,
			ATXCode:     req.ATXCode,
			Description: req.Description,
			Quantity:    req.Quantity,
			MaxQuantity: req.MaxQuantity,
		}

		err := h.uc.CreateMedicine(ctx.Context(), m)
		if err != nil {
			return err.ToHTTPError(ctx)
		}

		return de.OK.ToHTTPError(ctx)
	}
}
