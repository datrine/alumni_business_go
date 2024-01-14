package dtos

import "time"

type BasicLoginRequestJSONDTO struct {
	Identifier string //`json:"identifier"`
	Password   string //`json:"password"`
}

type RegisterUserRequestJSONDTO struct {
	MemberNumber string `json:"member_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type RegisterUserRequestFORMDTO struct {
	MemberNumber      string `validate:"required"`
	Email             string `validate:"required"`
	Password          string `validate:"required"`
	FirstName         string `validate:"required"`
	LastName          string `validate:"required"`
	ProfilePictureUrl string `validate:"required"`
}

type UpdateUserProfileRequestFORMDTO struct {
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
}

type User struct {
	MemberNumber   string    `json:"member_number"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Profession     *string   `json:"profession"`
	JobTitle       *string   `json:"job_title"`
	Education      *[]string `json:"education"`
	Certifications *[]string `json:"certifications"`
	Employer       *string   `json:"employer"`
	Industry       *string   `json:"industry"`
	Location       *string   `json:"location"`
	Skills         *[]string `json:"skills"`
	GraduationYear int       `json:"graduation_year"`
}