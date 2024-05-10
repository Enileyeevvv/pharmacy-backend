package postgres

import (
	"context"
	"database/sql"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type adapter struct {
	db *sqlx.DB
}

func NewAdapter(db *sqlx.DB) usecase.PGAdapter {
	return &adapter{
		db: db,
	}
}

func (a *adapter) FetchMedicinalProducts(
	ctx context.Context,
	limit, offset int,
) ([]usecase.MedicinalProduct, *de.DomainError) {
	var mps []MedicinalProduct
	err := a.db.SelectContext(ctx, &mps, queryFetchMedicines, limit, offset)

	if errors.Is(err, sql.ErrNoRows) {
		return make([]usecase.MedicinalProduct, 0), nil
	}

	if err != nil {
		log.Error(err)
		return nil, de.ErrFetchMedicinalProducts
	}

	return MapMedicinalProductSlice(mps), nil
}

func (a *adapter) CreateMedicinalProduct(ctx context.Context, mp usecase.MedicinalProduct) *de.DomainError {
	_, err := a.db.ExecContext(
		ctx,
		queryCreateMedicalProduct,
		mp.Name,
		mp.SellName,
		mp.ATXCode,
		mp.Description,
		mp.Quantity,
		mp.MaxQuantity)
	if err != nil {
		log.Error(err)
		return de.ErrCreateMedicalProduct
	}

	return nil
}
