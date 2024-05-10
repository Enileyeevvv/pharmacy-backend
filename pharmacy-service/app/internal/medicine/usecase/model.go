package usecase

type MedicinalProductJSON struct {
	Id          uint   `json:"id" gorm:"unique;primary_key;autoIncrement"`
	Name        string `json:"name" validate:"required,lte=255"`
	Description string `json:"description" validate:"required,lte=255"`
	Quantity    int    `json:"quantity"`
	MaxQuantity int    `json:"maxQuantity" validate:"required,gte=1"`
}

type MedicinalProduct struct {
	ID                  int
	Name                string
	SellName            string
	ATXCode             string
	Description         string
	PharmaceuticalGroup PharmaceuticalGroup
	Company             Company // todo where image?
	Quantity            int
	MaxQuantity         int
}

type Company struct {
	ID   int
	Name string
}

type PharmaceuticalGroup struct {
	ID   int
	Name string
}
