package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	PayerID    string
	PayerEmail string
	Amount     string
	Currency   string
	Platform   string
	Metadata   interface{} `gorm:"type:json"`
	Status     string
	CreatedAt  time.Time // Automatically managed by GORM for creation time
	UpdatedAt  time.Time // Automatically managed by GORM for update time
}
