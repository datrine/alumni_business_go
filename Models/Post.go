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
	Content     string
	Comments    []Comment `gorm:"foreignKey:ParentPostId"`
	ContentType string
	ActivatedAt sql.NullTime // Uses sql.NullTime for nullable time fields
	CreatedAt   time.Time    // Automatically managed by GORM for creation time
	UpdatedAt   time.Time    // Automatically managed by GORM for update time
}
