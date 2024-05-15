package postgres

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"
)

func MapMedicinalProductSlice(mps []MedicinalProduct) []usecase.MedicinalProduct {
	if mps == nil {
		return make([]usecase.MedicinalProduct, 0)
	}

	mpsData := make([]usecase.MedicinalProduct, 0)
	for _, mp := range mps {
		mpEntry := usecase.MedicinalProduct{
			ID:                      mp.ID,
			Name:                    mp.Name,
			SellName:                mp.SellName,
			ATXCode:                 mp.ATXCode,
			Description:             mp.Description,
			PharmaceuticalGroupID:   mp.PharmaceuticalGroupID,
			PharmaceuticalGroupName: mp.PharmaceuticalGroupName,
			CompanyID:               mp.CompanyID,
			CompanyName:             mp.CompanyName,
			Quantity:                mp.Quantity,
			MaxQuantity:             mp.MaxQuantity,
			ImageURL:                mp.ImageURL,
		}
		mpsData = append(mpsData, mpEntry)
	}

	return mpsData
}

func MapPatientSlice(ps []Patient) []usecase.Patient {
	if ps == nil {
		return make([]usecase.Patient, 0)
	}

	psData := make([]usecase.Patient, 0)
	for _, p := range ps {
		pEntry := usecase.Patient{
			ID:        p.ID,
			Name:      p.Name,
			Email:     p.Email,
			Birthday:  p.Birthday,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		}
		psData = append(psData, pEntry)
	}

	return psData
}

func MapPatient(p Patient) usecase.Patient {
	return usecase.Patient{
		ID:        p.ID,
		Name:      p.Name,
		Email:     p.Email,
		Birthday:  p.Birthday,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func MapPrescription(p Prescription) usecase.Prescription {
	var pharmacistIDPointer *int = nil
	if p.PharmacistID.Valid {
		pharmacistID := int(p.PharmacistID.Int64)
		pharmacistIDPointer = &pharmacistID
	}

	var pharmacistNamePointer *string = nil
	if p.PharmacistName.Valid {
		pharmacistName := p.PharmacistName.String
		pharmacistNamePointer = &pharmacistName
	}

	return usecase.Prescription{
		ID:                       p.ID,
		StampID:                  p.StampID,
		TypeID:                   p.TypeID,
		StatusID:                 p.StatusID,
		MedicinalProductID:       p.MedicinalProductID,
		MedicinalProductName:     p.MedicinalProductName,
		MedicinalProductQuantity: p.MedicinalProductQuantity,
		DoctorID:                 p.DoctorID,
		DoctorName:               p.DoctorName,
		PatientID:                p.PatientID,
		PatientName:              p.PatientName,
		PharmacistID:             pharmacistIDPointer,
		PharmacistName:           pharmacistNamePointer,
		CreatedAt:                p.CreatedAt,
		UpdatedAt:                p.UpdatedAt,
		ExpiredAt:                p.ExpiredAt,
	}
}

func MapPrescriptions(ps []Prescription) []usecase.Prescription {
	if ps == nil {
		return make([]usecase.Prescription, 0)
	}

	res := make([]usecase.Prescription, 0)

	for _, p := range ps {
		res = append(res, MapPrescription(p))
	}

	return res
}
