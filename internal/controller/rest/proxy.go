package rest

import (
	"GoProxyService/internal/usecase/proxy"

	"clevergo.tech/jsend"
	"github.com/gofiber/fiber/v2"
)

type ProxyHandler struct {
	service proxy.Service
}

func NewHandler(s proxy.Service) Handler {
	return &ProxyHandler{s}
}

func (h *ProxyHandler) SaveData(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	if err := h.service.SaveData(c.Body()); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(jsend.NewError(err.Error(), fiber.StatusBadRequest, nil))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}
