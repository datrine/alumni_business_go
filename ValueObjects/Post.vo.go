package valueobjects

import (
	"time"

	models "github.com/datrine/alumni_business/Models"
	providers "github.com/datrine/alumni_business/Providers"
)

type Post struct {
	ID                string
	MemberNumber      string // Uses sql.NullString to handle nullable strings
	Email             string
	Password          string
	FirstName         string
	LastName          string
	Profession        *string
	JobTitle          *string
	Education         *[]string
	Certifications    *[]string
	Employer          *string
	Industry          *string
	Location          *string
	Skills            *[]string
	ProfilePictureUrl string
	GraduationYear    int
	DOB               *time.Time
	ActivatedAt       *time.Time // Uses sql.NullTime for nullable time fields
	CreatedAt         *time.Time // Automatically managed by GORM for creation time
	UpdatedAt         *time.Time // Automatically managed by GORM for update time
}

func (postToAdd *Post) Post() (*Post, error) {
	id := postToAdd.ID
	model := &models.Account{
		ID: id,
	}
	result := providers.DB.First(model)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:                model.ID,
		Email:             model.Email,
		LastName:          model.LastName,
		FirstName:         model.FirstName,
		Profession:        &model.Profession.String,
		JobTitle:          &model.JobTitle.String,
		Education:         model.Education,
		Certifications:    model.Certifications,
		Employer:          &model.Employer.String,
		Industry:          &model.Industry.String,
		Location:          &model.Location.String,
		Skills:            model.Skills,
		ProfilePictureUrl: model.ProfilePictureUrl,
		GraduationYear:    model.GraduationYear,
		DOB:               &model.DOB.Time,
	}, nil
}
