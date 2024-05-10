package postgres

import "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"

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
