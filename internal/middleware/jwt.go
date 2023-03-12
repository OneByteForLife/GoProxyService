package middleware

import (
	"os"

	"clevergo.tech/jsend"
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
	return c.Status(fiber.StatusUnauthorized).JSON(jsend.NewError("error invalid jwt token", fiber.StatusUnauthorized, nil))
}
