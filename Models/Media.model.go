package models

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	PostId string `gorm:"foreignKey:ID"`
	URL    string
	Type   string
}
