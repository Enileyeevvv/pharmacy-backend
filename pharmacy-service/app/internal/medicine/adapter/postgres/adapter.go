package postgres

import (
	"context"
	"database/sql"
	de "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/domain_error"
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
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

func (a *adapter) CreateMedicinalProductTransaction(ctx context.Context, mp usecase.MedicinalProduct) *de.DomainError {
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error(err)
		return de.ErrCreateTransaction
	}

	defer func(tx *sqlx.Tx) {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Error(err)
		}
	}(tx)

	mpID, dErr := a.checkMedicinalProductExists(ctx, tx, mp)
	if dErr != nil {
		return dErr
	}
	if mpID == -1 {
		mpID, dErr = a.createMedicinalProduct(ctx, tx, mp)
		if dErr != nil {
			return dErr
		}
	}

	cID, dErr := a.checkCompanyExists(ctx, tx, mp)
	if dErr != nil {
		return dErr
	}
	if cID == -1 {
		cID, dErr = a.createCompany(ctx, tx, mp)
		if dErr != nil {
			return dErr
		}
	}

	mp.ID = mpID
	mp.CompanyID = cID

	dErr = a.upsertMedicinalProductCompany(ctx, tx, mp)
	if dErr != nil {
		return dErr
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return de.ErrCommitTransaction
	}

	return nil
}

func (a *adapter) checkMedicinalProductExists(
	ctx context.Context,
	tx *sqlx.Tx,
	mp usecase.MedicinalProduct,
) (int, *de.DomainError) {
	var mpID int
	err := tx.GetContext(
		ctx,
		&mpID,
		queryCheckMedicinalProductExists,
		strings.TrimSpace(strings.ToLower(mp.Name)),
		strings.TrimSpace(strings.ToLower(mp.SellName)))

	if errors.Is(err, sql.ErrNoRows) {
		return -1, nil
	}

	if err != nil {
		log.Error(err)
		return 0, de.ErrCheckMedicinalProductExists
	}

	return mpID, nil
}

func (a *adapter) checkCompanyExists(
	ctx context.Context,
	tx *sqlx.Tx,
	mp usecase.MedicinalProduct,
) (int, *de.DomainError) {
	var cID int
	err := tx.GetContext(
		ctx,
		&cID,
		queryCheckCompanyExists,
		strings.TrimSpace(strings.ToLower(mp.CompanyName)))

	if errors.Is(err, sql.ErrNoRows) {
		return -1, nil
	}

	if err != nil {
		log.Error(err)
		return 0, de.ErrCheckCompanyExists
	}

	return cID, nil
}

func (a *adapter) createMedicinalProduct(
	ctx context.Context,
	tx *sqlx.Tx,
	mp usecase.MedicinalProduct,
) (int, *de.DomainError) {
	var mpID int

	err := tx.GetContext(
		ctx,
		&mpID,
		queryCreateMedicalProduct,
		mp.Name,
		mp.SellName,
		mp.ATXCode,
		mp.Description,
		mp.Quantity,
		mp.MaxQuantity,
		mp.PharmaceuticalGroupID)
	if err != nil {
		log.Error(err)
		return 0, de.ErrCreateMedicalProduct
	}

	return mpID, nil
}

func (a *adapter) createCompany(ctx context.Context, tx *sqlx.Tx, mp usecase.MedicinalProduct) (int, *de.DomainError) {
	var cID int

	err := tx.GetContext(
		ctx,
		&cID,
		queryCreateCompany,
		mp.CompanyName)
	if err != nil {
		log.Error(err)
		return 0, de.ErrCreateCompany
	}

	return cID, nil
}

func (a *adapter) upsertMedicinalProductCompany(
	ctx context.Context,
	tx *sqlx.Tx,
	mp usecase.MedicinalProduct,
) *de.DomainError {
	_, err := tx.ExecContext(
		ctx,
		queryUpsertMedicinalProductCompany,
		mp.ID,
		mp.CompanyID,
		mp.ImageURL,
		mp.DosageFormID)
	if err != nil {
		log.Error(err)
		return de.ErrUpsertMedicinalProductCompany
	}

	return nil
}

func (a *adapter) FetchPatients(ctx context.Context, limit, offset int, name *string) ([]usecase.Patient, *de.DomainError) {
	patients := make([]Patient, 0)

	err := a.db.SelectContext(ctx, &patients, queryFetchPatients, limit, offset, name)

	if errors.Is(err, sql.ErrNoRows) {
		return make([]usecase.Patient, 0), nil
	}

	if err != nil {
		log.Error(err)
		return nil, de.ErrFetchPatients
	}

	return MapPatientSlice(patients), nil
}

func (a *adapter) GetPatient(ctx context.Context, id int) (usecase.Patient, *de.DomainError) {
	var patient Patient

	err := a.db.GetContext(ctx, &patient, queryGetPatient, id)

	if errors.Is(err, sql.ErrNoRows) {
		return usecase.Patient{}, nil
	}

	if err != nil {
		log.Error(err)
		return usecase.Patient{}, de.ErrGetPatient
	}

	return MapPatient(patient), nil
}

func (a *adapter) FetchPrescriptions(
	ctx context.Context,
	limit, offset int,
	patientID *int,
	patientName *string,
) ([]usecase.Prescription, *de.DomainError) {
	var ps []Prescription

	err := a.db.SelectContext(
		ctx,
		&ps,
		queryFetchPrescriptions,
		limit,
		offset,
		patientID,
		patientName)

	if errors.Is(err, sql.ErrNoRows) {
		return make([]usecase.Prescription, 0), nil
	}

	if err != nil {
		log.Error(err)
		return nil, de.ErrFetchPrescriptions
	}

	return MapPrescriptions(ps), nil
}

func (a *adapter) GetPrescription(ctx context.Context, id int) (usecase.Prescription, *de.DomainError) {
	var p Prescription

	err := a.db.GetContext(ctx, &p, queryGetPrescription, id)

	if errors.Is(err, sql.ErrNoRows) {
		return usecase.Prescription{}, nil
	}

	if err != nil {
		log.Error(err)
		return usecase.Prescription{}, de.ErrGetPrescription
	}

	return MapPrescription(p), nil
}

func (a *adapter) CreatePrescriptionTransaction(ctx context.Context, p usecase.Prescription) *de.DomainError {
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error(err)
		return de.ErrCreateTransaction
	}

	defer func(tx *sqlx.Tx) {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Error(err)
		}
	}(tx)

	res, dErr := a.createPrescription(ctx, tx, p)
	if dErr != nil {
		return dErr
	}

	dErr = a.updatePrescriptionHistory(ctx, tx, res)
	if dErr != nil {
		return dErr
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return de.ErrCommitTransaction
	}

	return nil
}

func (a *adapter) createPrescription(
	ctx context.Context,
	tx *sqlx.Tx,
	p usecase.Prescription,
) (Prescription, *de.DomainError) {
	var res Prescription

	err := tx.GetContext(
		ctx,
		&res,
		queryCreatePrescription,
		p.StampID,
		p.TypeID,
		p.MedicinalProductID,
		p.MedicinalProductQuantity,
		p.DoctorID,
		p.PatientID,
		p.ExpiredAt)
	if err != nil {
		log.Error(err)
		return Prescription{}, de.ErrCreatePrescription
	}

	return res, nil
}

func (a *adapter) CheckoutPrescriptionTransaction(ctx context.Context, p usecase.Prescription) *de.DomainError {
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error(err)
		return de.ErrCreateTransaction
	}

	defer func(tx *sqlx.Tx) {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Error(err)
		}
	}(tx)

	res, dErr := a.checkoutPrescription(ctx, tx, p)
	if dErr != nil {
		return dErr
	}

	if res.StatusID == 3 {
		mp, dErr := a.getMedicinalProduct(ctx, tx, res.MedicinalProductID)
		if dErr != nil {
			return dErr
		}

		dErr = a.subtractMedicinalProduct(ctx, tx, mp, res.MedicinalProductQuantity)
		if dErr != nil {
			return dErr
		}
	}

	dErr = a.updatePrescriptionHistory(ctx, tx, res)
	if dErr != nil {
		return dErr
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return de.ErrCommitTransaction
	}

	return nil
}

func (a *adapter) checkoutPrescription(
	ctx context.Context,
	tx *sqlx.Tx,
	p usecase.Prescription,
) (Prescription, *de.DomainError) {
	var res Prescription

	err := tx.GetContext(
		ctx,
		&res,
		queryCheckoutPrescription,
		p.ID,
		p.PharmacistID,
		p.StatusID)
	if err != nil {
		log.Error(err)
		return Prescription{}, de.ErrCheckoutPrescription
	}

	return res, nil
}

func (a *adapter) updatePrescriptionHistory(ctx context.Context, tx *sqlx.Tx, p Prescription) *de.DomainError {
	_, err := tx.ExecContext(
		ctx,
		queryUpdatePrescriptionHistory,
		p.ID,
		p.StatusID,
		p.DoctorID,
		p.PharmacistID)
	if err != nil {
		log.Error(err)
		return de.ErrUpdatePrescriptionHistory
	}

	return nil
}

func (a *adapter) FetchPrescriptionHistory(
	ctx context.Context,
	limit, offset, pID int,
) ([]usecase.PrescriptionHistory, *de.DomainError) {
	var ps []PrescriptionHistory

	err := a.db.SelectContext(ctx, &ps, queryFetchPrescriptionHistory, limit, offset, pID)

	if errors.Is(err, sql.ErrNoRows) {
		return make([]usecase.PrescriptionHistory, 0), nil
	}

	if err != nil {
		log.Error(err)
		return nil, de.ErrFetchPrescriptionHistory
	}

	return MapPrescriptionHistory(ps), nil
}

func (a *adapter) getMedicinalProduct(ctx context.Context, tx *sqlx.Tx, mpID int) (MedicinalProduct, *de.DomainError) {
	var mp MedicinalProduct
	err := tx.GetContext(ctx, &mp, queryGetMedicinalProduct, mpID)

	if errors.Is(err, sql.ErrNoRows) {
		return MedicinalProduct{}, nil
	}

	if err != nil {
		log.Error(err)
		return MedicinalProduct{}, de.ErrGetMedicinalProduct
	}

	return mp, nil
}

func (a *adapter) AddMedicinalProductTransaction(ctx context.Context, mpID, quantity int) *de.DomainError {
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error(err)
		return de.ErrCreateTransaction
	}

	defer func(tx *sqlx.Tx) {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Error(err)
		}
	}(tx)

	mp, dErr := a.getMedicinalProduct(ctx, tx, mpID)
	if dErr != nil {
		return dErr
	}

	dErr = a.addMedicinalProduct(ctx, tx, mp, quantity)
	if dErr != nil {
		return dErr
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return de.ErrCommitTransaction
	}

	return nil
}

func (a *adapter) addMedicinalProduct(
	ctx context.Context,
	tx *sqlx.Tx,
	mp MedicinalProduct,
	quantity int,
) *de.DomainError {
	if mp.Quantity+quantity > mp.MaxQuantity {
		return de.ErrQuantityTooHigh.WithParams("max", mp.MaxQuantity-mp.Quantity)
	}

	_, err := tx.ExecContext(ctx, queryAddMedicinalProduct, mp.ID, quantity)
	if err != nil {
		log.Error(err)
		return de.ErrAddMedicinalProduct
	}

	return nil
}

func (a *adapter) subtractMedicinalProduct(
	ctx context.Context,
	tx *sqlx.Tx,
	mp MedicinalProduct,
	quantity int,
) *de.DomainError {
	if mp.Quantity-quantity < 0 {
		return de.ErrQuantityTooHigh.WithParams("max", mp.Quantity)
	}

	_, err := tx.ExecContext(ctx, querySubtractMedicinalProduct, mp.ID, quantity)
	if err != nil {
		log.Error(err)
		return de.ErrSubtractMedicinalProduct
	}

	return nil
}
