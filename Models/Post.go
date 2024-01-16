package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	AuthorId    string `gorm:"foreignKey:AuthorId"`
	Title       string
	Text        string
	Media       []Media
	Comments    []Comment    `gorm:"foreignKey:ParentPostId"`
	Action      []PostAction `gorm:"foreignKey:ID;references:MemberNumber"`
	ContentType string
	ActivatedAt sql.NullTime // Uses sql.NullTime for nullable time fields
	CreatedAt   time.Time    // Automatically managed by GORM for creation time
	UpdatedAt   time.Time    // Automatically managed by GORM for update time
}
