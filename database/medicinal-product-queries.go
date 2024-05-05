package database

import "backend/app/models"

func GetMedicinalProduct() {
	medicinalProduct := models.MedicinalProduct{
		Id:          1,
		Name:        "Medical Product",
		Description: "Description",
	}

	GetDB().Find(&medicinalProduct)
}
