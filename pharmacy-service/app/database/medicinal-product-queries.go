package database

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/models"
)

func GetMedicinalProduct(filters *models.GetMedicinalProduct) ([]models.MedicinalProduct, error) {
	var medicinalProduct []models.MedicinalProduct

	result := GetDB().Limit(filters.Limit + 1).Offset(filters.Offset).Find(&medicinalProduct)

	return medicinalProduct, result.Error
}

func CreateMedicinalProduct(medProduct *models.MedicinalProduct) error {
	result := GetDB().Create(&medProduct)

	return result.Error
}
