package commands

import (
	"errors"

	dtos "github.com/datrine/alumni_business/Dtos/Command"
	models "github.com/datrine/alumni_business/Models"
	providers "github.com/datrine/alumni_business/Providers"
	utils "github.com/datrine/alumni_business/Utils"
)

func UpdateUserPassword(data *dtos.UpdateUserPasswordCommandDTO) (*AuthUserEntity, error) {
	model := &models.Account{
		ID:    data.ID,
		Email: data.Email,
	}

	result := providers.DB.Model(model).First(model)
	if result.Error != nil {
		return nil, result.Error
	}
	if model.Password != data.OldPassword {
		return nil, errors.New("wrong old password")
	}
	model.Password = data.NewPassword
	result2 := providers.DB.Save(&model)
	if result2.Error != nil {
		return nil, result.Error
	}

	token, err := utils.SetAuthClaims(&dtos.JWTPayload{})
	if result2.Error != nil {
		return nil, err
	}

	user := &AuthUserEntity{
		User:  nil,
		Token: token,
	}
	return user, nil
}
