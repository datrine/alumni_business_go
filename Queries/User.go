package queries

import (
	entities "github.com/datrine/alumni_business/Entities"
	models "github.com/datrine/alumni_business/Models"
	providers "github.com/datrine/alumni_business/Providers"
)

func GetAuthUserByEmail(email string) (*entities.User, error) {
	model := &models.Account{
		Email: email,
	}
	result := providers.DB.Model(model).First(model)
	if result.Error != nil {
		return nil, result.Error
	}
	userEntity := entities.GetUserEntityFromAccountModel(model)
	return userEntity, nil
}
