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

func (transaction *Transaction) AfterSave(tx *gorm.DB) (err error) {
	if transaction.Status == "VERIFIED" {
		accModel := &Account{
			ID: transaction.PayerID,
		}
		accModel.Has_Subscribed = true
		tx.Save(&accModel)
	}
	return
}
