package middleware

import (
	"clevergo.tech/jsend"
	"github.com/gofiber/fiber/v2"
)

// Проверка контента содержимого запроса
func CheckContentType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if content := c.Request().Header.ContentType(); string(content) != "application/json" {
			return c.Status(fiber.StatusBadRequest).JSON(jsend.NewError("error invalid content-type", fiber.StatusBadRequest, nil))
		}
		return c.Next()
	}
}
