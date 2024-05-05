package models

type User struct {
	ID        uint   `db:"id" json:"id" validate:"required,uuid" gorm:"primary_key"`
	Login     string `db:"login" json:"login" validate:"required,lte=255"`
	Password  string `db:"password" json:"password,omitempty" validate:"required,lte=255"`
	Status    int    `db:"status" json:"status" validate:"required,len=1"`
	CreatedAt int64  `gorm:"autoCreateTime" db:"created_at" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" db:"updated_at" json:"updated_at"`
}
