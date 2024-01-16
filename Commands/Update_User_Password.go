package commands

import (
	"errors"

	dtos "github.com/datrine/alumni_business/Dtos/Command"
	models "github.com/datrine/alumni_business/Models"
	providers "github.com/datrine/alumni_business/Providers"
)

func UpdateUserPassword(data *dtos.UpdateUserPasswordCommandDTO) error {
	model := &models.Account{
		ID:    data.ID,
		Email: data.Email,
	}

	result := providers.DB.Model(model).First(model)
	if result.Error != nil {
		return result.Error
	}
	if model.Password != data.OldPassword {
		return errors.New("wrong old password")
	}
	model.Password = data.NewPassword
	result2 := providers.DB.Save(&model)
	if result2.Error != nil {
		return result.Error
	}
	return nil
}
