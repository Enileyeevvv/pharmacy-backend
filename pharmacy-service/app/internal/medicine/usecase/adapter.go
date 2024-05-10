package usecase

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
)

type PGAdapter interface {
	FetchMedicinalProducts(ctx context.Context, limit, offset int) ([]MedicinalProduct, *de.DomainError)
	CreateMedicinalProduct(ctx context.Context, mp MedicinalProduct) *de.DomainError
}
