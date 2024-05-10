package http

type CreateMedicinalProductRequest struct {
	Name        string `json:"name" validate:"required,lte=255"`
	SellName    string `json:"sellName" validate:"required,lte=255"`
	ATXCode     string `json:"ATXCode" validate:"required,lte=255"`
	Description string `json:"description" validate:"required,lte=255"`
	Quantity    int    `json:"quantity" validate:"required,lte=255"`
	MaxQuantity int    `json:"maxQuantity" validate:"required,gte=1"`
}

type FetchMedicinalProductsRequest struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type MedicinalProduct struct {
	ID                      int    `json:"id"`
	Name                    string `json:"name"`
	SellName                string `json:"sellName"`
	ATXCode                 string `json:"ATXCode"`
	Description             string `json:"description"`
	PharmaceuticalGroupID   int    `json:"pharmaceuticalGroupID"`
	PharmaceuticalGroupName string `json:"pharmaceuticalGroupName"`
	CompanyID               int    `json:"companyID"`
	CompanyName             string `json:"companyName"`
	Quantity                int    `json:"quantity"`
	MaxQuantity             int    `json:"maxQuantity"`
	ImageURL                string `json:"imageURL"`
}

type FetchMedicinalProductsResponse struct {
	HasNext bool               `json:"hasNext"`
	Data    []MedicinalProduct `json:"data"`
}
