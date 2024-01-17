package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type PostAction struct {
	gorm.Model
	ID          string
	ActorId     string
	ActionType  string
	PostId      string
	ActivatedAt sql.NullTime // Uses sql.NullTime for nullable time fields
	CreatedAt   time.Time    // Automatically managed by GORM for creation time
	UpdatedAt   time.Time    // Automatically managed by GORM for update time
}
