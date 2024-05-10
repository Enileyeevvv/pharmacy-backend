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
