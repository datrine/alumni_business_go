package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	AuthorId     string
	ParentPostId string
	Content      string
	Comments     *[]Comment
	ContentType  string
	CreatedAt    time.Time // Automatically managed by GORM for creation time
	UpdatedAt    time.Time // Automatically managed by GORM for update time
}
