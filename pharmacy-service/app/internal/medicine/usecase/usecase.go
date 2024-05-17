package usecase

import (
	"context"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
)

type UseCase struct {
	pgAdp PGAdapter
}

func NewUseCase(pgAdp PGAdapter) *UseCase {
	return &UseCase{
		pgAdp: pgAdp,
	}
}

func (u *UseCase) FetchMedicinalProducts(
	ctx context.Context,
	limit, offset int,
) ([]MedicinalProduct, bool, *de.DomainError) {
	mps, err := u.pgAdp.FetchMedicinalProducts(ctx, limit, offset)
	if err != nil {
		return nil, false, err
	}

	hasNext := false
	if len(mps) > limit {
		hasNext = true
		mps = mps[:len(mps)-1]
	}

	return mps, hasNext, nil
}

func (u *UseCase) CreateMedicinalProduct(ctx context.Context, medicine MedicinalProduct) *de.DomainError {
	return u.pgAdp.CreateMedicinalProductTransaction(ctx, medicine)
}

func (u *UseCase) FetchPatients(
	ctx context.Context,
	limit, offset int,
	name *string,
) ([]Patient, bool, *de.DomainError) {
	ps, err := u.pgAdp.FetchPatients(ctx, limit, offset, name)
	if err != nil {
		return nil, false, err
	}

	hasNext := false
	if len(ps) > limit {
		hasNext = true
		ps = ps[:len(ps)-1]
	}

	return ps, hasNext, nil
}

func (u *UseCase) GetPatient(ctx context.Context, id int) (Patient, *de.DomainError) {
	p, err := u.pgAdp.GetPatient(ctx, id)
	if err != nil {
		return Patient{}, err
	}

	return p, nil
}

func (u *UseCase) FetchPrescriptions(
	ctx context.Context,
	limit, offset int,
	patientID *int,
	patientName *string,
) ([]Prescription, bool, *de.DomainError) {
	ps, err := u.pgAdp.FetchPrescriptions(ctx, limit, offset, patientID, patientName)
	if err != nil {
		return nil, false, err
	}

	hasNext := false
	if len(ps) > limit {
		hasNext = true
		ps = ps[:len(ps)-1]
	}

	return ps, hasNext, nil
}

func (u *UseCase) GetPrescription(ctx context.Context, id int) (Prescription, *de.DomainError) {
	p, err := u.pgAdp.GetPrescription(ctx, id)
	if err != nil {
		return Prescription{}, err
	}

	return p, nil
}

func (u *UseCase) CreatePrescription(ctx context.Context, p Prescription) *de.DomainError {
	return u.pgAdp.CreatePrescriptionTransaction(ctx, p)
}

func (u *UseCase) CheckoutPrescription(ctx context.Context, p Prescription) *de.DomainError {
	return u.pgAdp.CheckoutPrescriptionTransaction(ctx, p)
}

func (u *UseCase) FetchPrescriptionHistory(
	ctx context.Context,
	limit, offset, pID int,
) ([]PrescriptionHistory, bool, *de.DomainError) {
	ps, err := u.pgAdp.FetchPrescriptionHistory(ctx, limit, offset, pID)
	if err != nil {
		return nil, false, err
	}

	hasNext := false
	if len(ps) > limit {
		hasNext = true
		ps = ps[:len(ps)-1]
	}

	return ps, hasNext, nil
}

func (u *UseCase) AddMedicinalProduct(ctx context.Context, mpID, quantity int) *de.DomainError {
	return u.pgAdp.AddMedicinalProductTransaction(ctx, mpID, quantity)
}
