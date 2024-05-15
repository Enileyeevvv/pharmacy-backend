package http

import "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"

func MapCreateMedicinalProductRequest(req CreateMedicinalProductRequest) usecase.MedicinalProduct {
	return usecase.MedicinalProduct{
		Name:        req.Name,
		SellName:    req.SellName,
		ATXCode:     req.ATXCode,
		Description: req.Description,
		Quantity:    req.Quantity,
		MaxQuantity: req.MaxQuantity,

		PharmaceuticalGroupID: req.PharmaceuticalGroupID,

		CompanyName: req.CompanyName,
		ImageURL:    req.ImageURL,
	}
}

func MapFetchMedicinalProductsResponse(mps []usecase.MedicinalProduct, hasNext bool) FetchMedicinalProductsResponse {
	mpsData := make([]MedicinalProduct, 0)
	if mps != nil {
		for _, mp := range mps {
			mpEntry := MedicinalProduct{
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
	}

	return FetchMedicinalProductsResponse{
		HasNext: hasNext,
		Data:    mpsData,
	}
}

func MapFetchPatientsResponse(ps []usecase.Patient, hasNext bool) FetchPatientsResponse {
	psData := make([]Patient, 0)
	if ps != nil {
		for _, mp := range ps {
			pEntry := Patient{
				ID:        mp.ID,
				Name:      mp.Name,
				Email:     mp.Email,
				Birthday:  mp.Birthday,
				CreatedAt: mp.CreatedAt,
				UpdatedAt: mp.UpdatedAt,
			}

			psData = append(psData, pEntry)
		}
	}

	return FetchPatientsResponse{
		HasNext: hasNext,
		Data:    psData,
	}
}

func MapGetPatientResponse(p usecase.Patient) GetPatientResponse {
	return GetPatientResponse{
		Data: Patient{
			ID:        p.ID,
			Name:      p.Name,
			Email:     p.Email,
			Birthday:  p.Birthday,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
	}
}

func MapGetPrescriptionResponse(p usecase.Prescription) GetPrescriptionResponse {
	return GetPrescriptionResponse{
		Data: Prescription{
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
			PharmacistID:             p.PharmacistID,
			PharmacistName:           p.PharmacistName,
			CreatedAt:                p.CreatedAt,
			UpdatedAt:                p.UpdatedAt,
			ExpiredAt:                p.ExpiredAt,
		},
	}
}

func MapFetchPrescriptionsResponse(ps []usecase.Prescription, hasNext bool) FetchPrescriptionsResponse {
	psData := make([]Prescription, 0)
	if ps != nil {
		for _, p := range ps {
			pEntry := Prescription{
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
				PharmacistID:             p.PharmacistID,
				PharmacistName:           p.PharmacistName,
				CreatedAt:                p.CreatedAt,
				UpdatedAt:                p.UpdatedAt,
				ExpiredAt:                p.ExpiredAt,
			}

			psData = append(psData, pEntry)
		}
	}

	return FetchPrescriptionsResponse{
		HasNext: hasNext,
		Data:    psData,
	}
}
