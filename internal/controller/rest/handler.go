package rest

import "github.com/gofiber/fiber/v2"

type Handler interface {
	SaveData(c *fiber.Ctx) error
}
