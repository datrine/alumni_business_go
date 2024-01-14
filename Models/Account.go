package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Account struct {
	ID                string `gorm:"primaryKey"`
	MemberNumber      string `gorm:"unique"`
	Email             string `gorm:"unique"`
	Password          string
	FirstName         string
	LastName          string
	Profession        *string
	JobTitle          *string
	Education         pq.StringArray `gorm:"type:text[]"`
	Certifications    pq.StringArray `gorm:"type:text[]"`
	Employer          *string
	Industry          *string
	Location          *string
	Skills            pq.StringArray `gorm:"type:text[]"`
	ProfilePictureUrl string
	GraduationYear    int
	Has_Subscribed    bool `gorm:"default:false"`
	Approved          bool `gorm:"default:false"`
	ApprovedBy        *string
	Role              string
	DOB               sql.NullTime // Uses sql.NullTime for nullable time fields
	ActivatedAt       sql.NullTime // Uses sql.NullTime for nullable time fields
	CreatedAt         time.Time    // Automatically managed by GORM for creation time
	UpdatedAt         time.Time    // Automatically managed by GORM for update time
}
