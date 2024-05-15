package usecase

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

type Patient struct {
	ID        int
	Name      string
	Email     string
	Birthday  int
	CreatedAt int
	UpdatedAt int
}
