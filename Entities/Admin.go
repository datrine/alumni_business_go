package entities

import (
	"errors"
	"time"

	dtos "github.com/datrine/alumni_business/Dtos/Command"
	models "github.com/datrine/alumni_business/Models"
	providers "github.com/datrine/alumni_business/Providers"
	"github.com/rs/xid"
)

type Admin struct {
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

func (userToAdd *Admin) Register() error {
	id := xid.New()
	model := &models.Account{
		ID:                id.String(),
		Email:             userToAdd.Email,
		LastName:          userToAdd.LastName,
		FirstName:         userToAdd.FirstName,
		MemberNumber:      userToAdd.MemberNumber,
		ProfilePictureUrl: userToAdd.ProfilePictureUrl,
	}
	result := providers.DB.Create(model)
	err := result.Error
	return err
}

func (userToAdd *Admin) GetUserByID() (*User, error) {
	id := userToAdd.ID
	model := &models.Account{
		ID: id,
	}
	result := providers.DB.First(model)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &User{
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

func (userToAdd *Admin) UpdateUserProfile(data *dtos.UpdateUserProfileCommandData) (*User, error) {
	id := userToAdd.ID
	account := &models.Account{
		ID: id,
	}
	result := providers.DB.Model(account).Omit("id", "email", "password").Updates(data)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &User{
		ID:                account.ID,
		Email:             account.Email,
		LastName:          account.LastName,
		FirstName:         account.FirstName,
		Profession:        &account.Profession.String,
		JobTitle:          &account.JobTitle.String,
		Education:         account.Education,
		Certifications:    account.Certifications,
		Employer:          &account.Employer.String,
		Industry:          &account.Industry.String,
		Location:          &account.Location.String,
		Skills:            account.Skills,
		ProfilePictureUrl: account.ProfilePictureUrl,
		GraduationYear:    account.GraduationYear,
		DOB:               &account.DOB.Time,
	}, nil
}

func (userToAdd *Admin) UpdateUserPassword(data *dtos.UpdateUserPasswordCommandDTO) error {
	accountModel := &models.Account{
		Email: data.Email,
	}

	result := providers.DB.Model(accountModel).Updates(models.Account{Password: data.NewPassword})

	affectedRows := result.RowsAffected
	if affectedRows == 0 {
		return errors.New("failed to update user password,email not valid")
	}
	return result.Error
}
