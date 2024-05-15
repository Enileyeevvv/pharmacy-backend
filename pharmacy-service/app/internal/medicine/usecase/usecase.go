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
	mpID, err := u.pgAdp.CheckMedicinalProductExists(ctx, medicine)
	if err != nil {
		return err
	}

	cID, err := u.pgAdp.CheckCompanyExists(ctx, medicine)
	if err != nil {
		return err
	}

	if mpID == -1 {
		mpID, err = u.pgAdp.CreateMedicinalProduct(ctx, medicine)
		if err != nil {
			return err
		}
	}

	if cID == -1 {
		cID, err = u.pgAdp.CreateCompany(ctx, medicine)
		if err != nil {
			return err
		}
	}

	medicine.ID = mpID
	medicine.CompanyID = cID

	return u.pgAdp.UpsertMedicinalProductCompany(ctx, medicine)
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
) ([]Prescription, bool, *de.DomainError) {
	ps, err := u.pgAdp.FetchPrescriptions(ctx, limit, offset)
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
