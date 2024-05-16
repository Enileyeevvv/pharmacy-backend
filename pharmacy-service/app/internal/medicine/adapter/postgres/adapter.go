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

func (a *adapter) CheckMedicinalProductExists(ctx context.Context, mp usecase.MedicinalProduct) (int, *de.DomainError) {
	var mpID int
	err := a.db.GetContext(
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

func (a *adapter) CheckCompanyExists(ctx context.Context, mp usecase.MedicinalProduct) (int, *de.DomainError) {
	var cID int
	err := a.db.GetContext(
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

func (a *adapter) CreateMedicinalProduct(ctx context.Context, mp usecase.MedicinalProduct) (int, *de.DomainError) {
	var mpID int

	err := a.db.GetContext(
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

func (a *adapter) CreateCompany(ctx context.Context, mp usecase.MedicinalProduct) (int, *de.DomainError) {
	var cID int

	err := a.db.GetContext(
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

func (a *adapter) UpsertMedicinalProductCompany(ctx context.Context, mp usecase.MedicinalProduct) *de.DomainError {
	_, err := a.db.ExecContext(
		ctx,
		queryUpsertMedicinalProductCompany,
		mp.ID,
		mp.CompanyID,
		mp.ImageURL)
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

func (a *adapter) FetchPrescriptions(ctx context.Context, limit, offset int) ([]usecase.Prescription, *de.DomainError) {
	var ps []Prescription

	err := a.db.SelectContext(ctx, &ps, queryFetchPrescriptions, limit, offset)

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

func (a *adapter) CreatePrescription(ctx context.Context, p usecase.Prescription) *de.DomainError {
	var pID int

	err := a.db.GetContext(
		ctx,
		&pID,
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
		return de.ErrCreatePrescription
	}

	p.ID = pID
	p.StatusID = 1

	return a.updatePrescriptionHistory(ctx, p)
}

func (a *adapter) CheckoutPrescription(
	ctx context.Context,
	prescriptionID, pharmacistID, statusID int,
) *de.DomainError {
	var doctorID int

	err := a.db.GetContext(
		ctx,
		&doctorID,
		queryCheckoutPrescription,
		prescriptionID,
		pharmacistID,
		statusID)
	if err != nil {
		log.Error(err)
		return de.ErrCheckoutPrescription
	}

	p := usecase.Prescription{
		ID:           prescriptionID,
		StatusID:     statusID,
		DoctorID:     doctorID,
		PharmacistID: &pharmacistID,
	}

	return a.updatePrescriptionHistory(ctx, p)
}

func (a *adapter) updatePrescriptionHistory(ctx context.Context, p usecase.Prescription) *de.DomainError {
	_, err := a.db.ExecContext(
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
