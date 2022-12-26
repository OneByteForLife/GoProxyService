package middleware

import (
	"GoProxyService/internal/routes"

	"github.com/gofiber/fiber/v2"
)

// Проверка контента содержимого запроса
func CheckContentType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if content := c.Request().Header.ContentType(); string(content) != "application/json" {
			return c.Status(fiber.StatusBadRequest).JSON(routes.RespStatus("1.0", fiber.StatusBadRequest, "Incorrect Content-Type", nil))
		}
		return c.Next()
	}
}
