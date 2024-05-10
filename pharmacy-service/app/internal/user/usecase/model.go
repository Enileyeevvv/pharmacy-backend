package usecase

type User struct {
	ID           int    `db:"id" json:"id" validate:"required,uuid" gorm:"unique;primary_ke;autoIncrement"`
	Login        string `db:"login" json:"login" validate:"required,lte=255"`
	PasswordHash string `db:"password" json:"password,omitempty" validate:"required,lte=255"`
	Status       int    `db:"status" json:"status" validate:"required"`
	CreatedAt    int    `gorm:"autoCreateTime" db:"created_at" json:"created_at"`
	UpdatedAt    int    `gorm:"autoUpdateTime" db:"updated_at" json:"updated_at"`
	RoleID       int
}
