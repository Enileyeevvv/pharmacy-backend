package http

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"
)

type UseCase interface {
	FetchMedicinalProducts(ctx context.Context, limit, offset int) ([]usecase.MedicinalProduct, bool, *de.DomainError)
	CreateMedicinalProduct(ctx context.Context, medicine usecase.MedicinalProduct) *de.DomainError
	FetchPatients(ctx context.Context, limit, offset int, name *string) ([]usecase.Patient, bool, *de.DomainError)
	GetPatient(ctx context.Context, id int) (usecase.Patient, *de.DomainError)
	FetchPrescriptions(ctx context.Context, limit, offset int) ([]usecase.Prescription, bool, *de.DomainError)
	GetPrescription(ctx context.Context, id int) (usecase.Prescription, *de.DomainError)
}
