package postgres

type MedicinalProduct struct {
	ID                    int    `db:"id"`
	Name                  string `db:"name"`
	ATXCode               string `db:"atx_code"`
	Description           string `db:"description"`
	PharmaceuticalGroupID int    `db:"pharmaceutical_group_id"`
	Quantity              int    `db:"quantity"`
	MaxQuantity           int    `db:"max_quantity"`
}
