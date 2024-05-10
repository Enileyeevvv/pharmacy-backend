package http

import "github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/internal/medicine/usecase"

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
