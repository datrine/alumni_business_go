package commands

import (
	"strings"
	"time"

	dtos "github.com/datrine/alumni_business/Dtos/Command"
	entities "github.com/datrine/alumni_business/Entities"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUserEntity struct {
	User  *entities.User
	Token string
}

func BasicLogin(data *dtos.BasicLoginCommandDTO) (*AuthUserEntity, error) {
	userEntity, err := entities.BasicLogin(&entities.BasicLoginData{
		Identifier: data.Identifier,
		Password:   data.Password,
	})
	if err != nil {
		return nil, err
	}
	claims := jwt.MapClaims{
		"name": strings.Join([]string{userEntity.FirstName, userEntity.LastName}, " "),
		"role": userEntity.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	authUser := &AuthUserEntity{User: userEntity, Token: tokenString}

	return authUser, nil
}
