package utils

import (
	"errors"
	"strings"
	"time"

	commandDtos "github.com/datrine/alumni_business/Dtos/Command"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetAuthPayload(c *fiber.Ctx) (*commandDtos.JWTPayload, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	payload, ok := claims["sub"].(map[string]interface{})
	if !ok {
		return nil, errors.New("wrong claims")
	}
	id := payload["ID"].(string)
	email := payload["Email"].(string)
	return &commandDtos.JWTPayload{
		ID:    id,
		Email: email,
	}, nil
}

func SetAuthClaims(data *commandDtos.JWTPayload) (string, error) {
	claims := jwt.MapClaims{
		"name": strings.Join([]string{data.FirstName, data.LastName}, " "),
		"role": data.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"sub":  data,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
