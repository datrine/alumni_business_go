package dtos

import "time"

type UpdateUserProfileCommandData struct {
	ID                string
	FirstName         string
	LastName          string
	Profession        *string
	JobTitle          *string
	Education         []string
	Certifications    []string
	Employer          *string
	Industry          *string
	Location          *string
	Skills            []string
	ProfilePictureUrl string
	GraduationYear    int
	DOB               *time.Time
	ActivatedAt       *time.Time
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
