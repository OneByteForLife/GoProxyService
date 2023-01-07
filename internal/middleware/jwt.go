package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func CheckJwtToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"api_version": "1.0", "status_code": fiber.StatusUnauthorized, "description": "Incorrect or invalid access token!"})
}
