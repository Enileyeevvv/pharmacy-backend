package http

type CreateMedicinalProduct struct {
	Name        string `json:"name" validate:"required,lte=255"`
	Description string `json:"description" validate:"required,lte=255"`
	Quantity    *int   `json:"quantity" validate:"required,lte=255"`
	MaxQuantity int    `json:"maxQuantity" validate:"required,gte=1"`
}

type GetMedicinalProductRequest struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type GetMedicinalProductResponse struct {
	ID                  int                 `json:"id"`
	Name                string              `json:"name"`
	SellName            string              `json:"sellName"`
	ATXCode             string              `json:"ATXCode"`
	Description         string              `json:"description"`
	PharmaceuticalGroup PharmaceuticalGroup `json:"pharmaceuticalGroup"`
	Company             Company             `json:"company"` // todo where image?
	Quantity            int                 `json:"quantity"`
	MaxQuantity         int                 `json:"maxQuantity"`
}

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PharmaceuticalGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
