package middleware

import (
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

type AuthFailedError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var Auth = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(fiber.StatusUnauthorized).JSON(&AuthFailedError{
			Message: err.Error(),
			Status:  fiber.StatusUnauthorized,
		})
	},
})
