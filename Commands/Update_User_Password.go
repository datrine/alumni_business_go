package commands

import (
	"errors"
	"fmt"

	dtos "github.com/datrine/alumni_business/Dtos/Command"
	entities "github.com/datrine/alumni_business/Entities"
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
	result2 := providers.DB.Save(model)
	if result2.Error != nil {
		return nil, result.Error
	}
	fmt.Println("ooooooo")
	token, err := utils.SetAuthClaims(&dtos.JWTPayload{
		ID:        model.ID,
		Email:     model.Email,
		FirstName: model.FirstName,
		LastName:  model.LastName,
	})
	if err != nil {
		return nil, err
	}
	userEntity := entities.GetUserEntityFromAccountModel(model)
	user := &AuthUserEntity{
		User:  userEntity,
		Token: token,
	}
	return user, nil
}
