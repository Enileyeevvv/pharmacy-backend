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
	FetchPrescriptions(
		ctx context.Context,
		limit, offset int,
		patientID *int,
		patientName *string,
	) ([]usecase.Prescription, bool, *de.DomainError)
	GetPrescription(ctx context.Context, id int) (usecase.Prescription, *de.DomainError)
	CreatePrescription(ctx context.Context, p usecase.Prescription) *de.DomainError
	CheckoutPrescription(ctx context.Context, p usecase.Prescription) *de.DomainError
	FetchPrescriptionHistory(
		ctx context.Context,
		limit, offset, pID int,
	) ([]usecase.PrescriptionHistory, bool, *de.DomainError)
	AddMedicinalProduct(ctx context.Context, mpID, quantity int) *de.DomainError
}
