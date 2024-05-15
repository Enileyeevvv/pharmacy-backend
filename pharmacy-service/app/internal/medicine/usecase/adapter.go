package usecase

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
)

type PGAdapter interface {
	FetchMedicinalProducts(ctx context.Context, limit, offset int) ([]MedicinalProduct, *de.DomainError)
	CheckMedicinalProductExists(ctx context.Context, mp MedicinalProduct) (int, *de.DomainError)
	CheckCompanyExists(ctx context.Context, mp MedicinalProduct) (int, *de.DomainError)
	CreateMedicinalProduct(ctx context.Context, mp MedicinalProduct) (int, *de.DomainError)
	CreateCompany(ctx context.Context, mp MedicinalProduct) (int, *de.DomainError)
	UpsertMedicinalProductCompany(ctx context.Context, mp MedicinalProduct) *de.DomainError

	FetchPatients(ctx context.Context, limit, offset int, name *string) ([]Patient, *de.DomainError)
	GetPatient(ctx context.Context, id int) (Patient, *de.DomainError)
}
