package utils

import (
	"errors"

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
