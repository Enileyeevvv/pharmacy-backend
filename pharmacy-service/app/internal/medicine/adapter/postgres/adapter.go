package postgres

import (
	"context"
	"database/sql"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type adapter struct {
	db *sqlx.DB
}

func (a *adapter) FetchMedicinalProducts(
	ctx context.Context,
	limit, offset int,
) ([]MedicinalProduct, *de.DomainError) {
	var mps []MedicinalProduct
	err := a.db.SelectContext(ctx, &mps, queryFetchMedicines, limit, offset)

	if errors.Is(err, sql.ErrNoRows) {
		return []MedicinalProduct{}, nil
	}

	if err != nil {
		log.Error(err)
		return nil, de.ErrFetchMedicinalProducts
	}

	return mps, nil
}
