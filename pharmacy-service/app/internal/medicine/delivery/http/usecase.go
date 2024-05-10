package http

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"
)

type UseCase interface {
	FetchMedicinalProducts(ctx context.Context, limit, offset int) ([]usecase.MedicinalProduct, bool, *de.DomainError)
	CreateMedicine(ctx context.Context, medicine usecase.MedicinalProduct) *de.DomainError
}
