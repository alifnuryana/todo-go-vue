package middleware

import (
	"github.com/alifnuryana/go-auth-jwt/helpers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(helpers.Load("JWT_KEY")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "failed",
				"code":    fiber.StatusUnauthorized,
				"message": err.Error(),
			})
		},
	})
}
