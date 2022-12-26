package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RespStatus(api string, status int, description string, content interface{}) fiber.Map {
	return fiber.Map{
		"api_version": api,
		"status_code": status,
		"description": description,
		"content":     content,
	}
}

func SavingData(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Success", c.Body()))
}
