package postgres

import "database/sql"

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

type Patient struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Birthday  int    `db:"birthday"`
	CreatedAt int    `db:"created_at"`
	UpdatedAt int    `db:"updated_at"`
}

type Prescription struct {
	ID                       int            `db:"id"`
	StampID                  int            `db:"stamp_id"`
	TypeID                   int            `db:"type_id"`
	StatusID                 int            `db:"status_id"`
	MedicinalProductID       int            `db:"medicinal_product_id"`
	MedicinalProductName     string         `db:"medicinal_product_name"`
	MedicinalProductQuantity int            `db:"medicinal_product_quantity"`
	DoctorID                 int            `db:"doctor_id"`
	DoctorName               string         `db:"doctor_name"`
	PatientID                int            `db:"patient_id"`
	PatientName              string         `db:"patient_name"`
	PharmacistID             sql.NullInt64  `db:"pharmacist_id"`
	PharmacistName           sql.NullString `db:"pharmacist_name"`
	CreatedAt                int            `db:"created_at"`
	UpdatedAt                int            `db:"updated_at"`
	ExpiredAt                int            `db:"expired_at"`
}
