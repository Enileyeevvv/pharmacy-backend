package models

type CreateMedicinalProduct struct {
	Name        string `json:"name" validate:"required,lte=255"`
	Description string `json:"description" validate:"required,lte=255"`
}

type GetMedicinalProduct struct {
}
