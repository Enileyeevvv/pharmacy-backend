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
	DosageFormID            int
	DosageFormName          string
}

type Patient struct {
	ID        int
	Name      string
	Email     string
	Birthday  int
	CreatedAt int
	UpdatedAt int
}

type Prescription struct {
	ID                       int
	StampID                  int
	TypeID                   int
	StatusID                 int
	MedicinalProductID       int
	MedicinalProductName     string
	MedicinalProductQuantity int
	DoctorID                 int
	DoctorName               string
	PatientID                int
	PatientName              string
	PharmacistID             *int
	PharmacistName           *string
	CreatedAt                int
	UpdatedAt                int
	ExpiredAt                int
}

type PrescriptionHistory struct {
	ID             int
	PrescriptionID int
	DoctorID       int
	DoctorName     string
	PharmacistID   *int
	PharmacistName *string
	StatusID       int
	UpdatedAt      int
}
