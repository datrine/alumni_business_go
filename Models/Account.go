package models

import (
	"database/sql"
	"time"
)

type Account struct {
	ID                string `gorm:"primaryKey"`
	MemberNumber      string `gorm:"unique"`
	Email             string `gorm:"unique"`
	Password          string
	FirstName         string
	LastName          string
	Profession        sql.NullString
	JobTitle          sql.NullString
	Education         []string `gorm:"json"`
	Certifications    []string `gorm:"json"`
	Employer          sql.NullString
	Industry          sql.NullString
	Location          sql.NullString
	Skills            []string `gorm:"json"`
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
