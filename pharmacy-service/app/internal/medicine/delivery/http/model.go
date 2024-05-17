package http

type CreateMedicinalProductRequest struct {
	Name        string `json:"name" validate:"required,lte=255"`
	SellName    string `json:"sellName" validate:"required,lte=255"`
	ATXCode     string `json:"ATXCode" validate:"required,lte=255"`
	Description string `json:"description" validate:"required,lte=255"`
	Quantity    int    `json:"quantity" validate:"required,lte=255"`
	MaxQuantity int    `json:"maxQuantity" validate:"required,gte=1"`

	PharmaceuticalGroupID int `json:"pharmaceuticalGroupID" validate:"required"`

	CompanyName  string `json:"companyName" validate:"required"`
	ImageURL     string `json:"imageURL" validate:"required"`
	DosageFormID int    `json:"dosageFormID" validate:"required"`
}

type FetchMedicinalProductsRequest struct {
	Limit  int `query:"limit" validate:"required"`
	Offset int `query:"offset" validate:"required"`
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
	DosageFormID            int    `json:"dosageFormID"`
	DosageFormName          string `json:"dosageFormName"`
}

type FetchMedicinalProductsResponse struct {
	HasNext bool               `json:"hasNext"`
	Data    []MedicinalProduct `json:"data"`
}

type FetchPatientsRequest struct {
	Limit  int     `query:"limit" validate:"required"`
	Offset int     `query:"offset" validate:"required"`
	Name   *string `query:"name"`
}

type Patient struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Birthday  int    `json:"birthday"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

type FetchPatientsResponse struct {
	HasNext bool      `json:"hasNext"`
	Data    []Patient `json:"data"`
}

type GetPatientResponse struct {
	Data Patient `json:"data,omitempty"`
}

type FetchPrescriptionsRequest struct {
	Limit       int     `query:"limit" validate:"required"`
	Offset      int     `query:"offset" validate:"required"`
	PatientID   *int    `query:"patientID"`
	PatientName *string `query:"patientName"`
}

type Prescription struct {
	ID                       int     `json:"id"`
	StampID                  int     `json:"stampID"`
	TypeID                   int     `json:"typeID"`
	StatusID                 int     `json:"statusID"`
	MedicinalProductID       int     `json:"medicinalProductID"`
	MedicinalProductName     string  `json:"medicinalProductName"`
	MedicinalProductQuantity int     `json:"medicinalProductQuantity"`
	DoctorID                 int     `json:"doctorID"`
	DoctorName               string  `json:"doctorName"`
	PatientID                int     `json:"patientID"`
	PatientName              string  `json:"patientName"`
	PharmacistID             *int    `json:"pharmacistID"`
	PharmacistName           *string `json:"pharmacistName"`
	CreatedAt                int     `json:"createdAt"`
	UpdatedAt                int     `json:"updatedAt"`
	ExpiredAt                int     `json:"expiredAt"`
}

type FetchPrescriptionsResponse struct {
	HasNext bool           `json:"hasNext"`
	Data    []Prescription `json:"data"`
}

type GetPrescriptionResponse struct {
	Data Prescription `json:"data,omitempty"`
}

type CreateSinglePrescriptionRequest struct {
	MedicinalProductID int `json:"medicinalProductID" validate:"required"`
	PatientID          int `json:"patientID" validate:"required"`
	StampID            int `json:"stampID" validate:"required"`
	QuantityForCourse  int `json:"quantityForCourse" validate:"required"`
}

type CreateMultiplePrescriptionRequest struct {
	MedicinalProductID int `json:"medicinalProductID" validate:"required"`
	PatientID          int `json:"patientID" validate:"required"`
	StampID            int `json:"stampID" validate:"required"`
	QuantityInDose     int `json:"quantityInDose" validate:"required"`
	DoseCount          int `json:"doseCount" validate:"required"`
}

type SubmitPrescriptionRequest struct {
	ID int `json:"id" validate:"required"`
}

type CancelPrescriptionRequest struct {
	ID int `json:"id" validate:"required"`
}

type FetchPrescriptionHistoryRequest struct {
	Limit  int `query:"limit" validate:"required"`
	Offset int `query:"offset" validate:"required"`
}

type PrescriptionHistory struct {
	DoctorID       int     `json:"doctorID"`
	DoctorName     string  `json:"doctorName"`
	PharmacistID   *int    `json:"pharmacistID"`
	PharmacistName *string `json:"pharmacistName"`
	StatusID       int     `json:"statusID"`
	UpdatedAt      int     `json:"updatedAt"`
}

type FetchPrescriptionHistoryResponse struct {
	HasNext bool                  `json:"hasNext"`
	Data    []PrescriptionHistory `json:"data"`
}

type AddMedicinalProductRequest struct {
	ID       int `json:"id" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
}
