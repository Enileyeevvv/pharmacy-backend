package usecase

type MedicinalProductJSON struct {
	Id          uint   `json:"id" gorm:"unique;primary_key;autoIncrement"`
	Name        string `json:"name" validate:"required,lte=255"`
	Description string `json:"description" validate:"required,lte=255"`
	Quantity    int    `json:"quantity"`
	MaxQuantity int    `json:"maxQuantity" validate:"required,gte=1"`
}

type MedicinalProduct struct {
	ID                      int
	Name                    string
	SellName                string
	ATXCode                 string
	Description             string
	PharmaceuticalGroupID   int
	PharmaceuticalGroupName string
	CompanyID               int
	CompanyName             string
	Quantity                int
	MaxQuantity             int
	ImageURL                string
}
