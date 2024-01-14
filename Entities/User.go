package entities

import (
	"errors"
	"time"

	dtos "github.com/datrine/alumni_business/Dtos/Command"
	models "github.com/datrine/alumni_business/Models"
	providers "github.com/datrine/alumni_business/Providers"
	"github.com/rs/xid"
)

type User struct {
	ID                string
	MemberNumber      string
	Email             string
	Password          string
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
	Role              string
	DOB               *time.Time
	ActivatedAt       *time.Time // Uses sql.NullTime for nullable time fields
	CreatedAt         *time.Time // Automatically managed by GORM for creation time
	UpdatedAt         *time.Time // Automatically managed by GORM for update time
}

type BasicLoginData struct {
	Identifier string
	Password   string
}

func (userToAdd *User) Register() error {
	id := xid.New()
	model := models.Account{
		ID:                id.String(),
		Email:             userToAdd.Email,
		Password:          userToAdd.Password,
		LastName:          userToAdd.LastName,
		FirstName:         userToAdd.FirstName,
		MemberNumber:      userToAdd.MemberNumber,
		ProfilePictureUrl: userToAdd.ProfilePictureUrl,
		Role:              "USER",
	}
	result := providers.DB.Create(&model)
	err := result.Error
	userToAdd.ID = model.ID
	return err
}

func GetUserEntityFromAccountModel(model *models.Account) *User {
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
		Role:              model.Role,
	}
}

func GetUserByID(id string) (*User, error) {
	model := &models.Account{
		ID: id,
	}
	result := providers.DB.First(model)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return GetUserEntityFromAccountModel(model), nil
}

func BasicLogin(id *BasicLoginData) (*User, error) {
	account := &models.Account{}
	result := providers.DB.
		Where(&models.Account{Email: id.Identifier}).
		Or(&models.Account{MemberNumber: id.Identifier}).First(account)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return GetUserEntityFromAccountModel(account), nil
}

func (userToAdd *User) UpdateUserProfile(data *dtos.UpdateUserProfileCommandData) (*User, error) {
	id := userToAdd.ID
	account := &models.Account{
		ID: id,
	}
	result := providers.DB.Model(account).Omit("id", "email", "password", "member_number").Updates(data)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return GetUserEntityFromAccountModel(account), nil
}

func (userToAdd *User) UpdateUserPassword(data *dtos.UpdateUserPasswordCommandDTO) error {
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
