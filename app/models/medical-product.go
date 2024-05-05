package models

type MedicinalProduct struct {
	Id          uint   `json:"id" gorm:"unique;primary_ke;autoIncrement"`
	Name        string `json:"name" validate:"required,lte=255"`
	Description string `json:"description" validate:"required,lte=255"`
}
