package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	AuthorId    string
	Title       string
	Text        string
	Media       []Media
	Comments    []Comment
	Action      []PostAction `gorm:"foreignKey:PostId"`
	ContentType string
	ActivatedAt sql.NullTime // Uses sql.NullTime for nullable time fields
	CreatedAt   time.Time    // Automatically managed by GORM for creation time
	UpdatedAt   time.Time    // Automatically managed by GORM for update time
}
