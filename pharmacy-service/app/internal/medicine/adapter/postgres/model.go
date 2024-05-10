package postgres

type MedicinalProduct struct {
	ID                      int    `db:"id"`
	Name                    string `db:"name"`
	SellName                string `db:"sell_name"`
	ATXCode                 string `db:"atx_code"`
	Description             string `db:"description"`
	PharmaceuticalGroupID   int    `db:"pharmaceutical_group_id"`
	PharmaceuticalGroupName string `db:"pharmaceutical_group_name"`
	Quantity                int    `db:"quantity"`
	MaxQuantity             int    `db:"max_quantity"`
	CompanyID               int    `db:"company_id"`
	CompanyName             string `db:"company_name"`
	ImageURL                string `db:"image_url"`
}
