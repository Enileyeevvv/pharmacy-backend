package models

type CreateMedicinalProduct struct {
	Name        string `json:"name" validate:"required,lte=255"`
	Description string `json:"description" validate:"required,lte=255"`
	Quantity    *int   `json:"quantity" validate:"required,lte=255"`
	MaxQuantity int    `json:"maxQuantity" validate:"required,gte=1"`
}

type GetMedicinalProduct struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}
