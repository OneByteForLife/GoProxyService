package rest

import (
	"GoProxyService/internal/middleware"
	"time"

	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func Init(api Handler) {
	// Конифгурация сервера
	app := fiber.New(fiber.Config{
		JSONDecoder:  gojson.Unmarshal,
		JSONEncoder:  gojson.Marshal,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	})

	app.Use(logger.New())

	app.Post("api/v1/add", middleware.CheckContentType(), middleware.CheckJwtToken(), api.SaveData)

	if err := app.Listen(":8080"); err != nil {
		logrus.Fatalf("error up server: %v", err)
	}
}
