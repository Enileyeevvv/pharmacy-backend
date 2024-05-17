package usecase

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
)

type PGAdapter interface {
	FetchMedicinalProducts(ctx context.Context, limit, offset int) ([]MedicinalProduct, *de.DomainError)
	CreateMedicinalProductTransaction(ctx context.Context, mp MedicinalProduct) *de.DomainError

	FetchPatients(ctx context.Context, limit, offset int, name *string) ([]Patient, *de.DomainError)
	GetPatient(ctx context.Context, id int) (Patient, *de.DomainError)
	FetchPrescriptions(
		ctx context.Context,
		limit, offset int,
		patientID *int,
		patientName *string,
	) ([]Prescription, *de.DomainError)
	GetPrescription(ctx context.Context, id int) (Prescription, *de.DomainError)
	CreatePrescriptionTransaction(ctx context.Context, p Prescription) *de.DomainError
	CheckoutPrescriptionTransaction(ctx context.Context, p Prescription) *de.DomainError
	FetchPrescriptionHistory(ctx context.Context, limit, offset, pID int) ([]PrescriptionHistory, *de.DomainError)
	AddMedicinalProductTransaction(ctx context.Context, mpID, quantity int) *de.DomainError
}
